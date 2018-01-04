package main

import (
	"log"

	redis "gopkg.in/redis.v5"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:9851",
	})

	// cmd := redis.NewStringCmd("SET", "fleet", "truck1", "POINT", 23.32, 115.423)
	// client.Process(cmd)
	// v, _ := cmd.Result()
	// log.Println(v)

	// cmd1 := redis.NewStringCmd("GET", "test", "hono")
	// cmd1 := redis.NewStringCmd("GET", "fleet", "truck3")
	// cmd1 := redis.NewGeoLocationCmd("INTERSECTS", "p418", "BOUNDS 33.462 -112.268 33.491 -112.245")
	// cmd1 := redis.NewStringCmd("INTERSECTS", "p418", "BOUNDS", 33.462, -112.268, 33.491, -112.245)
	cmd1 := redis.NewStringCmd("INTERSECTS", "p418", "OBJECT", dataExample())
	// cmd1 := redis.NewStringCmd("GET", "p418", "http://opencoredata.org/id/dataset/740caec2-3bb2-473a-9dc2-d132dc51e2c1")
	// cmd1 := redis.NewStringCmd("SCAN", "fleet")

	client.Process(cmd1)
	v1, err := cmd1.Result()

	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	log.Println(v1)

	client.Close()
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
