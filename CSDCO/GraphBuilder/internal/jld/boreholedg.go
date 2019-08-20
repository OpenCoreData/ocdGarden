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
		"@id":                     fmt.Sprintf("http://opencoredata.org/id/do/%s", dm.Hole_ID),
		"@type":                   "http://opencoredata.org/voc/csdco/v1/Borehole",
		"http://schema.org/about": fmt.Sprintf("http://opencoredata.org/id/do/%s", dm.Expedition),
		"http://opencoredata.org/voc/csdco/v1/azimuth":         dm.Azimuth,
		"http://opencoredata.org/voc/csdco/v1/dip":             dm.Dip,
		"http://opencoredata.org/voc/csdco/v1/elevation":       dm.Elevation,
		"http://www.w3.org/2003/01/geo/wgs84_pos#lat":          dm.Lat,
		"http://www.w3.org/2003/01/geo/wgs84_pos#long":         dm.Long,
		"http://opencoredata.org/voc/csdco/v1/water_Depth":     dm.Water_Depth,
		"http://opencoredata.org/voc/csdco/v1/mblf_B":          dm.MBLF_B,
		"http://opencoredata.org/voc/csdco/v1/mblf_T":          dm.MBLF_T,
		"http://opencoredata.org/voc/csdco/v1/site":            dm.Site,
		"http://opencoredata.org/voc/csdco/v1/comment":         dm.Comment,
		"http://opencoredata.org/voc/csdco/v1/country":         dm.Country,
		"http://opencoredata.org/voc/csdco/v1/county_Region":   dm.County_Region,
		"http://opencoredata.org/voc/csdco/v1/date":            dm.Date,
		"http://opencoredata.org/voc/csdco/v1/expedition":      dm.Expedition,
		"http://opencoredata.org/voc/csdco/v1/hole":            dm.Hole,
		"http://opencoredata.org/voc/csdco/v1/hole_ID":         dm.Hole_ID,
		"http://opencoredata.org/voc/csdco/v1/IGSN":            dm.IGSN,
		"http://opencoredata.org/voc/csdco/v1/location":        dm.Location,
		"http://opencoredata.org/voc/csdco/v1/location_ID":     dm.Location_ID,
		"http://opencoredata.org/voc/csdco/v1/location_Type":   dm.Location_Type,
		"http://opencoredata.org/voc/csdco/v1/metadata_Source": dm.Metadata_Source,
		"http://opencoredata.org/voc/csdco/v1/original_ID":     dm.Original_ID,
		"http://opencoredata.org/voc/csdco/v1/pi":              dm.PI,
		"http://opencoredata.org/voc/csdco/v1/platform":        dm.Platform,
		"http://opencoredata.org/voc/csdco/v1/position":        dm.Position,
		"http://opencoredata.org/voc/csdco/v1/sample_Type":     dm.Sample_Type,
		"http://opencoredata.org/voc/csdco/v1/siteHole":        dm.SiteHole,
		"http://opencoredata.org/voc/csdco/v1/state_Province":  dm.State_Province,
		"http://opencoredata.org/voc/csdco/v1/platform_name":   dm.Platform_name,
		"http://opencoredata.org/voc/csdco/v1/platform_type":   dm.Platform_type,
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
