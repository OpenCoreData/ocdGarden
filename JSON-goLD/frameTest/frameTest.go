package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/kazarena/json-gold/ld"
)

type SpatialFrameRes struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Spatial struct {
		ID   string `json:"id"`
		Type string `json:"type"`
		Geo  struct {
			ID        string `json:"id"`
			Type      string `json:"type"`
			Latitude  string `json:"latitude"`
			Longitude string `json:"longitude"`
		} `json:"geo"`
	} `json:"spatial"`
}

func main() {
	fmt.Println("Frame testing")

	data := mockDataEvent()

	results := spatialFrame(data)

	for result := range results {
		fmt.Println(results[result].ID)
		fmt.Println(results[result].Spatial.Geo.Longitude)
		fmt.Println(results[result].Spatial.Geo.Latitude)
	}

}

func spatialFrame(jsonld string) []SpatialFrameRes {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")

	// frame := map[string]interface{}{
	// 	"@context": "http://schema.org/",
	// 	"@type":    "GeoCoordinates",
	// }

	frame := map[string]interface{}{
		"@context":   "http://schema.org/",
		"@type":      "Dataset",
		"@explicate": true,
		"@id":        "",
		"spatial": map[string]interface{}{
			"geo": map[string]interface{}{
				"latitude":  "",
				"longitude": "",
			},
		},
	}

	var myInterface interface{}
	err := json.Unmarshal([]byte(jsonld), &myInterface)
	if err != nil {
		log.Println("Error when transforming JSON-LD document to interface:", err)
	}

	framedDoc, err := proc.Frame(myInterface, frame, options) // do I need the options set in order to avoid the large context that seems to be generated?
	if err != nil {
		log.Println("Error when trying to frame document", err)
	}

	graph := framedDoc["@graph"]
	// ld.PrintDocument("JSON-LD graph section", graph)  // debug print....
	jsonm, err := json.MarshalIndent(graph, "", " ")
	if err != nil {
		log.Println("Error trying to marshal data", err)
	}

	// log.Printf("SPATIALFRAME: %s", graph)

	dss := make([]SpatialFrameRes, 0)
	err = json.Unmarshal(jsonm, &dss)
	if err != nil {
		log.Println("Error trying to unmarshal data to struct", err)
	}

	// log.Printf("This is the dss:  %v\n", dss)
	return dss
}

// mockDataEvent  returns example JSONLD for local testing if needed
func mockDataEvent() string {

	data := `  {
		"@context": "http://www.w3.org/ns/csvw",
		"dc:license": {
		 "@id": "http://opendefinition.org/licenses/cc-by/"
		},
		"dc:modified": {
		 "@type": "xsd:date",
		 "@value": "Sunday, 25-Oct-15 14:55:20 EDT"
		},
		"dc:publisher": {
		 "schema:name": "Open Core Data",
		 "schema:url": {
		  "@id": "http://opencoredata.org"
		 }
		},
		"dc:title": "104_643A_JanusAgeDatapoint_VbkOzEdY.csv",
		"dcat:keyword": [
		 "DSDP",
		 "ODP",
		 "IODP",
		 "JanusAgeDatapoint"
		],
		"tableSchema": {
		 "aboutUrl": "",
		 "columns": [
		  {
		   "datatype": "int64",
		   "dc:description": "",
		   "name": "Leg",
		   "required": false,
		   "titles": []
		  },
		  {
		   "datatype": "int64",
		   "dc:description": "",
		   "name": "Site",
		   "required": false,
		   "titles": []
		  },
		  {
		   "datatype": "string",
		   "dc:description": "",
		   "name": "Hole",
		   "required": false,
		   "titles": []
		  },
		  {
		   "datatype": "string",
		   "dc:description": "",
		   "name": "Age_model_type",
		   "required": false,
		   "titles": []
		  },
		  {
		   "datatype": "float64",
		   "dc:description": "",
		   "name": "Depth_mbsf",
		   "required": false,
		   "titles": []
		  },
		  {
		   "datatype": "sql.NullFloat64",
		   "dc:description": "",
		   "name": "Age_ma",
		   "required": false,
		   "titles": []
		  },
		  {
		   "datatype": "sql.NullString",
		   "dc:description": "",
		   "name": "Control_point_comment",
		   "required": false,
		   "titles": []
		  }
		 ],
		 "primaryKey": ""
		},
		"url": "http://opencoredata.org/id/dataset/8266648c-a5f1-4c8a-8889-584fa00a5584"
	   }`

	return data

}
