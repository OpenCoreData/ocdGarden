package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/karrick/godirwalk"
	"opencoredata.org/ocdGarden/CSDCODirWalkTests/godirwalk/datapackage"
	"opencoredata.org/ocdGarden/CSDCODirWalkTests/godirwalk/graph"
	"opencoredata.org/ocdGarden/CSDCODirWalkTests/godirwalk/kv"
	"opencoredata.org/ocdGarden/CSDCODirWalkTests/godirwalk/report"
)

// TODO: Work Tika flow back in to support discovery
// TODO: Add back in the age of file testing... (out during testing to allow more files to get through)
// TODO: ignore the .DS Store files, the code finds them..

// Examples calls
// godirwalk -dir="foo/bar" -index -report
// godirwalk -dir="foo/bar" -package  (package from index already in KV store)

func main() {
	// Build the output directories we need
	err := os.MkdirAll("./output/kvdata", os.ModePerm)
	err = os.MkdirAll("./output/packages", os.ModePerm)
	if err != nil {
		log.Println("in make dir all")
		panic(err)
	}

	// Set up our log file for runs...
	lf, err := os.OpenFile("output/logfile.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer lf.Close()
	log.SetOutput(lf)

	// command line flags
	dirToIndexPtr := flag.String("dir", "", "directory to index")
	indexPtr := flag.Bool("index", false, "a bool for index building")
	reportPtr := flag.Bool("report", false, "a bool for report building")
	graphPtr := flag.Bool("graph", false, "a bool for graph building")
	packagePtr := flag.Bool("package", false, "a bool for package building")
	resetkvPtr := flag.Bool("resetkv", false, "a bool for reseting the KV store")
	flag.Parse()

	dirname := *dirToIndexPtr

	// reset the KV store  TODO: (make this the default without a flag)
	if *resetkvPtr {
		if _, err := os.Stat("./output/kvdata/index.db"); err == nil {
			err = kv.DeleteKV()
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
			}
		}
	}
	// init the KV store if new, skips if present...
	kv.InitKV()

	// Index the files
	if *indexPtr {
		if dirname == "" {
			fmt.Println("You must provide a directory with -dir DIR/PATH")
			log.Println("You must provide a directory with -dir DIR/PATH")
			os.Exit(1)
		}
		log.Println("Begin index process")

		err = godirwalk.Walk(dirname, &godirwalk.Options{
			Callback: func(osPathname string, de *godirwalk.Dirent) error {
				projDir(de, osPathname, dirname)
				return nil
			},
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
	}

	// Get the index results and work with them
	f := kv.GetEntries() // this is ALL entries.. we need these for the "report" to check whitelist process

	// build csv
	if *reportPtr {
		flen := len(f)

		rows := make([][]string, flen)
		for i := range f {
			rows[i] = []string{f[i].Valid, f[i].ProjName, f[i].File, f[i].Measurement}
			log.Printf("Report line %d / %d \n", i, flen)
		}
		report.CSVReport("report.csv", rows)
	}

	// Get only files that are valid
	w := kv.WhiteList()

	// build graph
	if *graphPtr {
		graph.BuildGraph(w)
	}

	// build packages
	if *packagePtr {
		datapackage.BuildPackage(w, dirname)
	}
}

func projDir(de *godirwalk.Dirent, osPathname, dirname string) {
	pathElements := strings.Split(osPathname, "/")
	argElements := strings.Split(dirname, "/")

	if len(pathElements) > len(argElements) {
		projname := pathElements[len(argElements)]
		if de.IsDir() != true {
			fileIndex(projname, de, osPathname)
		}
	}
}

func fileIndex(projname string, de *godirwalk.Dirent, osPathname string) {
	si := strings.Index(osPathname, projname) + len(projname) // -1 to include the /
	// fmt.Println(ageInYears(osPathname)) // TODO  add back in age check here..  only index if creations time older

	asignPredicate(projname, osPathname[si:])
}

func asignPredicate(projname, osPathname string) {
	// fmt.Printf("Checking %s \n", osPathname)

	// the switch is on the directory name so I need to remove the filename and have only the path left to check on...
	dir, file := filepath.Split(osPathname)

	// Deal with root special
	if dir == "/" {
		// fmt.Printf("File in / : %s \n", file)
		// *mathch* these:  -metadata "metadata format Dtube Label_" "SRF"
		if caseInsenstiveContains(file, "-metadata") {
			kv.NewFileEntry("valid", projname, file, "metadata")
		}
		if caseInsenstiveContains(file, "metadata format Dtube Label_") {
			kv.NewFileEntry("valid", projname, file, "metadata format Dtube Label_")
		}
		if caseInsenstiveContains(file, "SRF") {
			kv.NewFileEntry("valid", projname, file, "SRF")

		}
		fileext := strings.ToLower(filepath.Ext(file))
		s := []string{".cml", ".car"}
		if contains(s, fileext) {
			kv.NewFileEntry("valid", projname, file, ".cml .car")
		}
		return
	}

	switch {
	case caseInsenstiveContains(dir, "Images"):
		fileext := strings.ToLower(filepath.Ext(osPathname))
		s := []string{".bmp", ".jpeg", ".jpg", "tif", "tiff"}
		if contains(s, fileext) {
			kv.NewFileEntry("valid", projname, osPathname, "Images")
		}
	case caseInsenstiveContains(dir, "Images/rgb"):
		fileext := strings.ToLower(filepath.Ext(osPathname))
		s := []string{".csv"}
		if contains(s, fileext) {
			// fmt.Printf("%s: IMAGES/RGB: %s\n", projname, osPathname)
			kv.NewFileEntry("valid", projname, osPathname, "Images/RGB")
		}
	case caseInsenstiveContains(dir, "Geotek Data/whole-core data"):
		fileext := strings.ToLower(filepath.Ext(osPathname))
		blkList := []string{".raw", ".dat", ".out", ".cal"}
		if !contains(blkList, fileext) {
			if caseInsenstiveContains(osPathname, "_HRMS") || caseInsenstiveContains(osPathname, "_XYZ") {
				s := []string{".xls", ".xlsx"}
				if contains(s, fileext) { // TODO  BLACKLIST Needed
					// fmt.Printf("%s: GEOTEK WhCr: %s\n", projname, osPathname)
					kv.NewFileEntry("valid", projname, osPathname, "GEOTEK WhCr")
				}
			}
		}
	case caseInsenstiveContains(dir, "Geotek Data/high-resolution MS data"):
		if caseInsenstiveContains(osPathname, "_HRMS") || caseInsenstiveContains(osPathname, "_XYZ") {
			fileext := strings.ToLower(filepath.Ext(osPathname))
			s := []string{".xls", ".xlsx"}
			if contains(s, fileext) { // TODO  BLACKLIST Needed
				// fmt.Printf("%s: GEOTEK HiRez: %s\n", projname, osPathname)
				kv.NewFileEntry("valid", projname, osPathname, "GEOTEK HiRez")
			}
		}
	case caseInsenstiveContains(dir, "ICD/"):
		fileext := strings.ToLower(filepath.Ext(osPathname))
		s := []string{".pdf"}
		if contains(s, fileext) && !caseInsenstiveContains(file, "ICD sheet.pdf") {
			// fmt.Printf("%s: ICD: %s\n", projname, osPathname)
			kv.NewFileEntry("valid", projname, osPathname, "ICD")
		}
	default:
		// in the root...
		// fmt.Printf("%s: NOT INDEXED:  %s \n", projname, osPathname)
		kv.NewFileEntry("notvalid", projname, osPathname, "")
		//row, _ = report.WriteNotebookRow(row, x, "notvalid", projname, file, "")
	}

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

// TODO:  just a place holder function to remind me abot this, Read into a string array and check for it in array
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
	ctime := time.Unix(int64(stat.Ctim.Sec), int64(stat.Ctim.Nsec))
	delta := time.Now().Sub(ctime)
	years := delta.Hours() / 24 / 365
	fmt.Printf("Create: %v   making it %.2f  years old\n", ctime, years)
	return years

	//return 2.0
}

func caseInsenstiveContains(a, b string) bool {
	return strings.Contains(strings.ToUpper(a), strings.ToUpper(b))
}

func checkPath(path string) error {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return err
		} else {
			return err
		}
	}

	return nil
}
