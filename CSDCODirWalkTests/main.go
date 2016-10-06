package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("CSDCO Directory walk filter testing")
	flag.Parse()
	root := flag.Arg(0)
	filepath.Walk(root, VisitFile)
}

// VisitFile from https://rosettacode.org/wiki/Walk_a_directory/Recursively#Go
func VisitFile(fp string, fi os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err) // can't walk here,
		return nil       // but continue walking elsewhere
	}
	if fi.IsDir() {
		return nil // not a file.  ignore.
	}

	// Look in all subdirectories of a project directory
	if caseInsenstiveContains(fp, "/") {

		// Looking for [ProjectID]-metadata
		// We really don't know the PID..  need to try and pull that from the path....
		matched, err := filepath.Match(strings.ToLower("*-metadata*"), strings.ToLower(fi.Name()))

		// look for Dtube lable name...
		if !matched {
			matched, err = filepath.Match(strings.ToLower("*metadata format Dtube Lable_*"), strings.ToLower(fi.Name())) // worry about case issue
		}

		// subsample metadata information
		if !matched {
			matched, err = filepath.Match(strings.ToLower("*SRF*"), strings.ToLower(fi.Name()))
		}

		// Corelyzer session file
		if !matched {
			matched, err = filepath.Match(strings.ToLower("*.cml"), strings.ToLower(fi.Name()))
		}

		// Corelyzer archive file
		if !matched {
			matched, err = filepath.Match(strings.ToLower("*.car"), strings.ToLower(fi.Name()))
		}
		if err != nil {
			fmt.Println(err) // malformed pattern
			return err       // this is fatal.
		}
		if matched {
			fmt.Printf("Match for (PID/Dtube/SRF/.cml/.car): %s\n", fp)
		}
	}

	// Walk all subdirectories?
	if caseInsenstiveContains(fp, "Images/") {
		matched, err := filepath.Match(strings.ToLower("*.jpg"), strings.ToLower(fi.Name()))

		if !matched {
			matched, err = filepath.Match(strings.ToLower("*.jpeg"), strings.ToLower(fi.Name()))
		}
		if !matched {
			matched, err = filepath.Match(strings.ToLower("*.tif"), strings.ToLower(fi.Name()))
		}
		if !matched {
			matched, err = filepath.Match(strings.ToLower("*.tiff"), strings.ToLower(fi.Name()))
		}
		if !matched {
			matched, err = filepath.Match(strings.ToLower("*.bmp"), strings.ToLower(fi.Name()))
		}

		if err != nil {
			fmt.Println(err) // malformed pattern
			return err       // this is fatal.
		}
		if matched {
			fmt.Printf("Image directory (jpg/jpeg/tif/bmp): %s\n", fp)
		}
	}

	// Walk all subdirectories?
	if caseInsenstiveContains(fp, "Images/rgb") {
		matched, err := filepath.Match(strings.ToLower("*.csv"), strings.ToLower(fi.Name()))
		if err != nil {
			fmt.Println(err) // malformed pattern
			return err       // this is fatal.
		}
		if matched {
			fmt.Printf("Image/rgb data in CSV: %s\n", fp)
		}
	}

	if caseInsenstiveContains(fp, "Geotek Data/whole-core data") {

		// black list this extensions in here: .raw .dat .out and .cal
		matched, err := filepath.Match(strings.ToLower("*.raw"), strings.ToLower(fi.Name()))
		if !matched {
			matched, err = filepath.Match(strings.ToLower("*.dat"), strings.ToLower(fi.Name()))
		}
		if !matched {
			matched, err = filepath.Match(strings.ToLower("*.out"), strings.ToLower(fi.Name()))
		}
		if !matched {
			matched, err = filepath.Match(strings.ToLower("*.cal"), strings.ToLower(fi.Name()))
		}

		if err != nil {
			fmt.Println(err) // malformed pattern
			return err       // this is fatal.
		}
		if matched {
			return nil // We return nil on match here since we got a postive from the Black list above
		}

		// if we don't drop out in the above black list, check for our white list pattern
		matched, err = filepath.Match(strings.ToLower("*_MSCL*"), strings.ToLower(fi.Name()))
		if matched {
			// now check for correct extensions
			matched, err = filepath.Match(strings.ToLower("*.xsl"), strings.ToLower(fi.Name()))
			if !matched {
				matched, err = filepath.Match(strings.ToLower("*.xsls"), strings.ToLower(fi.Name()))
			}
			if !matched {
				return nil // done with this test loop..
			}
			if matched {
				fmt.Printf("GeoTek Whole Core: %s\n", fp)
			}
			if err != nil {
				fmt.Println(err) // malformed pattern
				return err       // this is fatal.
			}
		}
	}

	if caseInsenstiveContains(fp, "Geotek Data/high-resolution MS data") {
		matched, err := filepath.Match(strings.ToLower("*_HRMS*"), strings.ToLower(fi.Name()))

		if !matched {
			matched, err = filepath.Match(strings.ToLower("*_XYZ*"), strings.ToLower(fi.Name()))
		}

		if err != nil {
			fmt.Println(err) // malformed pattern
			return err       // this is fatal.
		}
		if matched {
			// now check for correct extensions  //TODO  ask if I should add .csv here as well..

			matched, err = filepath.Match(strings.ToLower("*.xls"), strings.ToLower(fi.Name()))
			if !matched {
				matched, err = filepath.Match(strings.ToLower("*.xlsx"), strings.ToLower(fi.Name()))
			}
			if !matched {
				return nil // done with this test loop..
			}
			if matched {
				fmt.Printf("GeoTek High Res: %s\n", fp)
			}
			if err != nil {
				fmt.Println(err) // malformed pattern
				return err       // this is fatal.
			}
		}

	}

	// Walk all subdirectories?
	if caseInsenstiveContains(fp, "ICD/") {
		matched, err := filepath.Match("ICD sheet.pdf", strings.ToLower(fi.Name()))
		if matched {
			return nil // we matched above so get out now...
		}

		matched, err = filepath.Match(strings.ToLower("*.pdf"), strings.ToLower(fi.Name()))

		if err != nil {
			fmt.Println(err) // malformed pattern
			return err       // this is fatal.
		}
		if matched {
			fmt.Printf("ICD files: %s\n", fp)
		}
	}

	return nil
}

func dirFunc(path string) (int, error) {
	var count int
	var FQP string

	err := filepath.Walk(path, func(_ string, f os.FileInfo, err error) error {

		FQP = fmt.Sprintf("%s/%s", path, f.Name())

		fmt.Println(f.Name())

		if !f.IsDir() {
			count = count + 1

			dir, file := filepath.Split(FQP)

			fmt.Printf("For path:\t%s \nFor dir:\t%s\nFor file:\t%s \n\n", FQP, dir, file)

			// set up our cases here to check (need an int to switch on though)..  based on these different actions will be done on the files
			// and different "types" will be set in the RDF graph for the file types

			// look for "Geotek Data"
			if caseInsenstiveContains(FQP, "whole-core data") {
				if caseInsenstiveContains(file, "_MSCL") {
					fmt.Printf("Geotek WholeCore data\nFor path:\t%s \nFor dir:\t%s\nFor file:\t%s \n\n", FQP, dir, file)
				}
			}

			// look for "ICD Data"  (add looking for .pdf on file name)
			if caseInsenstiveContains(FQP, "/ICD") { // use / to match to a directory...  avoid hitting "icd" in a directory name string
				fmt.Printf("ICD data\nFor path:\t%s \nFor dir:\t%s\nFor file:\t%s \n\n", FQP, dir, file)
			}

			// look for "CSD Data"
			// if caseInsenstiveContains(dir, "CSD Data") {
			// 	fmt.Printf("CSD data\nFor path:\t%s \nFor dir:\t%s\nFor file:\t%s \n\n", FQP, dir, file)
			// }

			// fmt.Printf("For path:\t%s \nFor dir:\t%s\nFor file:\t%s \n\n", FQP, dir, file)
		}
		return err
	})

	return count, err
}

func caseInsenstiveContains(a, b string) bool {
	return strings.Contains(strings.ToUpper(a), strings.ToUpper(b))
}
