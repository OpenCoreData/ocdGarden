package main

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/kazarena/json-gold/ld"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	structs "opencoredata.org/ocdGarden/jsongold/structs"
)

func main() {
	// restTest()  // works fine
	stringTest() // working on this to use in the mongo mongoTest
	// mongoTest() // waiting to resolve the stringTest() call
}

// Calls the REST API to get a JSON-LD document and then that document is
//  1) converted to RDF
//  2) attempted to convert back to JSON-LD
// NOTE:  A newer version of the  JSON-LD is coming that will have no blank nodes
func restTest() {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	options.Format = "application/nquads"

	// triples, err := proc.ToRDF("http://localhost/api/v1/documents/download/e8fb758e-22ba-499d-92fb-8d653febcf28/JSON", options)
	triples, err := proc.ToRDF("http://localhost/api/v1/documents/download/218fda28-6763-470f-b8ba-6f3350e26fde/JSON", options)
	if err != nil {
		log.Println("Error when transforming JSON-LD document to RDF:", err)
		return
	}
	fmt.Println(triples) // os.Stdout.WriteString(triples.(string))

	// from RDF to JSON-LD test
	proc2 := ld.NewJsonLdProcessor()
	options2 := ld.NewJsonLdOptions("")
	doc, err := proc2.FromRDF(triples, options2)
	expanded, err := proc.Compact(doc, nil, options)
	ld.PrintDocument("JSON-LD expansion succeeded", expanded)
}

// Trys to take a simple JSON-LD string and process it. Having trouble here.. a git issue made
// in the jsongold repo...
func stringTest() {

	jsld := `{
  "@context": {
    "opencore":"http://opencore.org/",
    "glview":"http://glview.org/",
    "@vocab":"http://schema.org/"
  },
 "@type": "Dataset",
 "@id" : "http://opencoredata.org/id/dataset/e8fb758e-22ba-499d-92fb-8d653febcf28",
 "author": {
  "@id" : "http://doi.org/10.17616/R31W2T",
  "@type": "Organization",
  "description": "NSF funded International Ocean Discovery Program operated by JRSO",
  "name": "International Ocean Discovery Program",
  "url": "http://iodp.org"
 },
 "description": "Data set description",
 "distribution": {
  "@type": "DataDownload",
   "@id" : "http://opencoredata.org/id/dataset/e8fb758e-22ba-499d-92fb-8d653febcf28#distrubtion",
  "contentUrl": "http://opencoredata.org/id/dataset/e8fb758e-22ba-499d-92fb-8d653febcf28",
  "datePublished": "Sunday, 25-Oct-15 15:00:28 EDT",
  "encodingFormat": "text/tab-separated-values",
  "inLanguage": "en"
 },
 "glview:dataset": "194_1196B_JanusChemCarb_aZgZxnNI.csv",
 "glview:keywords": "DSDP, OPD, IODP, JanusChemCarb",
 "opencore:leg": "194",
 "opencore:site": "1196",
 "opencore:hole": "B",
 "opencore:measurement": "JanusChemCarb",
 "keywords": "DSDP, OPD, IODP, JanusChemCarb",
 "name": "194_1196B_JanusChemCarb_aZgZxnNI.csv",
 "spatial": {
  "@type": "Place",
    "@id" : "http://opencoredata.org/id/dataset/e8fb758e-22ba-499d-92fb-8d653febcf28#place",
  "geo": {
   "@type": "GeoCoordinates",
     "@id" : "http://opencoredata.org/id/dataset/e8fb758e-22ba-499d-92fb-8d653febcf28#geocoordinates",
   "latitude": "-21.01",
   "longitude": "152.86"
  }
 },
 "url": "http://opencoredata.org/id/dataset/e8fb758e-22ba-499d-92fb-8d653febcf28"
}`

	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")

	var myInterface interface{}
	err := json.Unmarshal([]byte(jsld), &myInterface)
	if err != nil {
		log.Println("Error when transforming JSON-LD document to interface:", err)
	}

	triples, err := proc.ToRDF(myInterface, options)
	fmt.Println(triples)

}

// This isn't working yet..  the JSON-LD in here has some errors (they can be fixed)
// Goal is to fix them, process through jsongold to make sure we are complient and
// then store the triples in a triple store for later extraction and use
func mongoTest() {
	session, err := getMongoCon()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// connect and get documents
	c := session.DB("test").C("schemaorg")
	result := structs.SchemaOrgMetadata{} // need this struct  (it's everywhere.   what can I do about that?   move only ocdServices from ocdWeb?)
	err = c.Find(bson.M{"url": "http://opencoredata.org/id/dataset/da39147d-deda-44ac-879d-684491a110fe"}).One(&result)
	if err != nil {
		log.Printf("URL lookup error: %v", err)
	}

	// context setting hack
	// result.Context = ` "opencore": "http://opencoredata.org/voc/1/", "glview": "http://geolink.org/view/1/", "schema": "http://schema.org/"`
	result.Context = "http://schema.org"
	// jsonldtext, _ := json.MarshalIndent(result, "", " ") // results as embeddale JSON-LD
	if err != nil {
		log.Printf("Error calling in GetFileBuyUUID : %v ", err)
	}
}

func getMongoCon() (*mgo.Session, error) {
	host := os.Getenv("localhost")
	return mgo.Dial(host)
}
