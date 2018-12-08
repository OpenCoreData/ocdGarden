package jld

import (
	"log"

	"opencoredata.org/ocdGarden/CSDCOGraphBuilder/utils"
)

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

// Build makes a JSON-LD file from the CSDCO metadata spreadsheet
func Build() string {
	log.Println("In JSON-LD Builder")

	cm := utils.CSDCO{}
	dm := CSDCODataset{}

	// TODO:  need an @ID for all levels
	doc := map[string]interface{}{
		"@type": "Action",
		"@id":   dm.ID,
		"http://schema.org/url":                                   dm.URL,
		"http://schema.org/description":                           dm.Description,
		"http://schema.org/keywords":                              dm.Keywords,
		"http://schema.org/license":                               "https://creativecommons.org/publicdomain/zero/1.0/",
		"http://schema.org/name":                                  dm.Name,
		"http://opencoredata.org/id/voc/csdco/v1/locationname":    cm.LocationName,
		"http://opencoredata.org/id/voc/csdco/v1/locationtype":    cm.LocationType,
		"http://opencoredata.org/id/voc/csdco/v1/project":         cm.LocationID,
		"http://opencoredata.org/id/voc/csdco/v1/locationid":      cm.LocationID,
		"http://opencoredata.org/id/voc/csdco/v1/site":            cm.LocationID,
		"http://opencoredata.org/id/voc/csdco/v1/hole":            cm.LocationID,
		"http://opencoredata.org/id/voc/csdco/v1/sitehole":        cm.LocationID,
		"http://opencoredata.org/id/voc/csdco/v1/originalid":      cm.LocationID,
		"http://opencoredata.org/id/voc/csdco/v1/holeid":          cm.LocationID,
		"http://opencoredata.org/id/voc/csdco/v1/platform":        cm.LocationID,
		"http://opencoredata.org/id/voc/csdco/v1/date":            cm.LocationID,
		"http://opencoredata.org/id/voc/csdco/v1/waterdepthm":     cm.LocationID,
		"http://opencoredata.org/id/voc/csdco/v1/country":         cm.LocationID,
		"http://opencoredata.org/id/voc/csdco/v1/state_province":  cm.LocationID,
		"http://opencoredata.org/id/voc/csdco/v1/county_region":   cm.LocationID,
		"http://opencoredata.org/id/voc/csdco/v1/pi":              cm.LocationID,
		"http://opencoredata.org/id/voc/csdco/v1/elevation":       cm.LocationID,
		"http://opencoredata.org/id/voc/csdco/v1/position":        cm.LocationID,
		"http://opencoredata.org/id/voc/csdco/v1/storagelocation": cm.LocationID,
		"http://opencoredata.org/id/voc/csdco/v1/sampletype":      cm.LocationID,
		"http://opencoredata.org/id/voc/csdco/v1/comment":         cm.LocationID,
		"http://opencoredata.org/id/voc/csdco/v1/mblft":           cm.LocationID,
		"http://opencoredata.org/id/voc/csdco/v1/mblfb":           cm.LocationID,
		"http://opencoredata.org/id/voc/csdco/v1/metadatasource":  cm.LocationID,
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
