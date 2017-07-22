package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/blevesearch/bleve"
)

func main() {
	phrase := "leg 127 januspaleosample"
	fmt.Print(string(callToJSON(phrase)))
}

func callToJSON(phrase string) string {
	indexPath := "/Users/dfils/Data/OCDDataVolumes/indexes/compositIndex"

	index, err := bleve.OpenUsing(indexPath, map[string]interface{}{
		"read_only": true,
	})
	if err != nil {
		log.Printf("error opening index %s: %v", indexPath, err)
	} else {
		log.Printf("registered index: at %s", indexPath)
	}

	query := bleve.NewMatchQuery(phrase)
	search := bleve.NewSearchRequestOptions(query, 10, 0, false) // no explanation
	search.Highlight = bleve.NewHighlight()                      // need Stored and IncludeTermVectors in index
	searchResults, err := index.Search(search)

	hits := searchResults.Hits // array of struct DocumentMatch

	for k, item := range hits {
		fmt.Printf("\n%d: %s, %f, %s, %v\n", k, item.Index, item.Score, item.ID, item.Fragments)
		for key, frag := range item.Fragments {
			fmt.Printf("%s   %s\n", key, frag)
		}
	}

	jsonResults, _ := json.MarshalIndent(hits, " ", " ")

	return string(jsonResults)
}
