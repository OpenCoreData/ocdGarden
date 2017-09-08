package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/deiu/rdf2go"
	"github.com/kazarena/json-gold/ld"
	tstore "github.com/wallix/triplestore"
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
		  "opencore":"http://opencore.org/voc/1",
		  "glview":"http://glview.org/voc/1/",
		  "@vocab":"http://schema.org/"
		},
	   "@type": "Dataset",
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
		"geo": {
		 "@type": "GeoCoordinates",
		 "latitude": "-21.01",
		 "longitude": "152.86"
		}
	   },
	   "url": "http://opencoredata.org/id/dataset/e8fb758e-22ba-499d-92fb-8d653febcf28"
	  }`

	const jsld2 = `{
		"@context": {
		  "opencore":"http://opencore.org/voc/1",
		  "glview":"http://glview.org/voc/1/",
		  "@vocab":"http://schema.org/"
		},
	   "@type": "Dataset",
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
		"geo": {
		 "@type": "GeoCoordinates",
		 "latitude": "-21.01",
		 "longitude": "152.86"
		}
	   },
	   "url": "http://opencoredata.org/id/dataset/e8fb758e-22ba-499d-92fb-8d653febcf28"
	  }`
	// ts()

	fmt.Println(jsonLDToRDF(jsld))

	// trying rdf2go
	baseUri := "https://example.org/foo" // Set a base URI
	g := rdf2go.NewGraph(baseUri)        // Create a new graph
	r := strings.NewReader(jsld)
	r2 := strings.NewReader(jsld2)
	g.Parse(r, "application/ld+json")  // r is an io.Reader
	g.Parse(r2, "application/ld+json") // r is an io.Reader

	fmt.Print(g.String())
}

// Trys to take a simple JSON-LD string and process it.
func jsonLDToRDF(jsonld string) string {

	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	options.Format = "application/nquads"

	var myInterface interface{}
	err := json.Unmarshal([]byte(jsonld), &myInterface)
	if err != nil {
		log.Println("Error when transforming JSON-LD document to interface:", err)
	}

	triples, err := proc.ToRDF(myInterface, options)
	if err != nil {
		log.Println("Error when transforming JSON-LD document to RDF:", err)
		return err.Error()
	}

	return triples.(string)
}

func ts() {
	fmt.Println("RDF Blank Node testing from JSON-LD")

	triples := tstore.Triples{}
	triples = append(triples,
		tstore.SubjPred("me", "name").StringLiteral("jsmith"),
		tstore.SubjPred("me", "age").IntegerLiteral(26),
		tstore.SubjPred("me", "male").BooleanLiteral(true),
		tstore.SubjPred("me", "born").DateTimeLiteral(time.Now()),
		tstore.SubjPred("me", "mother").Resource("mum#121287"),
	)

	triples = append(triples,
		tstore.SubjPred("me", "name").Bnode("jsmith"),
		tstore.BnodePred("me", "name").StringLiteral("jsmith"),
		tstore.SubjPred("me", "name").StringLiteralWithLang("jsmith", "en"),
	)

	triples = append(triples,
		tstore.SubjPred("me", "name").Bnode("jsmith"),
		tstore.BnodePred("me", "name").StringLiteral("jsmith"),
		tstore.SubjPred("me", "name").StringLiteralWithLang("jsmith", "en"),
	)

	test := tstore.SubjPred("me", "name").Bnode("randomUUID")
	nt := tstore.NewLenientNTEncoder(os.Stdout)
	nt.Encode(test)

	fmt.Print(triples.String())
}
