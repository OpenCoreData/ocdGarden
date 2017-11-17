package main

import (
	"fmt"
	"log"

	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", ":9851")
	if err != nil {
		log.Fatalf("Could not connect: %v\n", err)
	}
	defer c.Close()

	var value1 int
	var value2 interface{}
	// var value2 []GeoReturn
	reply, err := redis.Values(c.Do("INTERSECTS", "p418", "OBJECT", dataExample()))
	if err != nil {
		fmt.Printf("Error in reply %v \n", err)
	}
	if _, err := redis.Scan(reply, &value1, &value2); err != nil {
		fmt.Printf("Error in scan %v \n", err)
	}

	fmt.Println(value1)
	fmt.Println(value2)

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
