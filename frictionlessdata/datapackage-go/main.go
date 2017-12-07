package main

import (
	"fmt"
	"log"

	"github.com/frictionlessdata/datapackage-go/datapackage"
)

func main() {
	fmt.Println("Frictionless Data Package")

	pkg, err := datapackage.Load("./datapackage.json")
	// pkg, err := datapackage.Load("./population.json")
	if err != nil {
		log.Printf("Error in Load: %s", err)
	}

	resource := pkg.GetResource("boring") // needs a better error
	// resource := pkg.GetResource("population") // needs a better error
	contents, err := resource.ReadAll()
	if err != nil {
		log.Printf("Error in ReadAll: %s", err)
	}
	fmt.Println(contents)

}
