package janusQuerries

import (
	"bytes"
	"encoding/csv"
	"log"
	"os"
	"strings"
	"text/template"
)

type Voc struct {
	Name        string
	Description string
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

const ttlTemplate = `<http://opencoredata.org/voc/janus/1/{{.Name}}>
    a skos:Concept ;
    skos:definition "{{.Description}}" ;
    skos:inScheme <http://opencoredata.org/voc/janus/1/> .

`

func JanusQuerries() {
	csvdata := readMetaData()

	ht, err := template.New("Querry template").Parse(ttlTemplate)
	if err != nil {
		log.Printf("Querry template parse failed: %s", err)
	}

	rdfFile, err := os.Create("./ocdJanusQuerries.ttl")
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
	csvFile, err := os.Open("./janusQuerries/janusQuerry.csv")
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
		ob := Voc{Name: strings.ToLower(cleanString(line[0])),
			Description: cleanString(line[1])}

		observations[i-commentLines] = ob
	}

	return observations
}

func cleanString(input string) string {
	return strings.Replace(input, "\"", "'", -1)
}
