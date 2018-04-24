package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/karrick/godirwalk"
	"opencoredata.org/ocdGarden/CSDCODirWalkTests/godirwalk/datapackage"
	"opencoredata.org/ocdGarden/CSDCODirWalkTests/godirwalk/graph"
	"opencoredata.org/ocdGarden/CSDCODirWalkTests/godirwalk/kv"
	"opencoredata.org/ocdGarden/CSDCODirWalkTests/godirwalk/report"
)

// TODO:  work old Tika path in somehow?
// TODO: Add back in the age of file testing...

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

	// flags
	dirToIndexPtr := flag.String("dir", "", "directory to index")
	indexPtr := flag.Bool("index", false, "a bool for index build")
	reportPtr := flag.Bool("report", false, "a bool for report build")
	graphPtr := flag.Bool("graph", false, "a bool for graph building")
	packagePtr := flag.Bool("package", false, "a bool for package building")
	resetkvPtr := flag.Bool("resetkv", false, "a bool for reseting the KV store")
	flag.Parse()

	dirname := *dirToIndexPtr

	if dirname == "" || !*indexPtr {
		fmt.Println("You must provide -index flag and a directory with -dir DIR/PATH")
		log.Println("You must provide -index flag and a directory with -dir DIR/PATH")
		os.Exit(1)
	}

	if *resetkvPtr { // TODO:  invert this so the default is to reset the KV store
		if _, err := os.Stat("./output/kvdata/index.db"); err == nil {
			err = kv.DeleteKV()
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
			}
		}

	}

	kv.InitKV()

	// Index the files  // TODO  make this take a flag
	if *indexPtr {
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
	f := kv.GetEntries()

	// build excel
	if *reportPtr {
		x := report.InitNotebook()
		for i := range f {
			_, _ = report.WriteNotebookRow(i+1, x, f[i].Valid, f[i].ProjName, f[i].File, f[i].Measurement)
		}
		report.SaveNotebook(x)
	}

	// build graph
	if *graphPtr {
		graph.BuildGraph(f)
	}

	// build packages
	if *packagePtr {
		datapackage.BuildPackage(f)
	}
}

func projDir(de *godirwalk.Dirent, osPathname, dirname string) {
	pathElements := strings.Split(osPathname, "/")
	// fmt.Println(pathElements)
	argElements := strings.Split(dirname, "/")
	// fmt.Println(argElements)

	if len(pathElements) > len(argElements) {
		projname := pathElements[len(argElements)]
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
}

func asignPredicate(projname, osPathname string) {
	// fmt.Printf("Checking %s \n", osPathname)

	//  the switch is on the directory name
	//  so I need to remove the filename and have only the
	// path left to check on...

	dir, file := filepath.Split(osPathname)
	//x := report.InitNotebook()
	//row := 1

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

	// TODO...   assign the predicate and place all results in struct
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
	// fi, err := os.Stat(fp)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// stat := fi.Sys().(*syscall.Stat_t)
	// ctime := time.Unix(int64(stat.Ctimespec.Sec), int64(stat.Ctimespec.Nsec))
	// delta := time.Now() //.Sub(ctime)
	// years := delta.Hours() / 24 / 365
	// // fmt.Printf("Create: %v   making it %.2f  years old\n", ctime, years)
	// return years

	return 2.0
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
