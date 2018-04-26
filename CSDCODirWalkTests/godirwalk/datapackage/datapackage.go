package datapackage

import (
	"opencoredata.org/ocdGarden/CSDCODirWalkTests/godirwalk/kv"
)

// BuildPackage should use a pointer ...  :)
func BuildPackage(f []kv.FileMeta, dirname string) {

	// Need to set the package name
	// Need to set the package ID
	// How will the content URL be define?

	prjs := uniqueProjects(f)

	// Loop on each unique project, find all the files for that project
	// TODO  make a map to  map[projs]projFies   map[string][]string
	pf := make(map[string][]string)
	for p := range prjs {
		uf := projFiles(f, prjs[p])
		pf[prjs[p]] = uf
		// log.Printf("K: %s  V: %s \n", prjs[p], uf)
	}

	// log.Print(len(pf))

	packagedir := "/media/fils/seagate/packages"
	vaultdir := dirname
	tempdir := "/media/fils/seagate/tmp"

	PKGBuilder(pf, vaultdir, tempdir, packagedir)

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
