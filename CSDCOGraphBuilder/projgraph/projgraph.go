package projgraph

import (
	"log"

	"opencoredata.org/ocdGarden/CSDCOGraphBuilder/jld"
)

func BuildProjGraph() string {
	log.Println("In projgraph builder")

	// TODO:
	// read in the CSV data
	// find the unique proj names
	// get the project level metadata (not the holeid level)
	// this might be a small amount of the data

	//cm := utils.CSDCO{}
	dm := jld.CSDCODataset{}

	// TODO:  need an @ID for all levels
	doc := map[string]interface{}{
		"@type": "ResearchProject",
		"@id":   dm.ID,
		"http://schema.org/url":         dm.URL,
		"http://schema.org/description": dm.Description,
		"http://schema.org/keywords":    dm.Keywords,
		"http://schema.org/name":        dm.Name,
		"http://schema.org/publisher": map[string]interface{}{
			"@type": "Organization",
			"http://schema.org/description": dm.PublisherDesc,
			"http://schema.org/name":        dm.PublisherName,
			"http://schema.org/url":         dm.PublisherURL,
		},
	}

	context := map[string]interface{}{
		"@context": map[string]interface{}{
			"@vocab":  "http://schema.org/",
			"re3data": "http://example.org/re3data/0.1/",
		},
	}

	log.Println(doc)
	log.Println(context)

	return "done"

}
