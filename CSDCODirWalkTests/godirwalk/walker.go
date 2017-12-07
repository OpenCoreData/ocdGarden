package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/karrick/godirwalk"
)

func main() {
	dirToIndexPtr := flag.String("dir", ".", "directory to index")

	flag.Parse()
	dirname := *dirToIndexPtr

	err := godirwalk.Walk(dirname, &godirwalk.Options{
		Callback: func(osPathname string, de *godirwalk.Dirent) error {
			projDir(de, osPathname)
			return nil
		},
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func projDir(de *godirwalk.Dirent, osPathname string) {
	pathElements := strings.Split(osPathname, "/")
	if len(pathElements) > 6 {
		projname := pathElements[5]
		if de.IsDir() != true {
			fileIndex(projname, de, osPathname)
		}
	}
}

func fileIndex(projname string, de *godirwalk.Dirent, osPathname string) {

	// pathElements := strings.Split(osPathname, "/")
	// fmt.Println(pathElements[6:(len(pathElements - 1))])

	// fmt.Printf("\n %s %s %s \n", projname, de.Name(), osPathname)
	// Get first index of projname
	// remove to the left of that

	si := strings.Index(osPathname, projname) + len(projname) // -1 to include the /
	// fmt.Println(osPathname[si:])

	asignPredicate(projname, osPathname[si:])

	// getmd5  // use at UUID
	// getUUID

	// metadata
	// makeDC // calls metadata
	// makeSchemaOrg // calls metadata
	// makeFDPackage

	// tikaProcess
	// textIndex   // calls tikaProcess
	// graphIndex
}

func asignPredicate(projname, osPathname string) {

	// fmt.Printf("Checking %s \n", osPathname)

	//  the switch is on the directory name
	//  so I need to remove the filename and have only the
	// path left to check on...

	dir, file := filepath.Split(osPathname)

	// Deal with root special
	if dir == "/" {
		// fmt.Printf("File in / : %s \n", file)
		// *mathch* these:  -metadata "metadata format Dtube Label_" "SRF"
		if caseInsenstiveContains(file, "-metadata") {
			fmt.Printf("%s: ROOT -metadata:  %s \n", projname, file)
		}
		if caseInsenstiveContains(file, "metadata format Dtube Label_") {
			fmt.Printf("%s: ROOT DTUBE:  %s \n", projname, file)
		}
		if caseInsenstiveContains(file, "SRF") {
			fmt.Printf("%s: ROOT SRF:  %s \n", projname, file)
		}
		fileext := strings.ToLower(filepath.Ext(file))
		s := []string{".cml", ".car"}
		if contains(s, fileext) {
			fmt.Printf("%s: ROOT .car or .cml:  %s \n", projname, file)
		}
		return
	}

	switch {
	case caseInsenstiveContains(dir, "Images"):
		fileext := strings.ToLower(filepath.Ext(osPathname))
		s := []string{".bmp", ".jpeg", ".jpg", "tif", "tiff"}
		if contains(s, fileext) {
			fmt.Printf("%s: IMAGES: %s\n", projname, osPathname)
		}
	case caseInsenstiveContains(dir, "Images/rgb"):
		fileext := strings.ToLower(filepath.Ext(osPathname))
		s := []string{".csv"}
		if contains(s, fileext) {
			fmt.Printf("%s: IMAGES/RGB: %s\n", projname, osPathname)
		}
	case caseInsenstiveContains(dir, "Geotek Data/whole-core data"):
		fileext := strings.ToLower(filepath.Ext(osPathname))
		blkList := []string{".raw", ".dat", ".out", ".cal"}
		if !contains(blkList, fileext) {
			if caseInsenstiveContains(osPathname, "_HRMS") || caseInsenstiveContains(osPathname, "_XYZ") {
				s := []string{".xls", ".xlsx"}
				if contains(s, fileext) { // TODO  BLACKLIST Needed
					fmt.Printf("%s: GEOTEK WhCr: %s\n", projname, osPathname)
				}
			}
		}
	case caseInsenstiveContains(dir, "Geotek Data/high-resolution MS data"):
		if caseInsenstiveContains(osPathname, "_HRMS") || caseInsenstiveContains(osPathname, "_XYZ") {
			fileext := strings.ToLower(filepath.Ext(osPathname))
			s := []string{".xls", ".xlsx"}
			if contains(s, fileext) { // TODO  BLACKLIST Needed
				fmt.Printf("%s: GEOTEK HiRez: %s\n", projname, osPathname)
			}
		}
	case caseInsenstiveContains(dir, "ICD/"):
		fileext := strings.ToLower(filepath.Ext(osPathname))
		s := []string{".pdf"}
		if contains(s, fileext) && !caseInsenstiveContains(file, "ICD sheet.pdf") {
			fmt.Printf("%s: ICD: %s\n", projname, osPathname)
		}
	default:
		// in the root...
		fmt.Printf("%s: NOT INDEXED:  %s \n", projname, osPathname)
	}

	// TODO...   assign the predicate and place all resutls in struct
	// Then pretty report print the struct...

	//   match in /
	// "-metadata"  metadata // ns: http://opencoredata.org/id/voc/csdco/v1/
	// "metadata format Dtube Label_"  dtubeMetadata
	// "SRF"  srf
	// ".cml"  cml
	//  ".car" car

	// match in"Geotek Data/whole-core data"
	// BLACK LIST .raw, .dat, .out, .cal ->  NIL

	// "_MSCL"  .xls .xlsx   -> wholeCoreData

	// match in "Geotek Data/high-resolution MS data"
	// _HRMS  .xls .xlsx  -> geotekHighResMSdata
	// _XYZ    .xls  .xlsx -> geotekHighResMSdata

	// match ICD/
	// ICD sheet.pdf ->  icdFiles
	// .pdf  -> icdFiles  (why the above)

	// if .car only do metadata..  no inspection

}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

// Read into a string array and check for it in array
func inApprovedList(projectName string) bool {
	if projectName == "CAHO" {
		return true
	}
	return false
}

// ageInYears gets the age of a file as a float64 decimal value
func ageInYears(fp string) float64 {
	fi, err := os.Stat(fp)
	if err != nil {
		fmt.Println(err)
	}
	stat := fi.Sys().(*syscall.Stat_t)
	ctime := time.Unix(int64(stat.Ctimespec.Sec), int64(stat.Ctimespec.Nsec))
	delta := time.Now().Sub(ctime)
	years := delta.Hours() / 24 / 365
	// fmt.Printf("Create: %v   making it %.2f  years old\n", ctime, years)
	return years
}

func caseInsenstiveContains(a, b string) bool {
	return strings.Contains(strings.ToUpper(a), strings.ToUpper(b))
}
