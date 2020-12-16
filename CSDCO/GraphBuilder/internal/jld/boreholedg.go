package jld

import (
	"encoding/json"
	"fmt"

	"github.com/piprate/json-gold/ld"
)

type CSDCOBorehole struct {
	Azimuth         float64
	Dip             float64
	Elevation       float64
	Lat             float64
	Long            float64
	Water_Depth     float64
	MBLF_B          float64
	MBLF_T          float64
	Site            int64
	Comment         string
	Country         string
	County_Region   string
	Date            string
	Expedition      string
	Hole            string
	Hole_ID         string
	IGSN            string
	Location        string
	Location_ID     string
	Location_Type   string
	Metadata_Source string
	NGDC_Serial     string
	Original_ID     string
	PI              string
	Platform        string
	Position        string
	Sample_Type     string
	SiteHole        string
	State_Province  string
	Platform_name   string
	Platform_type   string
}

// BoreholeDG makes data graph for borehole
func BoreholeDG(dm CSDCOBorehole) ([]byte, error) {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")

	// cm := utils.CSDCO{}
	// guid := xid.New()   // move from opaque ID to Hole_ID name in URI

	// TODO:  need an @ID for all levels
	// TODO should I use the IGSN as the ID, not the XID
	doc := map[string]interface{}{
		"@id":                      fmt.Sprintf("https://opencoredata.org/id/do/%s", dm.Hole_ID),
		"@type":                    "https://opencoredata.org/voc/csdco/v1/Borehole",
		"https://schema.org/about": fmt.Sprintf("https://opencoredata.org/id/do/%s", dm.Expedition),
		"https://opencoredata.org/voc/csdco/v1/azimuth":   dm.Azimuth,
		"https://opencoredata.org/voc/csdco/v1/dip":       dm.Dip,
		"https://opencoredata.org/voc/csdco/v1/elevation": dm.Elevation,
		"https://schema.org/spatialCoverage": map[string]interface{}{
			"@type": "Place",
			"https://schema.org/geo": map[string]interface{}{
				"@type":                        "GeoCoordinates",
				"https://schema.org/latitude":  dm.Lat,
				"https://schema.org/longitude": dm.Long,
			},
		},
		"https://www.w3.org/2003/01/geo/wgs84_pos#lat":          dm.Lat,
		"https://www.w3.org/2003/01/geo/wgs84_pos#long":         dm.Long,
		"https://opencoredata.org/voc/csdco/v1/water_Depth":     dm.Water_Depth,
		"https://opencoredata.org/voc/csdco/v1/mblf_B":          dm.MBLF_B,
		"https://opencoredata.org/voc/csdco/v1/mblf_T":          dm.MBLF_T,
		"https://opencoredata.org/voc/csdco/v1/site":            dm.Site,
		"https://opencoredata.org/voc/csdco/v1/comment":         dm.Comment,
		"https://opencoredata.org/voc/csdco/v1/country":         dm.Country,
		"https://opencoredata.org/voc/csdco/v1/county_Region":   dm.County_Region,
		"https://opencoredata.org/voc/csdco/v1/date":            dm.Date,
		"https://opencoredata.org/voc/csdco/v1/expedition":      dm.Expedition,
		"https://opencoredata.org/voc/csdco/v1/hole":            dm.Hole,
		"https://opencoredata.org/voc/csdco/v1/hole_ID":         dm.Hole_ID,
		"https://opencoredata.org/voc/csdco/v1/IGSN":            dm.IGSN,
		"https://opencoredata.org/voc/csdco/v1/location":        dm.Location,
		"https://opencoredata.org/voc/csdco/v1/location_ID":     dm.Location_ID,
		"https://opencoredata.org/voc/csdco/v1/location_Type":   dm.Location_Type,
		"https://opencoredata.org/voc/csdco/v1/metadata_Source": dm.Metadata_Source,
		"https://opencoredata.org/voc/csdco/v1/original_ID":     dm.Original_ID,
		"https://opencoredata.org/voc/csdco/v1/pi":              dm.PI,
		"https://opencoredata.org/voc/csdco/v1/platform":        dm.Platform,
		"https://opencoredata.org/voc/csdco/v1/position":        dm.Position,
		"https://opencoredata.org/voc/csdco/v1/sample_Type":     dm.Sample_Type,
		"https://opencoredata.org/voc/csdco/v1/siteHole":        dm.SiteHole,
		"https://opencoredata.org/voc/csdco/v1/state_Province":  dm.State_Province,
		"https://opencoredata.org/voc/csdco/v1/platform_name":   dm.Platform_name,
		"https://opencoredata.org/voc/csdco/v1/platform_type":   dm.Platform_type,
	}

	context := map[string]interface{}{
		"@context": map[string]interface{}{
			"@vocab":            "https://schema.org/",
			"re3data":           "https://example.org/re3data/0.1/",
			"csdco":             "https://opencoredata.org/voc/csdco/v1/",
			"Azimuth":           "https://opencoredata.org/voc/csdco/v1/azimuth",
			"Dip":               "https://opencoredata.org/voc/csdco/v1/dip",
			"Elevation":         "https://opencoredata.org/voc/csdco/v1/elevation",
			"Lat":               "https://www.w3.org/2003/01/geo/wgs84_pos#lat",
			"Long":              "https://www.w3.org/2003/01/geo/wgs84_pos#long",
			"Water depth":       "https://opencoredata.org/voc/csdco/v1/water_Depth",
			"MBLF bottom":       "https://opencoredata.org/voc/csdco/v1/mblf_B",
			"MBLF top":          "https://opencoredata.org/voc/csdco/v1/mblf_T",
			"Site":              "https://opencoredata.org/voc/csdco/v1/site",
			"Comment":           "https://opencoredata.org/voc/csdco/v1/comment",
			"Country":           "https://opencoredata.org/voc/csdco/v1/country",
			"County Region":     "https://opencoredata.org/voc/csdco/v1/county_Region",
			"Date":              "https://opencoredata.org/voc/csdco/v1/date",
			"Expedition":        "https://opencoredata.org/voc/csdco/v1/expedition",
			"Hole":              "https://opencoredata.org/voc/csdco/v1/hole",
			"Hole ID":           "https://opencoredata.org/voc/csdco/v1/hole_ID",
			"IGSN":              "https://opencoredata.org/voc/csdco/v1/IGSN",
			"Location":          "https://opencoredata.org/voc/csdco/v1/location",
			"Location ID":       "https://opencoredata.org/voc/csdco/v1/location_ID",
			"Location Type":     "https://opencoredata.org/voc/csdco/v1/location_Type",
			"Metadata Source":   "https://opencoredata.org/voc/csdco/v1/metadata_Source",
			"Original ID":       "https://opencoredata.org/voc/csdco/v1/original_ID",
			"PI":                "https://opencoredata.org/voc/csdco/v1/pi",
			"Platform":          "https://opencoredata.org/voc/csdco/v1/platform",
			"Position":          "https://opencoredata.org/voc/csdco/v1/position",
			"Sample Type":       "https://opencoredata.org/voc/csdco/v1/sample_Type",
			"Site Hole":         "https://opencoredata.org/voc/csdco/v1/siteHole",
			"State or Province": "https://opencoredata.org/voc/csdco/v1/state_Province",
			"Platform Name":     "https://opencoredata.org/voc/csdco/v1/platform_name",
			"Platform Type":     "https://opencoredata.org/voc/csdco/v1/platform_type",
		},
	}

	compactedDoc, err := proc.Compact(doc, context, options)
	if err != nil {
		fmt.Println("Error when compacting", err)
	}

	return json.MarshalIndent(compactedDoc, "", " ")
}
