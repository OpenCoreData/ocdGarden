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

	// abstracts(session)
	// csvwmeta(session)
	// schemaorg(session)
	featuresAbsGeoJSON(session)

}

func featuresAbsGeoJSON(session *mgo.Session) {

	// connect and get documents
	csvw := session.DB("expedire").C("featuresAbsGeoJSON")
	var features []ocdstructs.ExpeditionGeoJSON
	err := csvw.Find(nil).All(&features)
	if err != nil {
		fmt.Printf("this is error %v \n", err)
	}

	// RDF item
	tr := []rdf.Triple{}

	// Loop on documents
	for _, item := range features {

		if item.Type != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/id/voc/janus/v1/type", item.Type))
		}
		if item.Hole != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/id/voc/janus/v1/hole", item.Hole))
		}
		if item.Expedition != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/id/voc/janus/v1/expedition", item.Expedition))
		}
		if item.Site != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/id/voc/janus/v1/site", item.Site))
		}
		if item.Program != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/id/voc/janus/v1/program", item.Program))
		}
		if item.Waterdepth != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/id/voc/janus/v1/waterdepth", item.Waterdepth))
		}
		if item.CoreCount != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/id/voc/janus/v1/corecount", item.CoreCount))
		}
		if item.Initialreportvolume != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/id/voc/janus/v1/initialreportvolume", item.Initialreportvolume))
		}
		if item.Coredata != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/id/voc/janus/v1/coredata", item.Coredata))
		}
		if item.Logdata != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/id/voc/janus/v1/logdata", item.Logdata))
		}
		if item.Geom != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/id/voc/janus/v1/geom", item.Geom))
		}
		if item.Scientificprospectus != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/id/voc/janus/v1/scientificprospectus", item.Scientificprospectus))
		}
		if item.CoreRecovery != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/id/voc/janus/v1/corerecovery", item.CoreRecovery))
		}
		if item.Penetration != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/id/voc/janus/v1/penetration", item.Penetration))
		}
		if item.Scientificreportvolume != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/id/voc/janus/v1/scientificreportvolume", item.Scientificreportvolume))
		}
		if item.Expeditionsite != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/id/voc/janus/v1/expeditionsite", item.Expeditionsite))
		}
		if item.Preliminaryreport != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/id/voc/janus/v1/preliminaryreport", item.Preliminaryreport))
		}
		if item.CoreInterval != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/id/voc/janus/v1/coreinterval", item.CoreInterval))
		}
		if item.PercentRecovery != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/id/voc/janus/v1/percentrecovery", item.PercentRecovery))
		}
		if item.Drilled != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/id/voc/janus/v1/drilled", item.Drilled))
		}
		if item.Vcdata != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/id/voc/janus/v1/vcdata", item.Vcdata))
		}
		if item.Note != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/id/voc/janus/v1/note", item.Note))
		}
		if item.Prcoeedingreport != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/id/voc/janus/v1/prcoeedingreport", item.Prcoeedingreport))
		}
		if item.Abstract != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/id/voc/janus/v1/abstract", item.Abstract))
		}

	}

	writeFile("./output/featuresAbsGeoJSON.nt", tr)

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

		tr = append(tr, SPOLiteral(item.URL, "http://opencoredata.org/id/voc/janus/v1/abstract", item.OpenCoreSite))

	}

	writeFile("./output/schemaorg.nt", tr)

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

	writeFile("./output/csvw.nt", tr)

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

	writeFile("./output/abstracts.nt", tr)

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

// SPOIRI return a new triple composed or two IRI's
func SPOIRI(subj, pred, obj string) rdf.Triple {

	newsub, err := rdf.NewIRI(subj)
	newpred0, err := rdf.NewIRI(pred)
	newobj0, err := rdf.NewIRI(obj)
	newtriple0 := rdf.Triple{Subj: newsub, Pred: newpred0, Obj: newobj0}

	if err != nil {
		log.Printf("this is error %v \n", err)
	}

	return newtriple0
}

// SPOLiteral return a new triple with literial object
func SPOLiteral(subj, pred, obj string) rdf.Triple {

	newsub, err := rdf.NewIRI(subj)
	newpred0, err := rdf.NewIRI(pred)
	newobj0, err := rdf.NewLiteral(obj)
	newtriple0 := rdf.Triple{Subj: newsub, Pred: newpred0, Obj: newobj0}

	if err != nil {
		log.Printf("this is error %v \n", err)
	}

	return newtriple0
}
