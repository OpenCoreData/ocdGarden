package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/knakk/rdf"
	sparql "github.com/knakk/sparql"
	// sparql "opencoredata.org/ocdCommons/sparqlclient"
)

type CSDCO struct {
	LocationName           string
	LocationType           string
	Project                string
	LocationID             string
	Site                   string
	Hole                   string
	SiteHole               string
	OriginalID             string
	HoleID                 string
	Platform               string
	Date                   string
	WaterDepthM            string
	Country                string
	State_Province         string
	County_Region          string
	PI                     string
	Lat                    string
	Long                   string
	Elevation              string
	Position               string
	StorageLocationWorking string
	StorageLocationArchive string
	SampleType             string
	Comment                string
	MblfT                  string
	MblfB                  string
	MetadataSource         string
}

// bring in the DataCite style struct to test serlizing to struct the SPARQL results
type DataCite struct {
	ExpDOI          string   // Is this the ID of the expedition or something else
	ExpURI          string   // something like http://data.rvdata.us/id/cruise/TN272 for R2R
	ResourceType    string   // Field_expedition
	CreatorName     string   // Open Core Data
	CreatorDOI      string   // re3data DOI  static   10.17616/R37936
	Title           string   // Expedition XXX on Joides Resoultion or CSDCO
	Abstract        string   // * abstract here...
	DateCollected   string   // ** Really a data of a specific format 2011-11-05/2011-12-17
	ContributorName string   // Joides Resolution Science Office || Continental Scientific Drilling Corrdinating Office
	RelatedDOIs     []string // 1 or more related DOI's
	Long            string   // longitude
	Lat             string   // latitude
	Publisher       string   // Rolling Deck to Repository (R2R) Program
	Version         string   // 1, 2, 3, etc
	PubYear         string   // 2016
}

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

# Get all the info on a HoleID from the CSDCO graph
# tag: CSDCOHoleID
SELECT *
WHERE 
{ 
  <{{.HOLEID}}>  ?p ?o .
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


# tag: focusedCall
SELECT DISTINCT ?uri ?date ?lat ?long ?holeid
WHERE 
{ 
  ?uri rdf:type <http://opencoredata.org/id/voc/csdco/v1/CSDCOProject> . 
  ?uri <http://opencoredata.org/id/voc/csdco/v1/project> "AAFBLP" . 
  ?uri <http://opencoredata.org/id/voc/csdco/v1/holeid> ?holeid .
  ?uri 	<http://opencoredata.org/id/voc/csdco/v1/date> ?date . 
  ?uri 	<http://www.w3.org/2003/01/geo/wgs84_pos#lat> ?lat .
  ?uri 	<http://www.w3.org/2003/01/geo/wgs84_pos#long> ?long .
}
`

func main() {
	// callTest1()
	callTest2()
	// callTest3()
}

func callTest1() {
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

func callTest2() {

	repo, err := sparql.NewRepo("http://localhost:9999/blazegraph/namespace/csdcov3/sparql")
	// repo, err := sparql.NewRepo("http://opencoredata.org/sparql")

	if err != nil {
		log.Printf("query make repo: %v\n", err)
	}

	f := bytes.NewBufferString(queries)
	bank := sparql.LoadBank(f)

	q, err := bank.Prepare("focusedCall")
	if err != nil {
		log.Printf("query bank prepair: %v\n", err)
	}

	res, err := repo.Query(q)

	if err != nil {
		log.Printf("query call: %v\n", err)
	}

	// fmt.Print(res)

	// Print loop testing
	bindingsTest := res.Results.Bindings // []map[string][]rdf.Term
	fmt.Println("res.Results.Bindings:")
	for k, i := range bindingsTest {
		fmt.Printf("At postion %v with %v and %v\n\n", k, i["long"].Value, i["lat"].Value)
	}

	bindingsTest2 := res.Bindings() // map[string][]rdf.Term
	fmt.Println("res.Bindings():")
	for k, i := range bindingsTest2 {
		fmt.Printf("At postion %v with %v \n\n", k, i)
	}

	solutionsTest := res.Solutions() // []map[string][]rdf.Term
	fmt.Println("res.Solutions():")
	for k, i := range solutionsTest {
		fmt.Printf("At postion %v with %v \n\n", k, i)
	}

	fmt.Println("call test for different sparql format")
	fmt.Println(ObjectValForPred(bindingsTest2, "uri", "holeid", "http://opencoredata/id/resource/csdco/project/aafblp-lp06-6a"))

}

func callTest3() {
	repo, err := sparql.NewRepo("http://localhost:9999/blazegraph/namespace/csdcov3/sparql")
	// repo, err := sparql.NewRepo("http://opencoredata.org/sparql")

	if err != nil {
		log.Printf("query make repo: %v\n", err)
	}

	f := bytes.NewBufferString(queries)
	bank := sparql.LoadBank(f)

	q, err := bank.Prepare("CSDCOHoleID", struct{ HOLEID string }{"http://opencoredata/id/resource/csdco/project/aafblp-llb06-2a"})
	if err != nil {
		log.Printf("query bank prepair: %v\n", err)
	}

	res, err := repo.Query(q)

	if err != nil {
		log.Printf("query call: %v\n", err)
	}

	data := CSDCO{}

	// Print loop testing
	bindingsTest := res.Results.Bindings // []map[string][]rdf.Term
	fmt.Println("res.Results.Bindings:")
	for k, i := range bindingsTest {
		fmt.Printf("At postion %v with %v and %v\n\n", k, i["p"].Value, i["o"].Value)
	}

	// data.Country = bindingsTest[0]["o"].Value // need to know index value

	bindingsTest2 := res.Bindings() // map[string][]rdf.Term
	fmt.Println("res.Bindings():")
	for k, i := range bindingsTest2 {
		fmt.Printf("At postion %v with %v \n\n", k, i)
	}

	// data.Hole

	solutionsTest := res.Solutions() // []map[string][]rdf.Term
	fmt.Println("res.Solutions():")
	for k, i := range solutionsTest {
		fmt.Printf("At postion %v with %v \n\n", k, i)
	}

	// what I need for a function then Is
	data.Hole = ObjectValForPred(bindingsTest2, "p", "o", "http://opencoredata.org/id/voc/csdco/v1/holeid")
	data.Country = ObjectValForPred(bindingsTest2, "p", "o", "http://opencoredata.org/id/voc/csdco/v1/country")
	data.Elevation = ObjectValForPred(bindingsTest2, "p", "o", "http://opencoredata.org/id/voc/csdco/v1/elevation")

	fmt.Println(data)

}

func ObjectValForPred(sparql map[string][]rdf.Term, predcol string, objectcol string, predicate string) string {
	pSlice := sparql[predcol]
	oSlice := sparql[objectcol]
	indexval := SliceIndex(len(pSlice), func(i int) bool { return pSlice[i].String() == predicate })
	// fmt.Println(indexval)
	// if -1 above return what?
	if indexval == -1 {
		return ""
	}
	return oSlice[indexval].String()
}

// SliceIndex return int location of item in slice  http://stackoverflow.com/questions/8307478/go-how-to-find-out-element-position-in-slice
func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}
