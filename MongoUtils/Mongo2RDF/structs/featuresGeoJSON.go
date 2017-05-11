package structs

// used in ocdBulk and ocdServices
type Expedition struct {
	Uri                    string
	Lat                    string
	Long                   string
	Hole                   string
	Expedition             string
	Site                   string
	Program                string
	Waterdepth             string
	CoreCount              string
	Initialreportvolume    string
	Coredata               string
	Logdata                string
	Geom                   string
	Scientificprospectus   string
	CoreRecovery           string
	Penetration            string
	Scientificreportvolume string
	Expeditionsite         string
	Preliminaryreport      string
	CoreInterval           string
	PercentRecovery        string
	Drilled                string
	Vcdata                 string
	Note                   string
	Prcoeedingreport       string
}

// ExpeditionGeoJSON is a GeoJSON compliant version of this struct.
// It should be used over the above struct to allow for spatial indexing in
// MongoDB
type ExpeditionGeoJSON struct {
	Uri string
	// Spatial             ExpSpatial       `json:"spatial"`
	// Type string `json:"type"`
	// Geo  ExpGeo    `json:"geometry"`
	// Properties   ExpProperties `json:"properties"`

	Type                   string    `json:"type"`
	Coordinates            []float64 `json:"coordinates"`
	Hole                   string
	Expedition             string
	Site                   string
	Program                string
	Waterdepth             string
	CoreCount              string
	Initialreportvolume    string
	Coredata               string
	Logdata                string
	Geom                   string
	Scientificprospectus   string
	CoreRecovery           string
	Penetration            string
	Scientificreportvolume string
	Expeditionsite         string
	Preliminaryreport      string
	CoreInterval           string
	PercentRecovery        string
	Drilled                string
	Vcdata                 string
	Note                   string
	Prcoeedingreport       string
	Abstract               string
}

// type ExpSpatial struct {
// 	Type string `json:"type"`
// 	Geo  ExpGeo    `json:"geometry"`
// }

// type ExpProperties struct {
//     Name   string `json:"name"`
// }

// type ExpGeo struct {
// 	Type      string `json:"type"`
//     Coordinates [][]string `json:"coordinates"`
// 	// Latitude  string `json:"latitude"`
// 	// Longitude string `json:"longitude"`
// }
