package report

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"regexp"

	"../vault/"
)

// CSVReport genreate the CSV version of the report
func CSVReport(name string, vh vault.VaultHoldings) {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)

	}
	s := reg.ReplaceAllString(name, "")

	f, err := os.Create(fmt.Sprintf("./output/report_%s.csv", s))
	if err != nil {
		log.Fatalf("Cannot open '%s': %s\n", name, err.Error())
	}
	defer func() {
		e := f.Close()
		if e != nil {
			log.Fatalf("Cannot close '%s': %s\n", name, e.Error())
		}
	}()

	var rows [][]string

	// comment out to remove headers
	headers := []string{"Project", "Type", "FileName", "FileExt", "Parent Directory", "Relative Path", "Voc URI"}
	rows = append(rows, headers)

	for _, item := range vh.Holdings {
		// if item.Type != "Unknown" && item.Type != "Directory" {
		if item.Type != "Directory" {
			// sa := []string{strconv.FormatBool(item.Public),
			sa := []string{item.Project,
				item.Type,
				item.FileName,
				item.FileExt,
				item.ParentDir,
				item.RelativePath,
				item.TypeURI}
			rows = append(rows, sa)
		}
	}

	w := csv.NewWriter(f)
	err = w.WriteAll(rows)
}
