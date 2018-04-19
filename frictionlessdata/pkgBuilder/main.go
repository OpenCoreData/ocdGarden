package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/frictionlessdata/datapackage-go/datapackage"
	"github.com/frictionlessdata/datapackage-go/validator"
)

// TODO integrate with CSDCO walker code
func main() {
	fmt.Println("Frictionless Data Package Bulder")

	// TODO
	// set up temp directory, copy files in, generate zip from that tmp directory
	dir, err := ioutil.TempDir("tmp", "")
	if err != nil {
		log.Fatal(err)
	}

	//defer os.RemoveAll(dir) // clean up

	// make data directory inside temp
	os.Mkdir(dir+"/data", os.ModePerm)

	// copy files
	err = copyFileContents("./dataVault/data.csv", dir+"/data/data.csv")
	if err != nil {
		log.Println("in copy file")
		panic(err)
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
		},
	}
	pkg, err := datapackage.New(descriptor, ".", validator.InMemoryLoader())
	if err != nil {
		log.Println("in desciptor builder")
		log.Println(err)
		panic(err)
	}

	err = pkg.Zip("package.zip")
	if err != nil {
		log.Println("in zip builder")
		log.Println(err)
		panic(err)
	}

}

func copyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}
