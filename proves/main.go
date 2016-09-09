package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type ProvESDataset struct {
	Id               string
	Doi              string
	Loction          string
	SourceInstrument string
	Collection       string
	Level            string
	Version          string
	Title            string
	Label            string
	Bundle           string
	Prov_type        string // EOS_DATASET

}

func main() {
	fmt.Printf("%s \n", dataSetProv())
}

func dataSetProv() string {
	// set up some of our boiler plate schema.org/Dataset elements
	// need date publishedOn, URL, lat long

	// geodata := Geo{Type: "GeoCoordinates", Latitude: latitude, Longitude: longitude}
	// spatial := Spatial{Type: "Place", Geo: geodata}
	// timenow := time.Now().Format(time.RFC850)
	// distribution := Distribution{Type: "DataDownload", ContentURL: uri, DatePublished: timenow, EncodingFormat: "text/tab-separated-values", InLanguage: "en"}
	// author := Author{Type: "Organization", Description: "NSF funded International Ocean Discovery Program operated by JRSO", Name: "International Ocean Discovery Program", URL: "http://iodp.org"}

	// // contextArray := []interface{"http://schema.org", {"glview": "http://schema.geolink.org/somethingIforgot"}}
	// kewords := fmt.Sprintf("DSDP, OPD, IODP, %s", measurement)

	schemametadata := ProvESDataset{Title: "this is the title", Level: "this is the level"}
	// schemametadata := SchemaOrgMetadata{Context:  ["http://schema.org", {"glview": "http://schema.geolink.org/somethingIforgot"}], Type: "Dataset"}

	schemaorgJSON, err := json.MarshalIndent(schemametadata, "", " ")

	if err != nil {
		log.Fatalf("JSON not encoded %v\n", err)
	}

	// JSON-LD parse snipit
	// dataparsed, _ := jsonld.ParseDataset(schemaorgJSON)
	// fmt.Printf("Serialized:\n %s \n\n", dataparsed.Serialize())

	return string(schemaorgJSON)
}
