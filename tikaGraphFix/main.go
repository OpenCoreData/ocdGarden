package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	rdf "github.com/knakk/rdf"
)

func main() {
	fmt.Println("Tika graph fixer")

	// Open the input file
	inFile, err := os.Open("./opencoredataorg_fdpjena.n3")
	defer inFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Decode the existing triples
	var inoutFormat rdf.Format
	inoutFormat = rdf.NTriples // FormatTTL
	dec := rdf.NewTripleDecoder(inFile, inoutFormat)
	tr, err := dec.DecodeAll()

	for _, item := range tr {
		u, err := url.Parse(item.Subj.String())
		if err != nil {
			fmt.Println(err)
		}

		pa := strings.Split(u.Path, "/")

		s, err := rdf.NewIRI(fmt.Sprintf("https://%s%s", u.Host, strings.Join(pa[0:4], "/")))
		if err != nil {
			fmt.Println(err)
		}

		p, err := rdf.NewIRI("http://schema.org/hasPart")
		if err != nil {
			fmt.Println(err)
		}

		o, err := rdf.NewIRI(item.Subj.String())
		if err != nil {
			fmt.Println(err)
		}

		t := rdf.Triple{Subj: s, Pred: p, Obj: o}
		fmt.Println(t.Serialize(rdf.NQuads))
		tr = append(tr, t)
	}

	// Create the output file
	outFile, err := os.Create("test.nt")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	// Write triples to a file
	enc := rdf.NewTripleEncoder(outFile, inoutFormat)
	err = enc.EncodeAll(tr)
	// err = enc.Encode(newtriple)
	enc.Close()
	if err != nil {
		log.Fatal(err)
	}
}
