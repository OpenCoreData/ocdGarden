package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/deiu/rdf2go"
)

type VoidDataset struct {
	ID                 string
	URL                string
	Description        string
	Keywords           string
	Name               string
	ContentURL         string
	AccrualPeriodicity string
	Issued             string
	License            string
	Publisher          string
	Title              string
	DataDump           string
	Source             string
	LandingPage        string
	DownloadURL        string
	MediaType          string
}

func main() {
	fmt.Println("VOID reader")
	output := VoidReader() // pass a URI or file reference   return []VoidDataset
	fmt.Printf("Datasets indexed from VoID file: %d \n", len(output))
}

func VoidReader() []VoidDataset {
	// Set a base URI  (is this the quad URI?)
	baseUri := "https://example.org/foo"

	// Create a new graph
	g := rdf2go.NewGraph(baseUri)
	file, _ := os.Open("void.ttl")
	nr := bufio.NewReader(file)

	// nr is an io.Reader
	g.Parse(nr, "text/turtle")

	fmt.Printf("Read in file with %d triples\n", g.Len())
	var vdsa []VoidDataset

	// read the triples from g
	triples := g.All(nil, nil, rdf2go.NewResource("http://rdfs.org/ns/void#Dataset"))
	for triple := range triples {
		var vds VoidDataset

		fmt.Printf("Found the URI: %s \n", triples[triple].String())

		vds.ID = triples[triple].Subject.RawValue()  //.String()  // hold what we are talking about
		vds.URL = triples[triple].Subject.RawValue() // The ID is the URL for this LOD case

		//    dcat := rdf2go.

		vds.Description = getObject(g, triples[triple].Subject, rdf2go.NewResource("http://purl.org/dc/terms/description"))
		// vds.Keywords = getObject(g, triples[triple].Subject, rdf2go.NewResource("http://purl.org/dc/terms/license"))
		vds.Name = getObject(g, triples[triple].Subject, rdf2go.NewResource("http://purl.org/dc/terms/title"))
		vds.ContentURL = getObject(g, triples[triple].Subject, rdf2go.NewResource("http://www.w3.org/ns/dcat#downloadURL"))
		vds.AccrualPeriodicity = getObject(g, triples[triple].Subject, rdf2go.NewResource("http://purl.org/dc/terms/accrualPeriodicity"))
		vds.Issued = getObject(g, triples[triple].Subject, rdf2go.NewResource("http://purl.org/dc/terms/issued"))
		vds.License = getObject(g, triples[triple].Subject, rdf2go.NewResource("http://purl.org/dc/terms/license"))
		vds.Publisher = getObject(g, triples[triple].Subject, rdf2go.NewResource("http://purl.org/dc/terms/publisher"))
		vds.Title = getObject(g, triples[triple].Subject, rdf2go.NewResource("http://purl.org/dc/terms/title"))
		vds.DataDump = getObject(g, triples[triple].Subject, rdf2go.NewResource("http://rdfs.org/ns/void#dataDump"))
		vds.Source = getObject(g, triples[triple].Subject, rdf2go.NewResource("http://purl.org/dc/terms/source"))
		vds.LandingPage = getObject(g, triples[triple].Subject, rdf2go.NewResource("http://www.w3.org/ns/dcat#landingPage"))
		vds.DownloadURL = getObject(g, triples[triple].Subject, rdf2go.NewResource("http://www.w3.org/ns/dcat#downloadURL"))
		vds.MediaType = getObject(g, triples[triple].Subject, rdf2go.NewResource("http://www.w3.org/ns/dcat#mediaType"))

		fmt.Println(vds.DownloadURL)

		vdsa = append(vdsa, vds)
	}

	return vdsa
}

func getObject(g *rdf2go.Graph, subjectURI, predicateURI rdf2go.Term) string {
	test := g.One(subjectURI, predicateURI, nil)

	if test != nil {
		return test.Object.String()
	} else {
		return ""
	}
}
