package main

import (
	"fmt"
	"log"

	"github.com/frictionlessdata/datapackage-go/datapackage"
	"github.com/frictionlessdata/datapackage-go/validator"
)

// TODO integrate with CSDCO walker code
func main() {
	fmt.Println("Frictionless Data Package Bulder")

	// Build descriptior and zip file
	descriptor := map[string]interface{}{
		"resources": []interface{}{
			map[string]interface{}{
				"name": "datatest1",
				// "path": "./data.csv",
				"path": "./data/data.csv",
				//"format": "csv",
				// "profile": "tabular-data-resource",
			},
		},
	}
	pkg, err := datapackage.New(descriptor, ".", validator.InMemoryLoader())
	if err != nil {
		log.Println("in descriptor builder")
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
