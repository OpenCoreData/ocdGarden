package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"

	rdf "github.com/knakk/rdf"
	"gopkg.in/mgo.v2"
	ocdstructs "opencoredata.org/ocdGarden/MongoUtils/Mongo2RDF/structs"
)

func main() {
	// call mongo and lookup the redirection to use...
	session, err := getMongoCon()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	abstracts(session)
	// csvwmeta(session)  // change this to use jsongold approach
	// schemaorg(session) // change this to use jsongold approach
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

		SPOIRI(item.Uri, "a", "http://opencoredata.org/voc/janus/v1/Feature")

		if item.Type != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Type", item.Type))
		}
		if item.Hole != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Hole", item.Hole))
		}
		if item.Expedition != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Expedition", item.Expedition))
		}
		if item.Site != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Site", item.Site))
		}
		if item.Program != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Program", item.Program))
		}
		if item.Waterdepth != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Waterdepth", item.Waterdepth))
		}
		if item.CoreCount != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Corecount", item.CoreCount))
		}
		if item.Initialreportvolume != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Initialreportvolume", item.Initialreportvolume))
		}
		if item.Coredata != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Coredata", item.Coredata))
		}
		if item.Logdata != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Logdata", item.Logdata))
		}
		if item.Geom != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Geom", item.Geom))
		}
		if item.Scientificprospectus != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Scientificprospectus", item.Scientificprospectus))
		}
		if item.CoreRecovery != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Corerecovery", item.CoreRecovery))
		}
		if item.Penetration != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Penetration", item.Penetration))
		}
		if item.Scientificreportvolume != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Scientificreportvolume", item.Scientificreportvolume))
		}
		if item.Expeditionsite != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Expeditionsite", item.Expeditionsite))
		}
		if item.Preliminaryreport != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Preliminaryreport", item.Preliminaryreport))
		}
		if item.CoreInterval != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Coreinterval", item.CoreInterval))
		}
		if item.PercentRecovery != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Percentrecovery", item.PercentRecovery))
		}
		if item.Drilled != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Drilled", item.Drilled))
		}
		if item.Vcdata != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/VCdata", item.Vcdata))
		}
		if item.Note != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Note", item.Note))
		}
		if item.Prcoeedingreport != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Prcoeedingreport", item.Prcoeedingreport))
		}
		if item.Abstract != "" {
			tr = append(tr, SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Abstract", stripCtlAndExtFromUnicode(item.Abstract)))
		}

	}

	writeFile("./output/janusAbstracts.nt", tr)

}

// update with material from jsongold
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

		// need the mapping like in:
		// 	if item.Title != "" {
		// 	tr = append(tr, SPOLiteral(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/title", stripCtlAndExtFromUnicode(item.Title)))
		// }

		tr = append(tr, SPOLiteral(item.URL, "http://opencoredata.org/id/voc/janus/v1/abstract", item.OpenCoreSite))

	}

	writeFile("./output/schemaorg.nt", tr)

}

// update with material from jsongold
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
	var csdcoAbs []ocdstructs.MdocsV2
	err := csvw.Find(nil).All(&csdcoAbs)
	if err != nil {
		fmt.Printf("this is error %v \n", err)
	}

	// RDF item
	tr := []rdf.Triple{}

	// Loop on documents
	for _, item := range csdcoAbs {
		// Make subject IRI
		// newsub, _ := rdf.NewIRI(fmt.Sprintf("http://opencoredata.org/id/resource/janus/query/%s", item.ID)) // Sprintf a correct URI here

		// title
		// newpred0, _ := rdf.NewIRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#title")
		// newobj0, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/JanusQuery")
		// newtriple0 := rdf.Triple{Subj: newsub, Pred: newpred0, Obj: newobj0}
		// tr = append(tr, newtriple0)

		abstractIRI := fmt.Sprintf("http://opencoredata.org/id/resource/csdco/abstract/%s", item.ID)

		// need a TYPE!
		SPOIRI(abstractIRI, "a", "http://opencoredata.org/id/voc/csdco/v1/Abstract")

		if item.Title != "" {
			tr = append(tr, SPOLiteral(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/title", stripCtlAndExtFromUnicode(item.Title)))
		}
		if item.Type != "" {
			tr = append(tr, SPOLiteral(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/type", item.Type))
		}
		if item.Year != 0 {
			tr = append(tr, SPOLiteral(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/year", strconv.Itoa(item.Year)))
		}
		if item.Source != "" {
			tr = append(tr, SPOLiteral(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/source", item.Source))
		}
		if item.ID != "" {
			tr = append(tr, SPOLiteral(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/id", item.ID))
		}
		if !item.Created.IsZero() {
			tr = append(tr, SPOLiteral(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/created", item.Created.String()))
		}
		if item.Profile_ID != "" {
			tr = append(tr, SPOLiteral(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/profileid", item.Profile_ID))
		}
		if item.Group_ID != "" {
			tr = append(tr, SPOLiteral(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/groupid", item.Group_ID))
		}
		if !item.Last_Modified.IsZero() {
			tr = append(tr, SPOLiteral(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/lastmodified", item.Last_Modified.String()))
		}
		if item.Abstract != "" {
			tr = append(tr, SPOLiteral(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/abstracttext", stripCtlAndExtFromUnicode(item.Abstract)))
		}
		if item.Identifiers.Doi != "" {
			tr = append(tr, SPOLiteral(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/doi", item.Identifiers.Doi))
		}
		if item.Identifiers.Issn != "" {
			tr = append(tr, SPOLiteral(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/issn", item.Identifiers.Issn))
		}

		// loop on Tags
		if len(item.Tags) > 0 {
			tagIRI := fmt.Sprintf("http://opencoredata.org/id/resource/csdco/abstract/%s#tag", item.ID)
			tr = append(tr, SPOIRI(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/tag", tagIRI))
			for _, tag := range item.Tags {
				tr = append(tr, SPOLiteral(tagIRI, "http://opencoredata.org/id/voc/csdco/v1/tag/value", stripCtlAndExtFromUnicode(tag)))
			}
		}

		// loop on Authors
		if len(item.Authors) > 0 {
			authorIRI := fmt.Sprintf("http://opencoredata.org/id/resource/csdco/abstract/%s#authors", item.ID)
			tr = append(tr, SPOIRI(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/authors", authorIRI))
			for _, author := range item.Authors {
				tr = append(tr, SPOLiteral(authorIRI, "http://opencoredata.org/id/voc/csdco/v1/author/firstname", stripCtlAndExtFromUnicode(author.First_Name)))
				tr = append(tr, SPOLiteral(authorIRI, "http://opencoredata.org/id/voc/csdco/v1/author/lastname", stripCtlAndExtFromUnicode(author.Last_Name)))
			}
		}

		// for _, author := range item.Authors {
		// 	fmt.Printf("FirstName: %s\n", author.First_Name)
		// 	fmt.Printf("LastName: %s\n", author.Last_Name)
		// }

	}

	writeFile("./output/csdcoAbstracts.nt", tr)

}

// ref:  https://rosettacode.org/wiki/Strip_control_codes_and_extended_characters_from_a_string#Go
func stripCtlAndExtFromUnicode(str string) string {
	isOk := func(r rune) bool {
		return r < 32 || r >= 127
	}
	// The isOk filter is such that there is no need to chain to norm.NFC
	t := transform.Chain(norm.NFKD, transform.RemoveFunc(isOk))
	// This Transformer could also trivially be applied as an io.Reader
	// or io.Writer filter to automatically do such filtering when reading
	// or writing data anywhere.
	str, _, _ = transform.String(t, str)
	return str
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
	inoutFormat = rdf.NTriples // Turtle NQuads Ntriples
	enc := rdf.NewTripleEncoder(outFile, inoutFormat)
	err = enc.EncodeAll(tr)
	// err = enc.Encode(newtriple)
	enc.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func getMongoCon() (*mgo.Session, error) {
	host := os.Getenv("opencore.dev")
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
