package main

import (
	"fmt"
	"io/ioutil"

	"github.com/frictionlessdata/datapackage-go/datapackage"
)

func main() {
	fmt.Println("Raw Read Test")

	//pkg, err := datapackage.Load("data/datapackage.json")
	// pkg, err := datapackage.Load("https://raw.githubusercontent.com/frictionlessdata/example-data-packages/master/simple-geojson/datapackage.json")
	pkg, err := datapackage.Load("http://raw.githubusercontent.com/frictionlessdata/example-data-packages/master/simple-geojson/datapackage.json")
	if err != nil {
		fmt.Println("Load error")
		fmt.Println(err)
	}

	so := pkg.GetResource("example")
	rc, err := so.RawRead()
	if err != nil {
		fmt.Println("rawread error")
		fmt.Println(err)
	}

	defer rc.Close()
	contents, err := ioutil.ReadAll(rc)
	if err != nil {
		fmt.Println("ioutil error")
		fmt.Println(err)
	}

	fmt.Println(string(contents))
}
