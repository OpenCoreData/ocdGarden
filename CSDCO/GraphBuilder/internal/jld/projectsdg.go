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
}

// ProjectDG make data graph for projects
func ProjectDG(dm Project) ([]byte, error) {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")

	// TODO need an @ID for all levels
	// TODO add the items to the map one at a time checking for null values
	// ref https://schema.org/ResearchProject

	// guid := xid.New()   // move from opaque ID to PROJ name in URI

	doc := map[string]interface{}{
		"@type":                         "ResearchProject",
		"@id":                           fmt.Sprintf("http://opencoredata.org/id/do/%s", dm.Expedition),
		"http://schema.org/url":         fmt.Sprintf("http://opencoredata.org/id/do/%s", dm.Expedition),
		"http://schema.org/name":        dm.FullName,
		"csdco:expedition":              dm.Expedition,
		"csdco:funding":                 dm.Funding,
		"csdco:technique":               dm.Technique,
		"csdco:discipline":              dm.Discipline,
		"csdco:linktitle":               dm.LinkTitle,
		"csdco:linkurl":                 dm.LinkURL,
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
			"@vocab":  "http://schema.org/",
			"re3data": "http://example.org/re3data/0.1/",
			"csdco":   "http://opencoredata.org/voc/csdco/1/",
		},
	}

	compactedDoc, err := proc.Compact(doc, context, options)
	if err != nil {
		fmt.Println("Error when compacting", err)
	}

	return json.MarshalIndent(compactedDoc, "", " ")
}
