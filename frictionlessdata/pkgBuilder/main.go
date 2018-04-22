package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/frictionlessdata/datapackage-go/datapackage"
	"github.com/frictionlessdata/datapackage-go/validator"
)

func main() {
	f := make(map[string][]string)
	f["test"] = []string{"./dataVault/data.csv", "./dataVault/population.csv"}
	fmt.Println(f)
	pkgBuilder(f)
}

// TODO integrate with CSDCO walker code
func pkgBuilder(f map[string][]string) {
	fmt.Println("Frictionless Data Package Bulder")
	fmt.Println(f)

	// TODO..  make the following a function and pass it a
	// map like I use in the directory walk program

	// TODO
	// set up temp directory, copy files in, generate zip from that tmp directory
	dir, err := ioutil.TempDir("tmp", "")
	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll(dir) // clean up

	// make data directory inside temp
	os.Mkdir(dir+"/data", os.ModePerm)

	// test cycle through map..  then extend the copy file section
	for k, v := range f {
		fmt.Printf("Name for the package: %s\n", k)
		for i := range v {
			fn := filepath.Base(v[i])
			// copy files
			err = copyFileContents(v[i], dir+"/data/"+fn)
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
	pkg, err := datapackage.New(descriptor, ".", validator.InMemoryLoader())
	if err != nil {
		log.Println("in descriptor builder")
		log.Println(err)
		panic(err)
	}

	err = pkg.Zip("../../packages/package.zip")
	if err != nil {
		log.Println("in zip builder")
		log.Println(err)
		panic(err)
	}

	err = os.Chdir("../..")
	if err != nil {
		log.Println("in change dir")
		panic(err)
	}

	// todo  rename package leveraging sha, mv package, delete old tmp directory and contents

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
