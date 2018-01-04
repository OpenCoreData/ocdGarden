package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/garyburd/redigo/redis"
	gj "github.com/kpawlik/geojson"
)

type Location struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

func main() {
	c, err := redis.Dial("tcp", "localhost:9851")
	if err != nil {
		log.Fatalf("Could not connect: %v\n", err)
	}
	defer c.Close()

	log.Print("connected")

	var value1 int
	var value2 []interface{}
	// var value2 []GeoReturn
	reply, err := redis.Values(c.Do("INTERSECTS", "p418", "LIMIT", "50000", "OBJECT", dataExample()))
	// reply, err := redis.Values(c.Do("SCAN", "p418"))
	if err != nil {
		fmt.Printf("Error in reply %v \n", err)
	}

	// for len(reply) > 0 {

	// if _, err := redis.Scan(reply, &value1, &value2); err != nil {
	// 	fmt.Printf("Error in scan %v \n", err)
	// }

	// for len(reply) > 0 {
	var title string
	rating := -1 // initialize to illegal value to detect nil.
	reply, err = redis.Scan(reply, &value1, &value2)
	if err != nil {
		fmt.Println(err)
		return
	}
	if rating == -1 {
		fmt.Println(title, "not-rated")
	} else {
		fmt.Println(title, rating)
	}
	// results, _ := tile38RespAsGeoJSON(value2)
	// fmt.Println(results)
	// }

	// sp := fmt.Sprintf("%s", value2)
	// fmt.Println(sp)

	results, _ := tile38RespAsGeoJSON(value2)
	fmt.Println(results)
	fmt.Printf("Count %d \n", value1)

}

func tile38RespAsGeoJSON(results []interface{}) (string, error) {

	// Build the geojson section
	var (
		// fc *gj.FeatureCollection
		f  *gj.Feature
		fa []*gj.Feature
	)

	for _, item := range results {
		valcast := item.([]interface{})
		val0 := fmt.Sprintf("%s", valcast[0])
		val1 := fmt.Sprintf("%s", valcast[1])
		fmt.Println(val0)
		fmt.Println(val1)

		loc := &Location{}
		err := json.Unmarshal([]byte(val1), loc)
		if err != nil {
			return "", err
		}

		fmt.Println(loc.Type)
		fmt.Println(loc.Coordinates[0])
		fmt.Println(loc.Coordinates[1])

		cd := gj.Coordinate{gj.Coord(loc.Coordinates[1]), gj.Coord(loc.Coordinates[0])} // is this long lat..  vs lat long?

		props := map[string]interface{}{"URL": val0}

		newp := gj.NewPoint(cd)
		f = gj.NewFeature(newp, props, nil)
		fa = append(fa, f)
	}

	fc := gj.FeatureCollection{Type: "FeatureCollection", Features: fa}
	gjstr, err := gj.Marshal(fc)
	if err != nil {
		log.Println(err)
	}

	return gjstr, nil
}

func dataExample() string {
	geoobject := `{
		"type": "FeatureCollection",
		"features": [
		  {
			"type": "Feature",
			"properties": {},
			"geometry": {
			  "type": "Polygon",
			  "coordinates": [
				[
				  [
					-113.90625,
					-27.994401411046148
				  ],
				  [
					66.09375,
					-27.994401411046148
				  ],
				  [
					66.09375,
					62.2679226294176
				  ],
				  [
					-113.90625,
					62.2679226294176
				  ],
				  [
					-113.90625,
					-27.994401411046148
				  ]
				]
			  ]
			}
		  }
		]
	  }`

	return geoobject
}
