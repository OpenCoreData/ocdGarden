package datapackage

import (
	"crypto/sha256"
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

// PKGBuilder integrate with CSDCO walker code
func PKGBuilder(f map[string][]string, vaultdir, tempdir, packagedir string) {
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
			//err = copyFileContents(fqp, pdd+"/"+fn)
			err = copyBySymLink(fqp, pdd+"/"+fn)

			fmt.Printf("------>>>>>>>   %s\n", strings.TrimPrefix(pdd+"/"+fn, dir+"/"+cleanstring(k)+"/"))
			fa = append(fa, strings.TrimPrefix(pdd+"/"+fn, dir+"/"+cleanstring(k)+"/"))

			if err != nil {
				log.Println("in copy file")
				panic(err)
			}
		}
		pm[k] = fa
	}

	// At this point the fles are copied into the project data directory..   can I calcualte a hash now?
	// shaval = calcSha(directory)   // takes a directory, reads all files and and generates the final sha value..
	// be sure to use io.Writer for this...

	// loop on pm, build a schema.org, put it into the temp dir, add it's path to the value array....
	m := make(map[string]string)
	for a, _ := range pm {

		projdir := fmt.Sprintf("%s/%s", dir, cleanstring(a))
		// // change working directory
		fmt.Printf("The projdir is %s\n", projdir)
		err = os.Chdir(projdir)
		log.Println(dir)
		if err != nil {
			log.Println("in change dir")
			panic(err)
		}

		err = os.MkdirAll("metadata", os.ModePerm)
		if err != nil {
			log.Println("in make dir all")
			panic(err)
		}

		sv := shaDataDir("./data")
		m[cleanstring(a)] = sv // TODO..  add this a map of a[sv] to use later for naming the file..

		// make a file..
		d1 := []byte(BuildSchema(dirProjName(a), a, sv))
		err = ioutil.WriteFile("./metadata/schemaorg.json", d1, 0644)
		if err != nil {
			log.Println("write the scheamaorg.json file")
			panic(err)
		}

		// make a metadata dir in the project dir
		// write jld to the tmp/proj/metadata directory
		// append tmp/proj/metadata/schemaorg.json to the file array at b  (the _ value for now)

	}

	// loop on the new map from above and move in and generate the datapackage json and zip packages...
	for i, j := range pm {
		fmt.Println("-------------  loop start  ---------------------------")
		fmt.Println(i)
		fmt.Println(j)

		i = cleanstring(i) // really want i to be some a hash or PID

		// Append in to the j array the presence of the schemaorg.json file
		j = append(j, "./metadata/schemaorg.json")

		projdir := fmt.Sprintf("%s/%s", dir, i)
		// // change working directory
		fmt.Printf("The projdir is %s \n", projdir)
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

		// TODO..  now change the name to be the shavalue .zip
		fmt.Printf("%s     to     %s/%s.zip ", zipfp, packagedir, m[i])
		err = os.Rename(zipfp, fmt.Sprintf("%s/%s.zip", packagedir, m[i]))
		if err != nil {
			fmt.Println(err)
			// os.Exit(1)
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

// This is a sym link..  not sure if the FDP package will work with sym links
// If we can stay on the same partiion/fs then hard links are an option.
func copyBySymLink(src, dst string) (err error) {
	fmt.Printf("Sym Link %s to %s\n", src, dst)
	os.Symlink(src, dst)
	return nil
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

func dirProjName(s string) string {
	sa := strings.Split(s, " ")
	return sa[0]
}

func shaDataDir(s string) string {

	h := sha256.New()

	fileList := make([]string, 0)
	e := filepath.Walk(s, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			fileList = append(fileList, path)
			fmt.Printf("In Sha with a file %s \n", path)

			f, err := os.Open(path)
			if err != nil {
				log.Print(err)
			}
			defer f.Close()

			if _, err := io.Copy(h, f); err != nil {
				log.Print(err)
			}
		}
		return err
	})

	if e != nil {
		log.Print(e) // I should make this a bit more fatal?  or no reason to "panic"  ;)
	}

	shavalue := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(shavalue)

	return shavalue
}
