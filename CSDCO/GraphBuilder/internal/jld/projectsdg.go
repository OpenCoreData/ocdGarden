package jld

import (
	"encoding/json"
	"fmt"

	"github.com/piprate/json-gold/ld"
)

// Project struct for describing a CSDCO project object `
type Project struct {
	Expedition    string
	FullName      string
	Funding       string
	Technique     string
	Discipline    string
	LinkTitle     string
	LinkURL       string
	Lab           string
	Repository    string
	Status        string
	StartDate     string
	Outreach      string
	Investigators string
	Abstract      string
	Features      []PjctFeature
}

// PjctFeature is a feature associated with a project
type PjctFeature struct {
	Name      string
	Latitude  float64
	Longitude float64
}

// ProjectDG make data graph for projects
func ProjectDG(dm Project) ([]byte, error) {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")

	// TODO need an @ID for all levels
	// TODO add the items to the map one at a time checking for null values
	// ref https://schema.org/ResearchProject

	// guid := xid.New()   // move from opaque ID to PROJ name in URI

	var features []map[string]interface{}

	for _, i := range dm.Features {
		feature := make(map[string]interface{})
		feature["@type"] = "GeoCoordinates"
		feature["http://schema.org/latitude"] = i.Latitude
		feature["http://schema.org/longitude"] = i.Longitude
		features = append(features, feature)
	}

	doc := map[string]interface{}{
		"@type":                  "ResearchProject",
		"@id":                    fmt.Sprintf("http://opencoredata.org/id/do/%s", dm.Expedition),
		"http://schema.org/url":  fmt.Sprintf("http://opencoredata.org/id/do/%s", dm.Expedition),
		"http://schema.org/name": dm.FullName,
		"csdco:expedition":       dm.Expedition,
		"csdco:funding":          dm.Funding,
		"csdco:technique":        dm.Technique,
		"csdco:discipline":       dm.Discipline,
		"csdco:linktitle":        dm.LinkTitle,
		"csdco:linkurl":          dm.LinkURL,
		"http://schema.org/location": map[string]interface{}{
			"@type":                 "http://schema.org/Place",
			"http://schema.org/geo": features,
		},
		"csdco:lab":                     dm.Lab,
		"csdco:repository":              dm.Repository,
		"csdco:status":                  dm.Status,
		"csdco:startdate":               dm.StartDate,
		"csdco:outreach":                dm.Outreach,
		"csdco:investigators":           dm.Investigators,
		"csdco:abstract":                dm.Abstract,
		"http://schema.org/description": dm.Abstract,
	}

	// Full_Name, Funding, Technique, Discipline, Link_Title, Link_URL, Lab, Repository, Status, Start_Date, Outreach, Investigators

	context := map[string]interface{}{
		"@context": map[string]interface{}{
			"@vocab":        "http://schema.org/",
			"re3data":       "http://example.org/re3data/0.1/",
			"csdco":         "http://opencoredata.org/voc/csdco/v1/",
			"Expedition":    "http://opencoredata.org/voc/csdco/v1/expedition",
			"Funding":       "http://opencoredata.org/voc/csdco/v1/funding",
			"Technique":     "http://opencoredata.org/voc/csdco/v1/technique",
			"Discipline":    "http://opencoredata.org/voc/csdco/v1/discipline",
			"Linktitle":     "http://opencoredata.org/voc/csdco/v1/linktitle",
			"Linkurl":       "http://opencoredata.org/voc/csdco/v1/linkurl",
			"Lab":           "http://opencoredata.org/voc/csdco/v1/lab",
			"Repository":    "http://opencoredata.org/voc/csdco/v1/repository",
			"Status":        "http://opencoredata.org/voc/csdco/v1/status",
			"Startdate":     "http://opencoredata.org/voc/csdco/v1/startdate",
			"Outreach":      "http://opencoredata.org/voc/csdco/v1/outreach",
			"Investigators": "http://opencoredata.org/voc/csdco/v1/investigators",
			"Abstract":      "http://opencoredata.org/voc/csdco/v1/abstract",
		},
	}

	compactedDoc, err := proc.Compact(doc, context, options)
	if err != nil {
		fmt.Println("Error when compacting", err)
	}

	return json.MarshalIndent(compactedDoc, "", " ")
}
