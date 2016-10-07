package main

import (
	"fmt"
	"log"
	"os"

	rdf "github.com/knakk/rdf"
	"gopkg.in/mgo.v2"
	ocdstructs "opencoredata.org/ocdGarden/Mongo2RDF/structs"
)

func main() {
	// call mongo and lookup the redirection to use...
	session, err := getMongoCon()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	abstracts(session)
	csvwmeta(session)
	schemaorg(session)
	featuresAbsGeoJSON(session)

}

func featuresAbsGeoJSON(session *mgo.Session) {

	// connect and get documents
	csvw := session.DB("test").C("schemaorg")
	var features []ocdstructs.ExpeditionGeoJSON
	err := csvw.Find(nil).All(&features)
	if err != nil {
		fmt.Printf("this is error %v \n", err)
	}

	// RDF item
	tr := []rdf.Triple{}

	// Loop on documents
	for _, item := range features {
		newsub, _ := rdf.NewIRI(fmt.Sprintf("http://opencoredata/id/resource/janus/query/%s", item.Uri)) // Sprintf a correct URI here

		// title
		newpred0, _ := rdf.NewIRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#title")
		newobj0, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/JanusQuery")
		newtriple0 := rdf.Triple{Subj: newsub, Pred: newpred0, Obj: newobj0}
		tr = append(tr, newtriple0)

	}

	writeFile("featuresAbsGeoJSON.nt", tr)

}

func schemaorg(session *mgo.Session) {

	// connect and get documents
	csvw := session.DB("test").C("schemaorg")
	var schemaDocs []ocdstructs.SchemaOrgMetadata
	err := csvw.Find(nil).All(&schemaDocs)
	if err != nil {
		fmt.Printf("this is error %v \n", err)
	}

	// RDF item
	tr := []rdf.Triple{}

	// Loop on documents
	for _, item := range schemaDocs {
		newsub, _ := rdf.NewIRI(fmt.Sprintf("http://opencoredata/id/resource/janus/query/%s", item.URL)) // Sprintf a correct URI here

		// title
		newpred0, _ := rdf.NewIRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#title")
		newobj0, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/JanusQuery")
		newtriple0 := rdf.Triple{Subj: newsub, Pred: newpred0, Obj: newobj0}
		tr = append(tr, newtriple0)

	}

	writeFile("schemaorg.nt", tr)

}

func csvwmeta(session *mgo.Session) {

	// connect and get documents
	csvw := session.DB("test").C("csvwmeta")
	var csvwDocs []ocdstructs.CSVWMeta
	err := csvw.Find(nil).All(&csvwDocs)
	if err != nil {
		fmt.Printf("this is error %v \n", err)
	}

	// RDF item
	tr := []rdf.Triple{}

	// Loop on documents
	for _, item := range csvwDocs {
		newsub, _ := rdf.NewIRI(fmt.Sprintf("http://opencoredata/id/resource/janus/query/%s", item.URL)) // Sprintf a correct URI here

		// title
		newpred0, _ := rdf.NewIRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#title")
		newobj0, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/JanusQuery")
		newtriple0 := rdf.Triple{Subj: newsub, Pred: newpred0, Obj: newobj0}
		tr = append(tr, newtriple0)

	}

	writeFile("csvw.nt", tr)

}

func abstracts(session *mgo.Session) {

	// connect and get documents
	csvw := session.DB("abstracts").C("csdco")
	var csdcoAbs []ocdstructs.Mdocs
	err := csvw.Find(nil).All(&csdcoAbs)
	if err != nil {
		fmt.Printf("this is error %v \n", err)
	}

	// RDF item
	tr := []rdf.Triple{}

	// Loop on documents
	for _, item := range csdcoAbs {
		// Make subject IRI
		newsub, _ := rdf.NewIRI(fmt.Sprintf("http://opencoredata/id/resource/janus/query/%s", item.ID)) // Sprintf a correct URI here

		// title
		newpred0, _ := rdf.NewIRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#title")
		newobj0, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/JanusQuery")
		newtriple0 := rdf.Triple{Subj: newsub, Pred: newpred0, Obj: newobj0}
		tr = append(tr, newtriple0)

		// for _, author := range item.Authors {
		// 	fmt.Printf("FirstName: %s\n", author.First_Name)
		// 	fmt.Printf("LastName: %s\n", author.Last_Name)
		// }

	}

	writeFile("abstracts.nt", tr)

}

func writeFile(name string, tr []rdf.Triple) {
	// Create the output file
	outFile, err := os.Create(name)
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

func getMongoCon() (*mgo.Session, error) {
	host := os.Getenv("localhost")
	return mgo.Dial(host)
}
