package main

import (
	// "encoding/json"
	"fmt"
	"github.com/blevesearch/bleve"
	// ocdServices "opencoredata.org/ocdServices/documents"
)

func main() {

	index, berr := bleve.Open("../bleveIndexer/csvw.bleve")
	if berr != nil {
		fmt.Printf("this is error %v \n", berr)
	}

	// search for some text
	query := bleve.NewMatchQuery("JanusXRFSample")
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)
	if err != nil {
		fmt.Printf("this is error %v \n", err)
	}
	fmt.Printf("Results %v\n\n", searchResults)

	// fuzzy search
	fuzzyq := bleve.NewFuzzyQuery("JanusXrfSample")
	fsearch := bleve.NewSearchRequest(fuzzyq)
	fsearchResults, err := index.Search(fsearch)
	if err != nil {
		fmt.Printf("this is error %v \n", err)
	}
	fmt.Printf("Fuzzy Results %v\n\n", fsearchResults)

}
