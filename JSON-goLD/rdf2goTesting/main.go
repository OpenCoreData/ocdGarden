package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/deiu/rdf2go"
)

type Address struct {
	Street string `predicate:"street"`
}

type Person struct {
	Name     string    `predicate:"name"`
	Age      int       `predicate:"age"`
	Size     int64     `predicate:"size"`
	Male     bool      `predicate:"male"`
	Birth    time.Time `predicate:"birth"`
	Surnames []string  `predicate:"surnames"`
	Addr     Address   `subject:"address"`
}

func main() {

	const jsld = `{
		"@context": {
		  "name": "http://schema.org/name",
		  "description": "http://schema.org/description",
		  "image": {
			"@id": "http://schema.org/image",
			"@type": "@id"
		  },
		  "geo": "http://schema.org/geo",
		  "latitude": {
			"@id": "http://schema.org/latitude",
			"@type": "xsd:float"
		  },
		  "longitude": {
			"@id": "http://schema.org/longitude",
			"@type": "xsd:float"
		  },
		  "xsd": "http://www.w3.org/2001/XMLSchema#"
		},
		"name": "The Empire State Building",
		"description": "The Empire State Building is a 102-story landmark in New York City.",
		"image": "http://www.civil.usherbrooke.ca/cours/gci215a/empire-state-building.jpg",
		"geo": {
		  "latitude": "40.75",
		  "longitude": "73.98"
		}
	  }`

	// trying rdf2go
	baseUri := "https://example.org/foo" // Set a base URI
	g := rdf2go.NewGraph(baseUri)        // Create a new graph
	r := strings.NewReader(jsld)
	g.Parse(r, "application/ld+json") // r is an io.Reader

	fmt.Print(g.String())

	/*

		Expect something like
		_:b0 <http://schema.org/description> "The Empire State Building is a 102-story landmark in New York City." .
		_:b0 <http://schema.org/geo> _:b1 .
		_:b0 <http://schema.org/image> <http://www.civil.usherbrooke.ca/cours/gci215a/empire-state-building.jpg> .
		_:b0 <http://schema.org/name> "The Empire State Building" .
		_:b1 <http://schema.org/latitude> "40.75"^^<http://www.w3.org/2001/XMLSchema#float> .
		_:b1 <http://schema.org/longitude> "73.98"^^<http://www.w3.org/2001/XMLSchema#float> .

		 Get
		_:n0 <http://schema.org/description> "The Empire State Building is a 102-story landmark in New York City."^^<http://www.w3.org/2001/XMLSchema#string> .
		_:n0 <http://schema.org/geo> _:n0 .
		_:n0 <http://schema.org/image> <http://www.civil.usherbrooke.ca/cours/gci215a/empire-state-building.jpg> .
		_:n0 <http://schema.org/name> "The Empire State Building"^^<http://www.w3.org/2001/XMLSchema#string> .
		_:n0 <http://schema.org/latitude> "40.75"^^<http://www.w3.org/2001/XMLSchema#float> .
		_:n0 <http://schema.org/longitude> "73.98"^^<http://www.w3.org/2001/XMLSchema#float> .


		Note the geo line 2 in both output.  The right most blank node can not be _:n0 again.

	*/
}
