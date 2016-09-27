package main

import (
	"bytes"
	"fmt"
	"log"

	sparql "github.com/knakk/sparql"
	// sparql "opencoredata.org/ocdCommons/sparqlclient"
)

const queries = `
# Comments are ignored, except those tagging a query.

# tag: test1
SELECT DISTINCT *
WHERE 
{ 
  ?uri rdf:type <http://opencoredata.org/id/voc/csdco/v1/CSDCOProject> . 
  ?uri <http://opencoredata.org/id/voc/csdco/v1/project> "AAFBLP" . 
  ?uri ?p ?o . 
}

# tag: test2
SELECT DISTINCT *
WHERE 
{ 
  ?uri ?p ?o . 
}
LIMIT 10

# tag: test3
select distinct ?Concept where {[] a ?Concept} LIMIT 100
`

func main() {
	repo, err := sparql.NewRepo("http://localhost:9999/blazegraph/namespace/csdco/sparql")
	// repo, err := sparql.NewRepo("http://opencoredata.org/sparql")

	if err != nil {
		log.Printf("query make repo: %v\n", err)
	}

	f := bytes.NewBufferString(queries)
	bank := sparql.LoadBank(f)

	q, err := bank.Prepare("test1")
	if err != nil {
		log.Printf("query bank prepair: %v\n", err)
	}

	res, err := repo.Query(q)

	if err != nil {
		log.Printf("query call: %v\n", err)
	}

	// Print loop testing
	bindingsTest := res.Results.Bindings // map[string][]rdf.Term
	fmt.Println("res.Results.Bindings:")
	for k, i := range bindingsTest {
		fmt.Printf("At postion %v with %v and %v\n\n", k, i["p"], i["o"])
	}

	bindingsTest2 := res.Bindings() // map[string][]rdf.Term
	fmt.Println("res.Bindings():")
	for k, i := range bindingsTest2 {
		fmt.Printf("At postion %v with %v \n\n", k, i)
	}

	solutionsTest := res.Solutions() // map[string][]rdf.Term
	fmt.Println("res.Solutions():")
	for k, i := range solutionsTest {
		fmt.Printf("At postion %v with %v \n\n", k, i)
	}

}
