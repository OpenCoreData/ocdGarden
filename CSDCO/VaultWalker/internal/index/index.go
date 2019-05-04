package index

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"opencoredata.org/ocdGarden/CSDCO/VaultWalker/internal/heuristics"
	"opencoredata.org/ocdGarden/CSDCO/VaultWalker/internal/vault"
)

// PathInspection function
func PathInspection(d, f string) (vault.VaultItem, error) {
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
	// fmt.Printf("Proj test: %v, %v \n", pathElements, argElements)
	if len(pathElements) > len(argElements) {
		proj = pathElements[len(argElements)]
	} else {
		proj = "/"
	}

	// don't index the / of the request
	// TODO  Add the black list check here leveraging caselessContainsSlice
	if proj == "/" {
		// fmt.Println("Proj test root: in the root dir")
		return vault.VaultItem{}, errors.New("Do not index the root directory of projects")
	}

	// don't index files in the / of the request
	fi, err := os.Stat(fmt.Sprintf("%s/%s", d, proj))
	if err != nil {
		// fmt.Printf("Proj test error: %s \n", err)
		return vault.VaultItem{}, err
	}
	if !fi.IsDir() {
		// fmt.Printf("Proj test isDir: %s is a file \n", f)
		return vault.VaultItem{}, errors.New("Do not index the root directory files")
	}

	// don't index directories with given prefix
	if strings.HasPrefix(proj, "!") {
		// fmt.Println("Proj test _: Project marked for skipping")
		return vault.VaultItem{}, errors.New("Do not index projects that start with _")
	}

	// get extension
	fe := filepath.Ext(f)

	// Typing..  look for type and info about f in directory d for  project proj
	t, uri, err := fileType(d, proj, f)
	if err != nil {
		// v := vault.VaultItem{Name: f, Type: nil, Public: false, Project: p}
		log.Println(err)
	}

	a, c := ageInYears(f)

	// TODO  Public is just set true.  so obviously meaningless..
	//  It's a place holder in case we need moratorium flags
	// Type == Unknown is really the "do no index flag"...   this is confusing..  need to resolve this in the code
	v := vault.VaultItem{Name: f, Type: t, Public: true, Project: proj,
		RelativePath: rel, FileName: file, ParentDir: dir, FileExt: fe, TypeURI: uri, Age: a, DateCreated: c.Format("2006-01-02")}

	return v, nil // proj == / then don't return it though..   need to return error since I can't return a nil struct
}

func fileType(d, proj, f string) (string, string, error) {
	// if directory.. note and get out
	// CAUTION: by not calling "open" on the file to "lock" it , this
	// code risks the type changing during the code operation....  the
	// odds of this are VERY VERY low however...
	fi, err := os.Stat(f)
	if err != nil {
		panic(err) // don't panic..   deal with this..
	}
	if fi.IsDir() {
		return "Directory", "", nil
	}

	// print the mod time stamp
	// fmt.Println(fi.ModTime().String())
	// fmt.Println(ageInYears(f)) // sys call function for age to get creation time (windows cifs only...  linux fs will not store)

	t := "Exclude" // by default..  we don't know what the file is  (could also return an error type for this)
	uri := ""
	tests := heuristics.CSDCOHTs()
	dir, _ := filepath.Split(f)

	for i := range tests {
		if caselessPrefix(d, proj, dir, tests[i].DirPattern) {
			// if caselessContains(dir, tests[i].DirPattern) { // TODO should become caselessPrefix(d, proj, dir, tests[i].DirPattern)
			if fileInDir(d, proj, tests[i].DirPattern, f) {
				if caselessContainsSlice(f, tests[i].FilePattern) {
					fileext := strings.ToLower(filepath.Ext(f))
					s := tests[i].FileExts
					if contains(s, fileext) {
						// fmt.Printf("%s == %s\n", f, tests[i].Comment) //  TODO  all NewFileEntry calls should use class URI, not name like "Images"
						t = tests[i].Comment
						uri = tests[i].URI
					}
				}
			}
		}
	}

	return t, uri, err
}

func fileInDir(d, proj, dp, f string) bool {
	a := fmt.Sprintf("%s/%s/%s", d, proj, dp)
	b := fmt.Sprintf("%s/", filepath.Dir(f))

	i := strings.Compare(a, b)
	e := false
	if i == 0 {
		e = true
	}

	// fmt.Printf("%t :: fileInDir: %s is in %s, %b \n", e, f, a, b)

	return e
}

func caselessContainsSlice(a string, b []string) bool {
	t := true // default to true so that 0 len string array is NOT a test.
	for i := range b {
		t = strings.Contains(strings.ToUpper(a), strings.ToUpper(b[i]))
		// 	fmt.Printf("CCS Tested %s against %s and got %t\n", a, b[i], t)
	}
	// fmt.Printf("CSS called, returning %t for %s \n", t, a)
	return t
}

func caselessContains(a, b string) bool {
	return strings.Contains(strings.ToUpper(a), strings.ToUpper(b))
}

// Test if a has prefix b    /dir1/dir2/dir3/filex  has /dir1/dir2
// To do this I need the base directory to remove from a
func caselessPrefix(base, proj, a, b string) bool {
	pref := fmt.Sprintf("%s/%s/", base, proj)
	// fmt.Printf("Directory Test: %s\n", pref)
	// fmt.Printf("Directory Test: %s\n", a)
	atl := strings.TrimPrefix(a, pref)
	// fmt.Printf("Directory Test: %s\n", atl)
	// fmt.Printf("Directory Test: In base %s in proj %s test if  %s has prefix %s result: %t \n", base, proj, strings.ToUpper(atl), strings.ToUpper(b),
	// strings.HasPrefix(strings.ToUpper(atl), strings.ToUpper(b)))
	return strings.HasPrefix(strings.ToUpper(atl), strings.ToUpper(b))
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

// ageInYears gets the age of a file as a float64 decimal value
func ageInYears(fp string) (float64, time.Time) {
	fi, err := os.Stat(fp)
	if err != nil {
		fmt.Println(err)
	}
	stat := fi.Sys().(*syscall.Stat_t)
	// ctime := time.Unix(int64(stat.Ctim.Sec), int64(stat.Ctim.Nsec))
	ctime := time.Unix(int64(stat.Mtim.Sec), int64(stat.Mtim.Nsec))
	delta := time.Now().Sub(ctime)
	years := delta.Hours() / 24 / 365
	// fmt.Printf("Create: %v   making it %.2f  years old\n", ctime, years)
	return round2(years, 0.01), ctime
}

func round2(x, unit float64) float64 {
	if x > 0 {
		return float64(int64(x/unit+0.5)) * unit
	}
	return float64(int64(x/unit-0.5)) * unit
}
