package main

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/kazarena/json-gold/ld"
	"gopkg.in/mgo.v2"
	structs "opencoredata.org/ocdGarden/jsongold/structs"
)

func main() {

	const jsld = `{
  "@context": {
    "opencore":"http://opencore.org/voc/1",
    "glview":"http://glview.org/voc/1/",
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

	csvwTest()
	mongoTest() // waiting to resolve the stringTest() call
	// triples := stringTest(jsld) // working on this to use in the mongo mongoTest
	// fmt.Println(triples)
	// restTest() // works fine

}
func getMongoCon() (*mgo.Session, error) {
	host := os.Getenv("localhost")
	return mgo.Dial(host)
}

// TODO a version of mongoTest for the CSVW metadata...
// this is for the the schema.org stuff
func csvwTest() {
	session, err := getMongoCon()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// connect and get documents
	c := session.DB("test").C("csvwmeta")
	results := []structs.CSVWMeta{} // need this struct  (it's everywhere.   what can I do about that?   move only ocdServices from ocdWeb?)
	// err = c.Find(bson.M{"url": "http://opencoredata.org/id/dataset/da39147d-deda-44ac-879d-684491a110fe"}).One(&result)
	err = c.Find(nil).Limit(3).All(&results) // c.Find(nil).Limit(3).All(&results)
	if err != nil {
		log.Printf("URL lookup error: %v", err)
	}

	for _, result := range results {
		// context setting hack
		result.Context.Schema = "http://www.w3.org/ns/csvw/" // this "schema" is the @voc in the struct..  confusing when not using schema.org
		result.Context.OpenCore = "http://opencore.org/voc/1/"
		result.Context.GeoLink = "http://glview.org/voc/1/"
		result.Context.Base = fmt.Sprintf("%s", result.URL) //  %s.json
		result.Type = "http://www.w3.org/ns/dcat#DataSet"

		// really this should be URL + .json ?
		result.Id = fmt.Sprintf("%s", result.URL) //  %s.json

		// add dc_publisher ID  TODO, set to DOI of OCD
		//result.Dc_publisher.Id = fmt.Sprintf("%s%s", result.URL, "#distribution")

		result.Dc_license.Type = "http://purl.org/dc/terms/RightsStatement"

		result.TableSchema.Id = fmt.Sprintf("%s#tableSchema", result.URL) //  %s.json
		result.TableSchema.Type = "TableSchema"

		// add tableschema ID
		result.Dc_publisher.Schema_url.Id = "http://opencoredata.org/voc/1/janus/"
		result.Dc_publisher.Id = fmt.Sprintf("%s#publisher", result.URL) //  %s.json
		result.Dc_publisher.Type = "http://purl.org/dc/terms/Publisher"  // change to a better type..  like RE3 type?

		// loop on column range to add ID and TYPE
		for index, column := range result.TableSchema.Columns {
			result.TableSchema.Columns[index].Id = fmt.Sprintf("%s#%s", result.URL, column.Name) //  %s.json
			result.TableSchema.Columns[index].Type = "Column"
			// TODO..  make a call to the vobulary graph and populate the description here too
		}

		jsonldtext, _ := json.MarshalIndent(result, "", " ") // results as embeddale JSON-LD

		fmt.Println("jsonld text--------------------------------")

		fmt.Println(string(jsonldtext))

		fmt.Println("jsonLDToRDF--------------------------------")

		fmt.Println(jsonLDToRDF(string(jsonldtext)))

		fmt.Println("rdfToJSONLD--------------------------------")

		fmt.Println(rdfToJSONLD(jsonLDToRDF(string(jsonldtext))))

	}
}

// this is a test function to take the RDF triple stream and rebuild the JSON-LD backs
func rdfToJSONLD(nquads string) string {

	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	options.Format = "application/nquads"

	doc, err := proc.FromRDF(nquads, options)
	expanded, err := proc.Compact(doc, nil, options)
	if err != nil {
		log.Println("Error when transforming nquads document to JSONDLD:", err)
		return err.Error()
	}

	b, _ := json.MarshalIndent(expanded, "", "  ")

	return string(b)
}

func PrintDocument(msg string, doc interface{}) {
	b, _ := json.MarshalIndent(doc, "", "  ")
	if msg != "" {
		os.Stdout.WriteString(msg)
		os.Stdout.WriteString("\n")
	}
	os.Stdout.Write(b)
	os.Stdout.WriteString("\n")
}

// this is for the the schema.org stuff
func mongoTest() {
	session, err := getMongoCon()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// connect and get documents
	c := session.DB("test").C("schemaorg")
	results := []structs.SchemaOrgMetadata{} // need this struct  (it's everywhere.   what can I do about that?   move only ocdServices from ocdWeb?)
	// err = c.Find(bson.M{"url": "http://opencoredata.org/id/dataset/da39147d-deda-44ac-879d-684491a110fe"}).One(&result)
	err = c.Find(nil).Limit(3).All(&results) // c.Find(nil).Limit(3).All(&results)
	if err != nil {
		log.Printf("URL lookup error: %v", err)
	}

	for _, result := range results {
		// context setting hack
		result.Context.Schema = "http://glview.org/voc/1/"
		result.Context.OpenCore = "http://opencore.org/voc/1"
		result.Context.GeoLink = "http://schema.org/"

		result.ID = result.URL
		result.Distribution.ID = fmt.Sprintf("%s%s", result.URL, "#distribution")
		result.Author.ID = fmt.Sprintf("%s%s", result.URL, "#author")
		result.Spatial.ID = fmt.Sprintf("%s%s", result.URL, "#spatial")
		result.Spatial.Geo.ID = fmt.Sprintf("%s%s", result.URL, "#geo")
		if err != nil {
			log.Printf("Error calling in GetFileBuyUUID : %v ", err)
		}

		jsonldtext, _ := json.MarshalIndent(result, "", " ") // results as embeddale JSON-LD
		// fmt.Println(string(jsonldtext))
		fmt.Println(jsonLDToRDF(string(jsonldtext)))
	}
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
	fmt.Println(triples)

	// from RDF to JSON-LD test
	proc2 := ld.NewJsonLdProcessor()
	options2 := ld.NewJsonLdOptions("")
	doc, err := proc2.FromRDF(triples, options2)
	expanded, err := proc.Compact(doc, nil, options)
	if err != nil {
		log.Println("Error when transforming JSON-LD document to RDF:", err)
		return
	}
	ld.PrintDocument("JSON-LD expansion succeeded", expanded)
}
