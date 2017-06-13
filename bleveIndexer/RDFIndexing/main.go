package main

import "fmt"

/*
TODO
Load NQ file
run through bleve index
*/

type RWGSchemaorg struct {
	Context      string `json:"@context"`
	Type         string `json:"@type"`
	Name         string `json:"name"`
	ContactPoint struct {
		Type        string `json:"@type"`
		Name        string `json:"name"`
		Email       string `json:"email"`
		URL         string `json:"url"`
		ContactType string `json:"contactType"`
	} `json:"contactPoint"`
	URL    string `json:"url"`
	Funder struct {
		Type string `json:"@type"`
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"funder"`
	MemberOf struct {
		Type                string `json:"@type"`
		ProgramName         string `json:"programName"`
		HostingOrganization struct {
			Type string `json:"@type"`
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"hostingOrganization"`
	} `json:"memberOf"`
	PotentialAction []struct {
		Type   string `json:"@type"`
		Target struct {
			Type        string `json:"@type"`
			URLTemplate string `json:"urlTemplate"`
			Description string `json:"description"`
			HTTPMethod  string `json:"httpMethod"`
		} `json:"target"`
	} `json:"potentialAction"`
}

func main() {
	fmt.Println("RDF bleve indexer")
}

func loadFacility() {
	// Given issue with triple version due to blank nodes...
	// load and index the JSON-LD directly

}

func loadGraph() {
	// Call A SPARQL on the RWG graph to build the relations to the URL for the
	// facility.  This is required since I am almost 100% blank nodes in the
	// RWG graph due to JSON-LD to RDF conversion reality.

}

/*
func indexTriples() {

	// open a new index
	mapping := bleve.NewIndexMapping()
	// analyzer := mapping.Ad

	index, berr := bleve.New("rwg.bleve", mapping)
	if berr != nil {
		fmt.Printf("Bleve error making index %v \n", berr)
	}

	// index some data
	for i, item := range csvwDocs {
		berr = index.Index(item.URL, item)
		fmt.Printf("Indexed item %d with URL %s\n", i, item.URL)
		if berr != nil {
			fmt.Printf("Bleve error indexing %v \n", berr)
		}
	}
}
*/
