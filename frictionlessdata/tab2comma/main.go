package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Conert tab to comma delimated csv")

	tabdata := csvAsComma("../testdata/208_1262C_JanusGraSection_tALVlUMg.csv")

	csvFile, err := os.Create("./csvComma.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()
	commaWriter := newWriter(csvFile)
	commaWriter.WriteAll(tabdata)
}

func csvAsComma(filename string) [][]string {
	csvFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.Comma = '\t' // Use tab-delimited instead of comma <---- here!
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return csvData
}

func newWriter(w io.Writer) (writer *csv.Writer) {
	writer = csv.NewWriter(w)
	return
}
