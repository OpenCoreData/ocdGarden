package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/knakk/rdf"
)

type ProvESDataset struct {
	Id               string
	Doi              string
	Loction          string
	SourceInstrument string
	Collection       string
	Level            string
	Version          string
	Title            string
	Label            string
	Bundle           string
	Prov_type        string // EOS_DATASET

}

func main() {
	// fmt.Printf("%s \n", dataSetProv())
	provTest1("http://opencoredata.org/doc/dataset/045deec9-94b2-445a-8fd2-43dbe90841fb")
}

func provTest1(URI string) {

	tr := []rdf.Triple{}

	newsub, _ := rdf.NewIRI(fmt.Sprintf("%s/prov", URI)) // Sprintf a correct URI here

	// type it an attribution
	newpred0, _ := rdf.NewIRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type")
	newobj0, _ := rdf.NewIRI("http://www.w3.org/ns/prov#Attribution")
	newtriple0 := rdf.Triple{Subj: newsub, Pred: newpred0, Obj: newobj0}

	// assign agent
	newpred1, _ := rdf.NewIRI("http://www.w3.org/ns/prov#agent")
	newobj1, _ := rdf.NewIRI("http://doi.org/10.17616/R37936") // re3data doi for OCD
	newtriple1 := rdf.Triple{Subj: newsub, Pred: newpred1, Obj: newobj1}

	// assign role
	newpred2, _ := rdf.NewIRI("http://www.w3.org/ns/prov#hasRole")
	newobj2, _ := rdf.NewLiteral("Publisher")
	newtriple2 := rdf.Triple{Subj: newsub, Pred: newpred2, Obj: newobj2}

	// assign association
	newpred3, _ := rdf.NewIRI("http://www.w3.org/ns/prov#wasAssociatedWith")
	newobj3, _ := rdf.NewIRI("http://www.w3.org/ns/prov#wasAssociatedWith")
	newtriple3 := rdf.Triple{Subj: newsub, Pred: newpred3, Obj: newobj3}

	// assign qualifiedUsage 1
	newpred4, _ := rdf.NewIRI("http://www.w3.org/ns/prov#qualifiedUsage")
	newobj4, _ := rdf.NewIRI(fmt.Sprintf("%s/prov#qu1", URI))
	newtriple4 := rdf.Triple{Subj: newsub, Pred: newpred4, Obj: newobj4}

	// assign qualifiedUsage 2
	newpred5, _ := rdf.NewIRI("http://www.w3.org/ns/prov#qualifiedUsage")
	newobj5, _ := rdf.NewIRI(fmt.Sprintf("%s/prov#qu2", URI))
	newtriple5 := rdf.Triple{Subj: newsub, Pred: newpred5, Obj: newobj5}

	// expand qualifiedUsage 1
	newpred6, _ := rdf.NewIRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type")
	newobj6, _ := rdf.NewIRI("http://www.w3.org/ns/prov#Usage")
	newtriple6 := rdf.Triple{Subj: newobj4, Pred: newpred6, Obj: newobj6}

	newpred7, _ := rdf.NewIRI("http://www.w3.org/ns/prov#entity")
	newobj7, _ := rdf.NewIRI(fmt.Sprintf("%s.csv", URI))
	newtriple7 := rdf.Triple{Subj: newobj4, Pred: newpred7, Obj: newobj7}

	newpred8, _ := rdf.NewIRI("http://www.w3.org/ns/prov#hasRole")
	newobj8, _ := rdf.NewIRI("http://www.w3.org/ns/csvw#csvEncodedTabularData")
	newtriple8 := rdf.Triple{Subj: newobj4, Pred: newpred8, Obj: newobj8}

	// expand qualifiedUsage 2
	newpred9, _ := rdf.NewIRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type")
	newobj9, _ := rdf.NewIRI("http://www.w3.org/ns/prov#Usage")
	newtriple9 := rdf.Triple{Subj: newobj5, Pred: newpred9, Obj: newobj9}

	newpred10, _ := rdf.NewIRI("http://www.w3.org/ns/prov#entity")
	newobj10, _ := rdf.NewIRI(fmt.Sprintf("%s/CSV", URI))
	newtriple10 := rdf.Triple{Subj: newobj5, Pred: newpred10, Obj: newobj10}

	newpred11, _ := rdf.NewIRI("http://www.w3.org/ns/prov#hasRole")
	newobj11, _ := rdf.NewIRI("http://www.w3.org/ns/csvw#tabularMetadata")
	newtriple11 := rdf.Triple{Subj: newobj5, Pred: newpred11, Obj: newobj11}

	tr = append(tr, newtriple0)
	tr = append(tr, newtriple1)
	tr = append(tr, newtriple2)
	tr = append(tr, newtriple3)
	tr = append(tr, newtriple4)
	tr = append(tr, newtriple5)
	tr = append(tr, newtriple6)
	tr = append(tr, newtriple7)
	tr = append(tr, newtriple8)
	tr = append(tr, newtriple9)
	tr = append(tr, newtriple10)
	tr = append(tr, newtriple11)

	// Create the output file
	outFile, err := os.Create("provTest.nt")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	// Write triples to a file
	var inoutFormat rdf.Format
	inoutFormat = rdf.NTriples // Turtle NQuads
	enc := rdf.NewTripleEncoder(outFile, inoutFormat)
	err = enc.EncodeAll(tr)
	// err = enc.Encode(newtriple)
	enc.Close()
	if err != nil {
		log.Fatal(err)
	}

}

func dataSetProv() string {
	// set up some of our boiler plate schema.org/Dataset elements
	// need date publishedOn, URL, lat long

	// geodata := Geo{Type: "GeoCoordinates", Latitude: latitude, Longitude: longitude}
	// spatial := Spatial{Type: "Place", Geo: geodata}
	// timenow := time.Now().Format(time.RFC850)
	// distribution := Distribution{Type: "DataDownload", ContentURL: uri, DatePublished: timenow, EncodingFormat: "text/tab-separated-values", InLanguage: "en"}
	// author := Author{Type: "Organization", Description: "NSF funded International Ocean Discovery Program operated by JRSO", Name: "International Ocean Discovery Program", URL: "http://iodp.org"}

	// // contextArray := []interface{"http://schema.org", {"glview": "http://schema.geolink.org/somethingIforgot"}}
	// kewords := fmt.Sprintf("DSDP, OPD, IODP, %s", measurement)

	schemametadata := ProvESDataset{Title: "this is the title", Level: "this is the level"}
	// schemametadata := SchemaOrgMetadata{Context:  ["http://schema.org", {"glview": "http://schema.geolink.org/somethingIforgot"}], Type: "Dataset"}

	schemaorgJSON, err := json.MarshalIndent(schemametadata, "", " ")

	if err != nil {
		log.Fatalf("JSON not encoded %v\n", err)
	}

	// JSON-LD parse snipit
	// dataparsed, _ := jsonld.ParseDataset(schemaorgJSON)
	// fmt.Printf("Serialized:\n %s \n\n", dataparsed.Serialize())

	return string(schemaorgJSON)
}
