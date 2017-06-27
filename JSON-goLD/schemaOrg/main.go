package main

import (
	"encoding/json"
	"fmt"

	"github.com/kazarena/json-gold/ld"
)

// DataSetMD a struct to hold metadata about a data set
type DataSetMD struct {
	ID          string
	URL         string
	Description string
	Keywords    string
	Name        string
	ContentURL  string
}

// DataCatalog is a struct to hold metadata about data catalogs
type DataCatalog struct {
	ID          string
	URL         string
	Description string
}

func main() {
	fmt.Println("Schema.org JSON-LD packet type builder")

	dm1 := DataSetMD{ID: "ID",
		URL:         "http://opencoredata.org/id/dataset/81876abc-9233-45e4-9f0d-294e8158571a",
		Description: "Data set description",
		Keywords:    "DSDP, OPD, IODP, JanusThinSectionImage",
		Name:        "206_1256B_JanusThinSectionImage_uVJzZlfW.csv",
		ContentURL:  "http://opencoredata.org/id/dataset/81876abc-9233-45e4-9f0d-294e8158571a"}

	dm2 := DataSetMD{ID: "ID",
		URL:         "http://opencoredata.org/id/dataset/81876abc-9233-45e4-9f0d-294e8158571a",
		Description: "Data set description",
		Keywords:    "DSDP, OPD, IODP, JanusThinSectionImage",
		Name:        "206_1256B_JanusThinSectionImage_uVJzZlfW.csv",
		ContentURL:  "http://opencoredata.org/id/dataset/81876abc-9233-45e4-9f0d-294e8158571a"}

	dsa := []DataSetMD{}
	dsa = append(dsa, dm1)
	dsa = append(dsa, dm2)

	dc := DataCatalog{ID: "http://opencoredata.org/catalogs", URL: "http://opencoredata.org/catalogs",
		Description: "Can I use this approach to reference this catalog from type WebSite"}

	// Build a schema.org/DataCatalog entry
	cat, _ := catalogBuilder(dc, dsa)
	fmt.Println(string(cat))

	// Build a schema.org/Dataset entry
	dset, _ := dsetBuilder(dm1)
	fmt.Println(string(dset))

}

func dsetBuilder(dm DataSetMD) ([]byte, error) {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")

	doc := map[string]interface{}{
		"@type": "Dataset",
		"@id":   dm.ID,
		"http://schema.org/url":         dm.URL,
		"http://schema.org/description": dm.Description,
		"http://schema.org/keywords":    dm.Keywords,
		"http://schema.org/name":        dm.Name,
		"http://schema.org/distribution": map[string]interface{}{
			"@type": "DataDownload",
			"http://schema.org/contentUrl": dm.ContentURL,
		},
	}

	context := map[string]interface{}{
		"@context": map[string]interface{}{
			"@vocab":  "http://schema.org/",
			"re3data": "http://example.org/re3data/0.1/",
		},
	}

	compactedDoc, err := proc.Compact(doc, context, options)
	if err != nil {
		fmt.Println("Error when compacting", err)
	}

	return json.MarshalIndent(compactedDoc, "", " ")
}

func catalogBuilder(dc DataCatalog, dsa []DataSetMD) ([]byte, error) {

	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")

	// array of maps
	var dsArray []map[string]interface{}

	// we are basically going from stuct to map so we can pass the expected
	// type to the json-ld tools
	for _, v := range dsa {
		datasets := make(map[string]interface{})
		datasets["@type"] = "Dataset"
		datasets["description"] = v.Description
		datasets["url"] = v.URL
		dsArray = append(dsArray, datasets)
	}

	doc := map[string]interface{}{
		"@type": "DataCatalog",
		"@id":   dc.ID,
		"http://schema.org/url":         dc.URL,
		"http://schema.org/description": dc.Description,
		"http://schema.org/dataset":     dsArray,
	}

	context := map[string]interface{}{
		"@context": map[string]interface{}{
			"@vocab":  "http://schema.org/",
			"re3data": "http://example.org/re3data/0.1/",
		},
	}

	compactedDoc, err := proc.Compact(doc, context, options)
	if err != nil {
		fmt.Println("Error when compacting", err)
	}

	return json.MarshalIndent(compactedDoc, "", " ")
}
