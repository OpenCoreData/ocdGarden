package main

import (
	"fmt"

	"github.com/akhenakh/regionagogo"
)

type testData struct {
	Latitude  float64
	Longitude float64
}

func main() {
	fmt.Println("Test of using https://github.com/akhenakh/regionagogo")

	msg := testData{Latitude: 45.4, Longitude: 67.3}

	gs := regionagogo.NewGeoSearch("region.db")
	r := gs.StabbingQuery(msg.Latitude, msg.Longitude)

}
