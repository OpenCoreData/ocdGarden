package index

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"../heuristics"
	"../vault"
)

// PathInspection function
func PathInspection(d, f string) vault.VaultItem {
	var proj, rel, dir, file string

	// or is dir?   might not work since I have strings
	if filepath.Ext(f) != "" {
		r, err := filepath.Rel(d, f) // TODO   this error is important..    deal with it!!!!
		if err != nil {
			log.Printf("Error in filepath Rel: %s\n", err)
		}
		rel = r
		//dir = filepath.Dir(r)
		dir, file = filepath.Split(r) // if done off rel, would dir work here?
	}

	pathElements := strings.Split(f, "/")
	argElements := strings.Split(d, "/")
	if len(pathElements) > len(argElements) {
		proj = pathElements[len(argElements)]
	} else {
		proj = "/"
	}

	fe := filepath.Ext(f)

	t, uri, err := fileType(f)
	if err != nil {
		// v := vault.VaultItem{Name: f, Type: nil, Public: false, Project: p}
		log.Println(err)
	}

	// TODO  Public is just set true.  so obviously meaningless..
	//  It's a place holder in case we need moratorium flags
	// Type == Unknown is really the "do no index flag"...   this is confusing..  need to resolve this in the code
	v := vault.VaultItem{Name: f, Type: t, Public: true, Project: proj,
		RelativePath: rel, FileName: file, ParentDir: dir, FileExt: fe, TypeURI: uri}
	return v
}

func fileType(f string) (string, string, error) {
	// if directory.. note and get out
	// CAUTION: by not calling "open" on the file to "lock" it , this
	// code risks the type changing during the code operation....  the
	// odds of this are VERY low...
	fi, err := os.Stat(f)
	if err != nil {
		panic(err) // don't panic..   deal with this..
	}
	if fi.IsDir() {
		return "Directory", "", nil
	}

	t := "Unknown" // by default..  we don't know what the file is  (could also return an error type for this)
	uri := ""
	tests := heuristics.CSDCOHTs()
	dir, _ := filepath.Split(f)

	for i := range tests {
		// fmt.Println("Starting a file test")
		// fmt.Printf("Test: %s %s \n", tests[i].FileExts, tests[i].Comment)
		if caselessContains(dir, tests[i].DirPattern) {
			if caselessContainsSlice(f, tests[i].FilePattern) {
				fileext := strings.ToLower(filepath.Ext(f))
				s := tests[i].FileExts
				if contains(s, fileext) {
					fmt.Printf("%s == %s\n", f, tests[i].Comment) //  TODO  all NewFileEntry calls should use class URI, not name like "Images"
					t = tests[i].Comment
					uri = tests[i].URI
				}
			}
		}
	}

	return t, uri, err
}

func caselessContainsSlice(a string, b []string) bool {
	t := true // default to true so that 0 len string array is NOT a test.
	for i := range b {
		t = strings.Contains(strings.ToUpper(a), strings.ToUpper(b[i]))
		// fmt.Printf("Tested %s against %s and got %t\n", a, b[i], t)
	}

	return t
}

func caselessContains(a, b string) bool {
	return strings.Contains(strings.ToUpper(a), strings.ToUpper(b))
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}
