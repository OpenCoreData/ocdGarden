package datapackage

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/piprate/json-gold/ld"
)

// http://opencoredata.org/collections/csdco/project/MBGLH
// or
// http://opencoredata.org/collections/csdco/project/LAUCA
// http://opencoredata.org/id/csdco/dataset/
// http://opencoredata.org/id/ocd/dataset/

// VoidDataset is a struct to hold items from a VOiD file that
// describe a dataset  https://developers.google.com/search/docs/data-types/datasets
type CSDCODataset struct {
	ID                 string // simple UUID
	URL                string // http://opencoredata.org/id/csdco/dataset/
	Description        string // A short summary describing a dataset.
	Keywords           string // Keywords summarizing the dataset.
	Name               string // CSDCO Project X dataset Y
	ContentURL         string // http://opencoredata.org/api/v1/documents/download/PROJNAME.csv
	AccrualPeriodicity string
	Issued             string
	License            string             // CC zero
	Publisher          string             // CSDCO RE3 DOI
	Title              string             // same as name?
	DataDump           string             // ????
	Source             string             // CSDCO Project X
	LandingPage        string             // as opposed to URL above?
	DownloadURL        string             // as opposed to conenturl above?
	MediaType          string             // zip file application/zip
	SameAs             string             // not used, but might be with Carp Lake
	Version            string             // 0.1.1
	VariableMeasured   []VariableMeasured // What does the dataset measure? (e.g., temperature, pressure)
	PublisherDesc      string             // CSDCO desc
	PublisherName      string             // CSDCO name
	PublisherURL       string             // CSDCO URL
	Latitude           string             //  can't be used?   or is there _one_ for the project?
	Longitude          string             //  "    "    "    "   "
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
// TODO:  need to pass fileName fileURI projname datapakcage.json?  (to get things from it?)
func BuildSchema(projname, workingdir, shavalue string) string {
	log.Println("in the schema.org build function")

	// FYI
	// Data from the CSDCO graph:
	// Location Name,Location Type,Project,Location ID,
	// Site,Hole,SiteHole,Original ID,Hole ID,Platform,
	// Date,Water Depth (m),Country,State_Province,County_Region,
	// PI,Lat,Long,Elevation,Position,Storage Location Working,
	// Storage Location Archive,Sample Type,Comment,mblf T,mblf B,
	// Metadata Source

	//ocg := ocdGraphCall  // returns a CSDCOProject{} struct

	dm := CSDCODataset{}

	//  Associate with a PROJECT and with contained FILE IDs?
	dm.ID = fmt.Sprintf("http://opencoredata.org/pkg/id/%s", shavalue) // need an ID approach for the PACKAGE  (proj + sha hash?)
	dm.Description = fmt.Sprintf("A CSDCO data package for  project %s (%s)", projname, workingdir)
	dm.Keywords = "CSDCO, Continental Scientific Drilling"
	dm.Name = fmt.Sprintf("%s.zip", shavalue)
	dm.ContentURL = fmt.Sprintf("http://opencoredata.org/pkg/id/%s.zip", shavalue)
	dm.PublisherDesc = "Continental Scientific Drilling Coordination Office"
	dm.PublisherName = "CSDCO"
	dm.PublisherURL = "https://csdco.umn.edu/"
	dm.Latitude = "0.0"  // from graph  ocg.lat
	dm.Longitude = "0.0" // from graph  ocg.long

	schemaorg, _ := dsetBuilder(dm)

	return string(schemaorg)
}

func ocdGraphCall() {

	// make a SPARQL call to opencore to get info about a project.
}

func dsetBuilder(dm CSDCODataset) ([]byte, error) {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")

	// Deal with the arrays first:  here..  variable measured
	// array of maps
	var vma []map[string]interface{}

	// wont have VarMeas...  but might have MesTech
	// ..   really we would just sappend into the description of the dataset the
	// file types present
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
		"http://schema.org/url":         dm.ID,
		"http://schema.org/description": dm.Description,
		"http://schema.org/keywords":    dm.Keywords,
		"http://schema.org/license":     "https://creativecommons.org/publicdomain/zero/1.0/",
		"http://schema.org/name":        dm.Name,
		"http://schema.org/distribution": map[string]interface{}{
			"@type": "DataDownload",
			"http://schema.org/contentUrl": dm.ContentURL,
			"http://schema.org/fileFormat": "application/vnd.datapackage+json",
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
		log.Println("Error when compacting", err)
	}

	return json.MarshalIndent(compactedDoc, "", " ")
}
