package datapackage

import (
	"fmt"
	"opencoredata.org/ocdGarden/CSDCODirWalkTests/godirwalk/kv"
)

// BuildPakcage should use a pointer ...  :)
func BuildPackage(f []kv.FileMeta) {
	fmt.Println("package builder")

	prjs := uniqueProjects(f)

	// Loop on each unique project, find all the files for that project
	// TODO  make a map to  map[projs]projFies   map[string][]string
	for p := range prjs {
		fmt.Println(prjs[p])
		uf := projFiles(f, prjs[p])
		fmt.Println(uf)
	}

	// TODO..
	// then for each one pull the unique fles..
	// gather the files,
	// build the manafest
	// build the metadata
	// assemble the package

}

func projFiles(f []kv.FileMeta, prj string) []string {
	var files []string
	m := map[string]bool{}

	for _, v := range f {
		if v.ProjName == prj { // todo..  use strings compare
			if !m[v.File] {
				m[v.File] = true
				files = append(files, v.File)
			}
		}
	}

	return files
}

func uniqueProjects(f []kv.FileMeta) []string {
	var unique []string
	m := map[string]bool{}

	for _, v := range f {
		if !m[v.ProjName] {
			m[v.ProjName] = true
			unique = append(unique, v.ProjName)
		}
	}

	return unique
}
