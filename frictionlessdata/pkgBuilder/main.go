package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	//"github.com/frictionlessdata/datapackage-go/datapackage"
	//"github.com/frictionlessdata/datapackage-go/validator"
)

func main() {
	fmt.Println("Frictionless Data Package Bulder")
	f := make(map[string][]string)
	f["proj1"] = []string{"./dataVault/testproj1/data.csv", "./dataVault/testproj1/population.csv"}
	f["proj2"] = []string{"./dataVault/testproj2/data.csv", "./dataVault/testproj2/subbin2-1/population.csv", "./dataVault/testproj2/subbin2-1/subsub/test.csv"}
	pkgBuilder(f)
}

// TODO integrate with CSDCO walker code
func pkgBuilder(f map[string][]string) {
	fmt.Println(f)

	// TODO..  make the following a function and pass it a
	// map like I use in the directory walk program

	// TODO
	// set up temp directory, copy files in, generate zip from that tmp directory
	dir, err := ioutil.TempDir("tmp", "")
	if err != nil {
		log.Fatal(err)
	}

	// defer os.RemoveAll(dir) // clean up

	// test cycle through map..  then extend the copy file section
	for k, v := range f {
		fmt.Printf("Name for the package: %s\n", k)
		// make data directory inside temp
		projdir := fmt.Sprintf("%s/%s", dir, k)
		projdatadir := fmt.Sprintf("%s/%s/data", dir, k)
		os.Mkdir(projdir, os.ModePerm)
		os.Mkdir(projdatadir, os.ModePerm)
		//d, err := makeDescriptor(v)
		//if err != nil {
		//	log.Println("Calling make desriptor")
		//	panic(err)
		//}
		//fmt.Println(string(d))
		for i := range v {
			pdd := projdatadir
			// TODO need concepts of "vault" location...  in order to strip concpets...
			// vp := "./dataVault/"+k
			fn := filepath.Base(v[i])
			dir := filepath.Dir(v[i])
			//fmt.Printf("\nDir %s : %s  \n", dir, fn)

			rel, err := filepath.Rel(dir, "dataVault")
			rpd := len(strings.Split(rel, "/"))
			//fmt.Printf("%d     %q: %q %v\n", rpd, "dataVault", rel, err)
			if rpd > 1 {
				dirsplit := strings.Split(dir, "/")
				sp := fmt.Sprintf("%s/%s", pdd, strings.Join(dirsplit[len(dirsplit)-(rpd-1):len(dirsplit)], "/")) // ref: https://github.com/golang/go/wiki/SliceTricks
				pdd = sp
			}

			fmt.Printf("Projdata is %s \n", pdd)
			err = os.MkdirAll(pdd, os.ModePerm)
			if err != nil {
				log.Println("in make dir all")
				panic(err)
			}
			// copy files  os.MkdirAll  filepath.Dir
			err = copyFileContents(v[i], pdd+"/"+fn)
			if err != nil {
				log.Println("in copy file")
				panic(err)
			}
			// TODO ..  need to scope in everyting below making
			// what I can into external func calls...
		}
	}

	// change working directory
	err = os.Chdir(dir)
	log.Println(dir)
	if err != nil {
		log.Println("in change dir")
		panic(err)
	}

	// TODO  look inside the ocdWeb utils package for schema build, that is how I make this....
	// TODO  make this a function call
	// Build descriptior and zip file
	descriptor := map[string]interface{}{
		"resources": []interface{}{
			map[string]interface{}{
				"name":   "datatest1",
				"path":   "./data/data.csv",
				"format": "csv",
				// "profile": "tabular-data-resource",
			},
			map[string]interface{}{
				"name":   "population",
				"path":   "./data/population.csv",
				"format": "csv",
			},
		},
	}

	fmt.Println(len(descriptor))

	//	pkg, err := datapackage.New(descriptor, ".", validator.InMemoryLoader())
	//	if err != nil {
	//		log.Println("in descriptor builder")
	//		log.Println(err)
	//		panic(err)
	//	}

	//	err = pkg.Zip("../../packages/package.zip")
	//	if err != nil {
	//		log.Println("in zip builder")
	//		log.Println(err)
	//		panic(err)
	//	}

	err = os.Chdir("../..")
	if err != nil {
		log.Println("in change dir")
		panic(err)
	}

	// todo  rename package leveraging sha, mv package, delete old tmp directory and contents

}

func makeDescriptor(f []string) ([]byte, error) {
	var vma []map[string]interface{}

	for i := range f {
		vm := make(map[string]interface{})
		vm["name"] = f[i]     // base name only  (might be dups in different sub dirs
		vm["path"] = f[i]     // tmp + data + path
		vm["format"] = "file" //  remove?  replace with something else from spec...
		vma = append(vma, vm)
	}

	descriptor := map[string]interface{}{
		"resources": vma,
	}

	return json.MarshalIndent(descriptor, "", " ")
}

func copyFileContents(src, dst string) (err error) {
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
