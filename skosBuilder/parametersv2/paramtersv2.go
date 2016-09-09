package parametersv2

import (
	"bytes"
	"encoding/csv"
	"log"
	"os"
	"strings"
	"text/template"
)

type Voc struct {
	Go_struct_name     string
	Go_struct_type     string
	Json_name          string
	Code               string
	Json_unit          string
	Json_unit_descript string
	Json_descript      string
}

const ttlSchema = `@prefix rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#> .
@prefix dc: <http://purl.org/dc/elements/1.1/> .
@prefix skos: <http://www.w3.org/2004/02/skos/core#> .
@prefix foo: <http://www.foo.org/foo/> .

<http://opencoredata.org/voc/janus/1/>
    dc:creator "Open Core Data" ;
    dc:description "A vocabulary for describing parameters used in the DSDP/OPD/IODP Janus Database" ;
    dc:title "OCD Janus Controlled Vocabulary" ;
    a skos:ConceptScheme .

`

const ttlTemplate = `<http://opencoredata.org/voc/janus/1/{{.Go_struct_name}}>
    a skos:Concept ;
    skos:definition "{{.Json_descript}}" ;
    skos:inScheme <http://opencoredata.org/voc/janus/1/> ;
    {{if .Go_struct_type}}foo:datatype "{{.Go_struct_type}}" ; {{end}}
    {{if .Json_unit}}foo:unit "{{.Json_unit}}" ; {{end}}
    {{if .Json_unit_descript}}foo:unitDescription "{{.Json_unit_descript}}" ; {{end}}
    skos:prefLabel "{{.Json_name}}" .

`

func Parametersv2() {
	csvdata := readMetaData()

	ht, err := template.New("some template").Parse(ttlTemplate)
	if err != nil {
		log.Printf("template parse failed: %s", err)
	}

	rdfFile, err := os.Create("./ocdJanusParamsv2SKOS.ttl")
	if err != nil {
		panic(err)
	}
	defer rdfFile.Close()

	rdfFile.WriteString(ttlSchema)

	for _, item := range csvdata {
		// log.Println(item)
		var buff = bytes.NewBufferString("")
		err = ht.Execute(buff, item)
		if err != nil {
			log.Printf("RDF template execution failed: %s", err)
		}
		rdfFile.WriteString(string(buff.Bytes()))
		// log.Printf("%s", string(buff.Bytes()))
	}
}

func readMetaData() []Voc {
	csvFile, err := os.Open("./parametersv2/gostructs.csv")
	defer csvFile.Close()
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(csvFile)
	//  don't require given field count since some may not have hole (fix with sparql query update too?)

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
		ob := Voc{Go_struct_name: strings.ToLower(cleanString(line[0])),
			Go_struct_type:     cleanString(line[1]),
			Json_name:          cleanString(line[2]),
			Code:               cleanString(line[3]),
			Json_unit:          cleanString(line[4]),
			Json_unit_descript: cleanString(line[5]),
			Json_descript:      cleanString(line[6])}

		observations[i-commentLines] = ob
	}

	return observations
}

func cleanString(input string) string {
	return strings.Replace(input, "\"", "'", -1)
}
