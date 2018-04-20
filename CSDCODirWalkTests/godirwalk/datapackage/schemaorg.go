package datapackage

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/piprate/json-gold/ld"
)

// VoidDataset is a struct to hold items from a VOiD file that
// describe a dataset  https://developers.google.com/search/docs/data-types/datasets
type CSDCODataset struct {
	ID                 string
	URL                string // type URL: Location of a page describing the dataset.
	Description        string // A short summary describing a dataset.
	Keywords           string // Keywords summarizing the dataset.
	Name               string // A descriptive name of a dataset (e.g., “Snow depth in Northern Hemisphere”)
	ContentURL         string
	AccrualPeriodicity string
	Issued             string
	License            string
	Publisher          string // Person, Org The name of the dataset creator (person or organization).
	Title              string
	DataDump           string
	Source             string
	LandingPage        string
	DownloadURL        string
	MediaType          string
	SameAs             string             // type URL: Other URLs that can be used to access the dataset page.
	Version            string             // The version number for this dataset.
	VariableMeasured   []VariableMeasured // What does the dataset measure? (e.g., temperature, pressure)
	PublisherDesc      string
	PublisherName      string
	PublisherURL       string
	Latitude           string
	Longitude          string
}

// TODO
// may not have VariableMeasured but might have measurement technique
// http://pending.schema.org/measurementTechnique
//

// VariableMeasured at http://pending.schema.org/variableMeasured
type VariableMeasured struct {
	Name        string
	UnitText    string
	Description string
	URL         string
}

// BuildSchema make a type Dataset JSON-LD for a given CSDCO project
func BuildSchema(projname string) string {
	log.Println("in the schema.org build function")

	// FYI
	// Data from the CSDCO graph:
	//  Location Name,Location Type,Project,Location ID,
	// Site,Hole,SiteHole,Original ID,Hole ID,Platform,
	// Date,Water Depth (m),Country,State_Province,County_Region,
	// PI,Lat,Long,Elevation,Position,Storage Location Working,
	// Storage Location Archive,Sample Type,Comment,mblf T,mblf B,
	// Metadata Source

	dm := CSDCODataset{}

	//  Associate with a PROJECT and with contained FILE IDs?
	dm.ID = "ID" // need an ID approach for the PACKAGE  (proj + sha hash?)
	dm.Description = "The description of the data set"
	dm.Keywords = "find some keywords for this section"
	dm.Name = "The name of the dataset" // same as ID?  plus .zip?
	dm.ContentURL = "http://example.org/foo"
	dm.PublisherDesc = "Description of CSDCO"
	dm.PublisherName = "Name of the publisher"
	dm.PublisherURL = "CSDCO home page"
	dm.Latitude = "45.0"  // from graph
	dm.Longitude = "12.5" // from graph

	schemaorg, _ := dsetBuilder(dm)

	return string(schemaorg)
}

func dsetBuilder(dm CSDCODataset) ([]byte, error) {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")

	// Deal with the arrays first:  here..  variable measured
	// array of maps
	var vma []map[string]interface{}

	// we are basically going from struct to map so we can pass the expected
	// type to the json-ld tools
	for _, v := range dm.VariableMeasured {
		vm := make(map[string]interface{})
		vm["@type"] = "PropertyValue"
		vm["unitText"] = v.UnitText
		vm["url"] = v.URL
		vm["name"] = v.Name
		vm["description"] = v.Description
		vma = append(vma, vm)
	}

	doc := map[string]interface{}{
		"@type": "Dataset",
		"@id":   dm.ID,
		"http://schema.org/url":         dm.URL,
		"http://schema.org/description": dm.Description,
		"http://schema.org/keywords":    dm.Keywords,
		"http://schema.org/license":     "https://creativecommons.org/publicdomain/zero/1.0/",
		"http://schema.org/name":        dm.Name,
		"http://schema.org/distribution": map[string]interface{}{
			"@type": "DataDownload",
			"http://schema.org/contentUrl": dm.ContentURL,
		},
		"http://schema.org/publisher": map[string]interface{}{
			"@type": "Organization",
			"http://schema.org/description": dm.PublisherDesc,
			"http://schema.org/name":        dm.PublisherName,
			"http://schema.org/url":         dm.PublisherURL,
		},
		"http://schema.org/spatialCoverage": map[string]interface{}{
			"@type": "Place",
			"http://schema.org/geo": map[string]interface{}{
				"@type": "GeoCoordinates",
				"http://schema.org/latitude":  dm.Latitude,
				"http://schema.org/longitude": dm.Longitude,
			},
		},
		"http://schema.org/variableMeasured": vma,
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
