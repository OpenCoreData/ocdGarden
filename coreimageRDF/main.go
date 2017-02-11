package main

import (
	"fmt"
	"log"
	"os"

	rdf "github.com/knakk/rdf"
)

func main() {
	fmt.Println("Take the RDF from coreimage work and convert into a more modern version")
	coreimage()

}

func coreimage() {
	// Open the input file
	inFile, err := os.Open("./coreimages.ttl")
	defer inFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Decode the existing triples
	var inoutFormat rdf.Format
	inoutFormat = rdf.Turtle // FormatNT
	dec := rdf.NewTripleDecoder(inFile, inoutFormat)
	tr, err := dec.DecodeAll()

	for _, item := range tr {
		fmt.Println(item.Subj)
		fmt.Println(item.Obj)
		fmt.Println(item.Pred)
	}

	// load into RDF

	// transcode into new RDF
}
