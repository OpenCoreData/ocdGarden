package build

import (
	"fmt"

	"github.com/frictionlessdata/datapackage-go/datapackage"
	"github.com/frictionlessdata/datapackage-go/validator"
)

func Pkg() error {

	fmt.Println("In package builder")
	// make temp directory
	// load in data and metadata
	// build and copy in descriptor
	// call pkg builders  https://github.com/frictionlessdata/datapackage-go

	descriptor := map[string]interface{}{
		"resources": []interface{}{
			map[string]interface{}{
				"name":    "books",
				"path":    "books.csv",
				"format":  "csv",
				"profile": "tabular-data-resource",
				"schema": map[string]interface{}{
					"fields": []interface{}{
						map[string]interface{}{"name": "author", "type": "string"},
						map[string]interface{}{"name": "title", "type": "string"},
						map[string]interface{}{"name": "year", "type": "integer"},
					},
				},
			},
		},
	}

	_, err := datapackage.New(descriptor, ".", validator.InMemoryLoader())
	if err != nil {
		panic(err)
	}

	return nil
}

// // func makeDescriptor(f []string) ([]byte, error) {
// func makeDescriptor(f []string) (map[string]interface{}, error) {
// 	var vma []interface{} //  was []map[string]interface{}  //

// 	for i := range f {
// 		vm := make(map[string]interface{})
// 		vm["name"] = filepath.Base(f[i]) // base name only  (might be dups in different sub dirs?
// 		path := f[i]
// 		// remove leading / if it exist...  ???  added later..  check for a regression issue
// 		if strings.HasPrefix(path, "data//") {
// 			path = strings.Replace(path, "data//", "data/", 1)
// 		}
// 		vm["path"] = path // tmp + data + path  (need to strip any ./ or / at the start..)
// 		// v["format"] = "file" //  remove?  replace with something else from spec...
// 		vma = append(vma, vm)
// 	}

// 	descriptor := map[string]interface{}{
// 		"resources": vma,
// 	}

// 	return descriptor, nil
// }

func pkgToS3() error {
	return nil
}

func pkgToDisk() error {
	return nil
}
