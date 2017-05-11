package main

import (
	"bytes"
	"compress/gzip"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/knakk/sparql"
)

type sampleinfo struct {
	Igsn       string
	OriginID   string
	ExternalID string
}

type sampleinfo2 struct {
	OriginID   string
	Igsn       string
	ExternalID string
}

const queries = `
# Comments are ignored, except those tagging a query.

# tag: samplecall
PREFIX qb:  <http://purl.org/linked-data/cube#>
PREFIX iodp: <http://data.oceandrilling.org/core/1/>
PREFIX janus: <http://data.oceandrilling.org/janus/>
PREFIX sdmx-dimension:  <http://purl.org/linked-data/sdmx/2009/dimension#>
SELECT   ?ob ?leg ?site ?hole ?top ?bottom
FROM <http://data.oceandrilling.org/janus/>
WHERE {
  ?sliceKey iodp:leg ?leg  .
  ?sliceKey iodp:site ?site  .
  ?slice  qb:sliceStructure <http://data.oceandrilling.org/janus/sliceBysampleCount> .
  ?slice qb:sliceStructure ?sliceKey .
  ?slice qb:observation ?ob .
  ?ob janus:sampleid {{.}} .
  ?ob janus:stopinterval1000 ?top .
  ?ob janus:sbottominterval1000 ?bottom  .
  ?ob janus:hole ?hole
}

# tag: blazecall
PREFIX qb:  <http://purl.org/linked-data/cube#>
PREFIX iodp: <http://data.oceandrilling.org/core/1/>
PREFIX janus: <http://data.oceandrilling.org/janus/>
PREFIX sdmx-dimension:  <http://purl.org/linked-data/sdmx/2009/dimension#>
SELECT   ?ob ?leg ?site ?hole ?top ?bottom
WHERE {
  ?sliceKey iodp:leg ?leg  .
  ?slice  qb:sliceStructure <http://data.oceandrilling.org/janus/sliceBysampleCount> .
  ?slice qb:sliceStructure ?sliceKey .
  ?slice qb:observation ?ob .
  ?ob janus:sampleid  "{{.}}"^^xsd:decimal .
  ?ob janus:stopinterval1000 ?top .
  ?ob janus:sbottominterval1000 ?bottom  .
  ?ob janus:hole ?hole
}
`

func main() {

	// csvFile()
	tabFile()
}

func tabFile() {
	fmt.Println("IGSN to Open Core Sample resolution")

	f, err := os.Open("odp_sample_id.tab")
	// f, err := os.Open("sample2.tab")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// gr, err := gzip.NewReader(f)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer gr.Close()

	cr := csv.NewReader(f)
	cr.Comma = '\t' // single quote for rune assignment recall
	recs, err := cr.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// var data []sampleinfo

	for _, rec := range recs {
		input := sampleinfo2{rec[0], rec[1], rec[2]}
		// fmt.Println(input)
		sparqlCall(input.OriginID)
	}
}

func csvFile() {
	fmt.Println("IGSN to Open Core Sample resolution")

	f, err := os.Open("sample.csv.gz")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	gr, err := gzip.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}
	defer gr.Close()

	cr := csv.NewReader(gr)
	cr.Comma = '|' // single quote for rune assignment recall
	recs, err := cr.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// var data []sampleinfo

	for _, rec := range recs {
		input := sampleinfo{rec[0], rec[1], rec[2]}
		// fmt.Println(input)
		sparqlCall(input.OriginID)
	}
}

func sparqlCall(input string) {

	// repo, err := sparql.NewRepo("http://data.oceandrilling.org/sparql")
	repo, err := sparql.NewRepo("http://localhost:9999/blazegraph/namespace/januscube/sparql")

	if err != nil {
		log.Printf("Make repo: %v\n", err)
	}

	f := bytes.NewBufferString(queries)
	bank := sparql.LoadBank(f)

	q, err := bank.Prepare("blazecall", input)
	if err != nil {
		log.Printf("Bank prepair: %v\n", err)
	}

	res, err := repo.Query(q)

	if err != nil {
		log.Printf("Query call: %v\n", err)
	}

	// Print loop testing
	// bindingsTest := res.Results.Bindings // []map[string][]rdf.Term
	// fmt.Println("res.Results.Bindings:")
	// for k, i := range bindingsTest {
	// 	fmt.Printf("At postion %v with %v and %v\n\n", k, i["long"].Value, i["lat"].Value)
	// }

	// bindingsTest2 := res.Bindings() // map[string][]rdf.Term
	// fmt.Println("res.Bindings():")
	// for k, i := range bindingsTest2 {
	// 	fmt.Printf("At postion %v with %v \n\n", k, i)
	// }

	solutionsTest := res.Solutions() // []map[string][]rdf.Term
	// fmt.Println("res.Solutions():")
	for _, i := range solutionsTest {
		// fmt.Printf("At postion %v with %v \n\n", k, i)
		fmt.Printf("%s with %s \n", input, i["ob"])
	}

	if len(solutionsTest) < 1 {
		fmt.Printf("Nothing found for %s \n", input)
	}

	// fmt.Println(ObjectValForPred(bindingsTest2, "uri", "holeid", "http://opencoredata/id/resource/csdco/project/aafblp-lp06-6a"))

}
