package main

import "encoding/xml"

// IGSN sample id struct
type IGSNOLD struct {
	XMLName                 xml.Name `xml:"samples"`
	UserCode                string   `xml:"sample>user_code"`
	Sampletype              string   `xml:"sample>sample_type"`
	Name                    string   `xml:"sample>name"`
	Material                string   `xml:"sample>material"`
	Igsn                    string   `xml:"sample>igsn,omitempty"`
	Description             string   `xml:"sample>description,omitempty"`
	AgeMin                  string   `xml:"sample>age_min,omitempty"`
	AgeMax                  string   `xml:"sample>age_max,omitempty"`
	AgeUnit                 string   `xml:"sample>age_unit,omitempty"`
	Collectionmethod        string   `xml:"sample>collection_method,omitempty"`
	Latitude                string   `xml:"sample>latitude,omitempty"`
	Longitude               string   `xml:"sample>longitude,omitempty"`
	Elevation               string   `xml:"sample>elevation,omitempty"`
	PrimaryLocationName     string   `xml:"sample>primary_location_name,omitempty"`
	Country                 string   `xml:"sample>country,omitempty"`
	Province                string   `xml:"sample>province,omitempty"`
	County                  string   `xml:"sample>county,omitempty"`
	Collector               string   `xml:"sample>collector,omitempty"`
	CollectionStartDate     string   `xml:"sample>collection_start_date,omitempty"`
	CollectionDatePrecision string   `xml:"sample>collection_date_precision,omitempty"`
	OriginalArchive         string   `xml:"sample>original_archive,omitempty"`
}
