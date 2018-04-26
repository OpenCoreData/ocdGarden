package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/frictionlessdata/datapackage-go/datapackage"
	"github.com/frictionlessdata/datapackage-go/validator"
)

// TODO  make all directories and filenames match regex:  "^([-a-z0-9._/])+$"

func main() {
	fmt.Println("Frictionless Data Package Bulder")
	f := make(map[string][]string)

	// f["proj1"] = []string{"./dataVault/testproj1/data.csv", "./dataVault/testproj1/population.csv"}
	// f["proj2"] = []string{"./dataVault/testproj2/data.csv", "./dataVault/testproj2/subbin2-1/population.csv", "./dataVault/testproj2/subbin2-1/subsub/test.csv"}

	// /media/fils/seagate/dataVault  (set by -dir)
	f["testproj1"] = []string{"data.csv", "population.csv"}
	f["testproj2"] = []string{"data.csv", "subbin2-1/population.csv", "subbin2-1/subsub/test.csv"}

	packagedir := "/media/fils/seagate/packages"
	vaultdir := "/media/fils/seagate/dataVault"
	tempdir := "/media/fils/seagate/tmp"

	pkgBuilder(f, vaultdir, tempdir, packagedir)
}

// TODO integrate with CSDCO walker code
func pkgBuilder(f map[string][]string, vaultdir, tempdir, packagedir string) {
	//fmt.Println(f)

	// set up temp directory, copy files in, generate zip from that tmp directory
	dir, err := ioutil.TempDir(tempdir, "")
	if err != nil {
		log.Fatal(err)
	}

	// defer os.RemoveAll(dir) // clean up

	pm := make(map[string][]string)

	for k, v := range f {
		fmt.Printf("Name for the package: %s\n", k)
		projdir := fmt.Sprintf("%s/%s", dir, k)
		projdatadir := fmt.Sprintf("%s/%s/data", dir, k)
		os.Mkdir(projdir, os.ModePerm)
		os.Mkdir(projdatadir, os.ModePerm)

		fa := []string{}

		for i := range v {
			pdd := projdatadir
			fn := filepath.Base(v[i])
			d := filepath.Dir(v[i])

			if len(strings.Split(d, "/")) > 1 {
				fmt.Println(d)
				dirsplit := strings.Split(d, "/")
				fmt.Println(dirsplit)
				sp := fmt.Sprintf("%s/%s", pdd, d)
				pdd = sp
			}

			pdd = cleanstring(pdd)
			fmt.Printf("Projdata dir is %s \n", pdd)

			err = os.MkdirAll(pdd, os.ModePerm)
			if err != nil {
				log.Println("in make dir all")
				panic(err)
			}

			fqp := fmt.Sprintf("%s/%s/%s", vaultdir, k, v[i])
			fn = cleanstring(fn)
			err = copyFileContents(fqp, pdd+"/"+fn)

			fa = append(fa, strings.TrimPrefix(pdd+"/"+fn, dir+"/"+k+"/"))

			if err != nil {
				log.Println("in copy file")
				panic(err)
			}
		}
		pm[k] = fa
	}

	fmt.Println(pm)

	// loop on the new map from above and move in and generate the datapackage json and zip packages...
	for i, j := range pm {
		fmt.Println("-------------  loop start  ---------------------------")
		fmt.Println(i)
		fmt.Println(j)

		projdir := fmt.Sprintf("%s/%s", dir, i)
		// // change working directory
		fmt.Printf("The projdir is %s\n", projdir)
		err = os.Chdir(projdir)
		log.Println(dir)
		if err != nil {
			log.Println("in change dir")
			panic(err)
		}

		descriptor, err := makeDescriptor(j)
		if err != nil {
			log.Println("in make descriptor call")
			panic(err)
		}
		fmt.Println(descriptor)

		pkg, err := datapackage.New(descriptor, ".", validator.InMemoryLoader())
		if err != nil {
			log.Println(err)
			panic(err)
		}

		zipfp := fmt.Sprintf("%s/%s.zip", packagedir, i)
		err = pkg.Zip(zipfp)
		if err != nil {
			log.Println(err)
			panic(err)
		}

		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("dir check 1 %s \n", pwd)

		err = os.Chdir("../../..") // TODO: need to replace with a explicate set directory.
		fmt.Println("changing back up...")
		if err != nil {
			log.Println("chdir back up...")
			panic(err)
		}

		pwd, err = os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("dir check 2 %s \n", pwd)

		fmt.Println("-------------  loop end  ---------------------------")
	}

}

// func makeDescriptor(f []string) ([]byte, error) {
func makeDescriptor(f []string) (map[string]interface{}, error) {
	var vma []interface{} //  was []map[string]interface{}  //

	for i := range f {
		vm := make(map[string]interface{})
		vm["name"] = filepath.Base(f[i]) // base name only  (might be dups in different sub dirs
		vm["path"] = f[i]                // tmp + data + path
		// vm["format"] = "file" //  remove?  replace with something else from spec...
		vma = append(vma, vm)
	}

	descriptor := map[string]interface{}{
		"resources": vma,
	}

	// OLD
	// descriptor = map[string]interface{}{
	// 	"resources": []interface{}{
	// 		map[string]interface{}{
	// 			"name":   "datatest1",
	// 			"path":   "./data/data.csv",
	// 			"format": "csv",
	// 			// "profile": "tabular-data-resource",
	// 		},
	// 		map[string]interface{}{
	// 			"name":   "population",
	// 			"path":   "./data/population.csv",
	// 			"format": "csv",
	// 		},
	// 	},
	// }

	// j, _ := json.MarshalIndent(descriptor, "", " ")
	// fmt.Println(string(j))

	return descriptor, nil
}

func copyFileContents(src, dst string) (err error) {
	fmt.Printf("Copy %s to %s\n", src, dst)
	in, err := os.Open(src)
	if err != nil {
		log.Println(err)
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		log.Println(err)
		return
	}
	err = out.Sync()
	return
}

func cleanstring(s string) string {
	// Make a Regex to say we only want
	reg, err := regexp.Compile("[^a-z0-9._/]+")
	if err != nil {
		log.Fatal(err)
	}
	sl := strings.ToLower(s)
	processedString := reg.ReplaceAllString(sl, "")

	return processedString
}
