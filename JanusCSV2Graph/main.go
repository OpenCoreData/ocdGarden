package main

import (
	// "bytes"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	// "text/template"

	rdf "github.com/knakk/rdf"
)

type HDR struct {
	Table_name      string
	Struct_name     string
	Struct_descript string
	Query_descript  string
}

type JSONSKOS struct {
	Json_name          string
	Json_unit          string
	Json_unit_descript string
	Json_descript      string
}

type Voc struct {
	Struct_name        string
	Table_name         string
	Column_name        string
	Column_id          string
	Go_struct_name     string
	Go_struct_type     string
	Json_name          string
	Code               string
	Xs_type            string
	Json_unit          string
	Json_unit_descript string
	Json_descript      string
}

func main() {
	csvdata := readMetaData()
	hdrdata := readHDRData()
	jsdata := readJSONSKOSData()

	// RDF item
	tr := []rdf.Triple{}

	for _, item := range hdrdata {
		newsub, _ := rdf.NewIRI(fmt.Sprintf("http://opencoredata/id/resource/janus/query/%s", item.Table_name)) // Sprintf a correct URI here

		newpred0, _ := rdf.NewIRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type")
		newobj0, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/JanusQuery")
		newtriple0 := rdf.Triple{Subj: newsub, Pred: newpred0, Obj: newobj0}

		newpred1, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/json_name")
		newpred2, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/json_unit")
		newpred3, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/json_unit_descript")
		newpred4, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/json_descript")

		newobj1, _ := rdf.NewLiteral(item.Table_name)
		newobj2, _ := rdf.NewLiteral(item.Struct_name)
		newobj3, _ := rdf.NewLiteral(item.Struct_descript)
		newobj4, _ := rdf.NewLiteral(item.Query_descript)

		newtriple1 := rdf.Triple{Subj: newsub, Pred: newpred1, Obj: newobj1}
		newtriple2 := rdf.Triple{Subj: newsub, Pred: newpred2, Obj: newobj2}
		newtriple3 := rdf.Triple{Subj: newsub, Pred: newpred3, Obj: newobj3}
		newtriple4 := rdf.Triple{Subj: newsub, Pred: newpred4, Obj: newobj4}

		tr = append(tr, newtriple0)
		if newtriple1.Obj.String() != "" {
			tr = append(tr, newtriple1)
		}
		if newtriple2.Obj.String() != "" {
			tr = append(tr, newtriple2)
		}
		if newtriple3.Obj.String() != "" {
			tr = append(tr, newtriple3)
		}
		if newtriple4.Obj.String() != "" {
			tr = append(tr, newtriple4)
		}
	}

	for _, item := range jsdata {
		newsub, _ := rdf.NewIRI(fmt.Sprintf("http://opencoredata/id/resource/janus/measure/%s", item.Json_name)) // Sprintf a correct URI here

		newpred0, _ := rdf.NewIRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type")
		newobj0, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/JanusMeasurement")
		newtriple0 := rdf.Triple{Subj: newsub, Pred: newpred0, Obj: newobj0}

		newpred1, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/json_name")
		newpred2, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/json_unit")
		newpred3, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/json_unit_descript")
		newpred4, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/json_descript")

		newobj1, _ := rdf.NewLiteral(item.Json_name)
		newobj2, _ := rdf.NewLiteral(item.Json_unit)
		newobj3, _ := rdf.NewLiteral(item.Json_unit_descript)
		newobj4, _ := rdf.NewLiteral(item.Json_descript)

		newtriple1 := rdf.Triple{Subj: newsub, Pred: newpred1, Obj: newobj1}
		newtriple2 := rdf.Triple{Subj: newsub, Pred: newpred2, Obj: newobj2}
		newtriple3 := rdf.Triple{Subj: newsub, Pred: newpred3, Obj: newobj3}
		newtriple4 := rdf.Triple{Subj: newsub, Pred: newpred4, Obj: newobj4}

		tr = append(tr, newtriple0)
		if newtriple1.Obj.String() != "" {
			tr = append(tr, newtriple1)
		}
		if newtriple2.Obj.String() != "" {
			tr = append(tr, newtriple2)
		}
		if newtriple3.Obj.String() != "" {
			tr = append(tr, newtriple3)
		}
		if newtriple4.Obj.String() != "" {
			tr = append(tr, newtriple4)
		}

	}

	// Loop on the ocd_meatadata.txt file
	// RDF version  (many of these sets of 12 could be combined for reduced line count)
	for _, item := range csvdata {
		newsub, _ := rdf.NewIRI(fmt.Sprintf("http://opencoredata/id/resource/janus/%s/%s", item.Struct_name, item.Column_id)) // Sprintf a correct URI here

		newpred0, _ := rdf.NewIRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type")
		newobj0, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/JanusQuerySet")
		newtriple0 := rdf.Triple{Subj: newsub, Pred: newpred0, Obj: newobj0}

		newpred1, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/struct_name")
		newpred2, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/table_name")
		newpred3, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/column_name")
		newpred4, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/column_id")
		newpred5, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/go_struct_name")
		newpred6, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/go_struct_type")
		newpred7, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/JanusMeasurement")
		// newpred7, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/json_name")
		newpred8, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/code")
		newpred9, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/xs_type")
		// newpred10, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/json_unit")
		// newpred11, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/json_unit_descript")
		// newpred12, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/json_descript")

		newobj1, _ := rdf.NewLiteral(item.Struct_name)
		// make IRI that links up to items from _hdr file
		// newobj2, _ := rdf.NewLiteral(item.Table_name)
		newobj2, _ := rdf.NewIRI(fmt.Sprintf("http://opencoredata/id/resource/janus/query/%s", strings.ToLower(item.Table_name)))
		newobj3, _ := rdf.NewLiteral(item.Column_name)
		newobj4, _ := rdf.NewLiteral(item.Column_id)
		newobj5, _ := rdf.NewLiteral(item.Go_struct_name)
		newobj6, _ := rdf.NewLiteral(item.Go_struct_type)
		// newobj7, _ := rdf.NewLiteral(item.Json_name) // make IRI that links up to items from _json_skos file
		newobj7, _ := rdf.NewIRI(fmt.Sprintf("http://opencoredata/id/resource/janus/measure/%s", strings.ToLower(item.Json_name)))
		newobj8, _ := rdf.NewLiteral(item.Code)
		newobj9, _ := rdf.NewLiteral(item.Xs_type)
		// newobj10, _ := rdf.NewLiteral(item.Json_unit)          // resolved by JSON_name IRI
		// newobj11, _ := rdf.NewLiteral(item.Json_unit_descript) // resolved by JSON_name IRI
		// newobj12, _ := rdf.NewLiteral(item.Json_descript)      // resolved by JSON_name IRI

		newtriple1 := rdf.Triple{Subj: newsub, Pred: newpred1, Obj: newobj1}
		newtriple2 := rdf.Triple{Subj: newsub, Pred: newpred2, Obj: newobj2}
		newtriple3 := rdf.Triple{Subj: newsub, Pred: newpred3, Obj: newobj3}
		newtriple4 := rdf.Triple{Subj: newsub, Pred: newpred4, Obj: newobj4}
		newtriple5 := rdf.Triple{Subj: newsub, Pred: newpred5, Obj: newobj5}
		newtriple6 := rdf.Triple{Subj: newsub, Pred: newpred6, Obj: newobj6}
		newtriple7 := rdf.Triple{Subj: newsub, Pred: newpred7, Obj: newobj7}
		newtriple8 := rdf.Triple{Subj: newsub, Pred: newpred8, Obj: newobj8}
		newtriple9 := rdf.Triple{Subj: newsub, Pred: newpred9, Obj: newobj9}
		// newtriple10 := rdf.Triple{Subj: newsub, Pred: newpred10, Obj: newobj10}
		// newtriple11 := rdf.Triple{Subj: newsub, Pred: newpred11, Obj: newobj11}
		// newtriple12 := rdf.Triple{Subj: newsub, Pred: newpred12, Obj: newobj12}

		tr = append(tr, newtriple0)
		if newtriple1.Obj.String() != "" {
			tr = append(tr, newtriple1)
		}
		if newtriple2.Obj.String() != "" {
			tr = append(tr, newtriple2)
		}
		if newtriple3.Obj.String() != "" {
			tr = append(tr, newtriple3)
		}
		if newtriple4.Obj.String() != "" {
			tr = append(tr, newtriple4)
		}
		if newtriple5.Obj.String() != "" {
			tr = append(tr, newtriple5)
		}
		if newtriple6.Obj.String() != "" {
			tr = append(tr, newtriple6)
		}
		if newtriple7.Obj.String() != "" {
			tr = append(tr, newtriple7)
		}
		if newtriple8.Obj.String() != "" {
			tr = append(tr, newtriple8)
		}
		if newtriple9.Obj.String() != "" {
			tr = append(tr, newtriple9)
		}
		// if newtriple10.Obj.String() != "" {
		// 	tr = append(tr, newtriple10)
		// }
		// if newtriple11.Obj.String() != "" {
		// 	tr = append(tr, newtriple11)
		// }
		// if newtriple12.Obj.String() != "" {
		// 	tr = append(tr, newtriple12)
		// }

	}

	// Create the output file
	outFile, err := os.Create("test.nt")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	// Write triples to a file
	var inoutFormat rdf.Format
	inoutFormat = rdf.NTriples // Turtle NQuads
	enc := rdf.NewTripleEncoder(outFile, inoutFormat)
	err = enc.EncodeAll(tr)
	// err = enc.Encode(newtriple)
	enc.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func readJSONSKOSData() []JSONSKOS {
	csvFile, err := os.Open("./ocd_metadata_json_skos.txt")
	defer csvFile.Close()
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(csvFile)
	r.Comma = '\t' // Use tab-delimited instead of comma

	r.FieldsPerRecord = -1
	lines, err := r.ReadAll()
	if err != nil {
		log.Fatalf("Error reading all lines: %v", err)
	}

	commentLines := 1
	observations := make([]JSONSKOS, len(lines)-commentLines) //  why did I do this and not a slice?  then just append?

	for i, line := range lines {
		if i < commentLines {
			continue // skip header lines
		}

		//line0, err := strconv.ParseFloat(line[0], 64)
		if err != nil {
			panic(err)
		}

		// ToDo..   filter function for each string to remove quotes and things like degree symbols
		ob := JSONSKOS{Json_name: strings.ToLower(cleanString(line[0])),
			Json_unit:          cleanString(line[1]),
			Json_unit_descript: cleanString(line[2]),
			Json_descript:      cleanString(line[3])}

		observations[i-commentLines] = ob
	}

	return observations
}

func readHDRData() []HDR {
	csvFile, err := os.Open("./ocd_metadata_hdr.txt")
	defer csvFile.Close()
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(csvFile)
	r.Comma = '\t' // Use tab-delimited instead of comma

	r.FieldsPerRecord = -1
	lines, err := r.ReadAll()
	if err != nil {
		log.Fatalf("Error reading all lines: %v", err)
	}

	commentLines := 1
	observations := make([]HDR, len(lines)-commentLines) //  why did I do this and not a slice?  then just append?

	for i, line := range lines {
		if i < commentLines {
			continue // skip header lines
		}

		//line0, err := strconv.ParseFloat(line[0], 64)
		if err != nil {
			panic(err)
		}

		// ToDo..   filter function for each string to remove quotes and things like degree symbols
		ob := HDR{Table_name: strings.ToLower(cleanString(line[0])),
			Struct_name:     cleanString(line[1]),
			Struct_descript: cleanString(line[2]),
			Query_descript:  cleanString(line[3])}

		observations[i-commentLines] = ob
	}

	return observations
}

func readMetaData() []Voc {
	csvFile, err := os.Open("./ocd_metadata.txt")
	defer csvFile.Close()
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(csvFile)
	r.Comma = '\t' // Use tab-delimited instead of comma

	r.FieldsPerRecord = -1
	lines, err := r.ReadAll()
	if err != nil {
		log.Fatalf("Error reading all lines: %v", err)
	}

	commentLines := 1
	observations := make([]Voc, len(lines)-commentLines) //  why did I do this and not a slice?  then just append?

	for i, line := range lines {
		if i < commentLines {
			continue // skip header lines
		}

		//line0, err := strconv.ParseFloat(line[0], 64)
		if err != nil {
			panic(err)
		}

		// ToDo..   filter function for each string to remove quotes and things like degree symbols
		ob := Voc{Struct_name: strings.ToLower(cleanString(line[0])),
			Table_name:         cleanString(line[1]),
			Column_name:        cleanString(line[2]),
			Column_id:          cleanString(line[3]),
			Go_struct_name:     cleanString(line[4]),
			Go_struct_type:     cleanString(line[5]),
			Json_name:          cleanString(line[6]),
			Code:               cleanString(line[7]),
			Xs_type:            cleanString(line[8]),
			Json_unit:          cleanString(line[9]),
			Json_unit_descript: cleanString(line[10]),
			Json_descript:      cleanString(line[11])}

		observations[i-commentLines] = ob
	}

	return observations
}

func cleanString(input string) string {
	// s := fmt.Sprintf("%+q\n", strings.Replace(input, "\"", "'", -1))
	// log.Printf("in with %s out with %s", input, s)
	// return s
	s := strings.Replace(input, "\"", "'", -1)
	s2 := strings.Replace(s, ")", "", -1)
	s3 := strings.Replace(s2, "(", "", -1)
	s4 := strings.Replace(s3, "–", "", -1)
	return s4
}
