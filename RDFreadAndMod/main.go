package main

import (
	"fmt"
	"log"
	"os"

	rdf "github.com/knakk/rdf"
	"gopkg.in/mgo.v2"
	ocdCommons "opencoredata.org/ocdCommons/structs"
)

type CruiseGL struct {
	Expedition    string `json:"expedition"`
	Cruisetype    string `json:"cruisetype"`
	Endportcall   string `json:"endportcall"`
	Operator      string `json:"operator"`
	Participant   string `json:"participant"`
	Program       string `json:"program"`
	Scheduler     string `json:"scheduler"`
	Startportcall string `json:"startportcall"`
	Legsitehole   string `json:"legsitehole"`
	Track         string `json:"track"`
	Vessel        string `json:"vessel"`
	Note          string `json:"note"`
	Uri           string `json:"uri"`
}

func main() {
	fmt.Println("Read and write RDF")

	// Open the input file
	inFile, err := os.Open("/Users/dfils/src/go/src/opencoredata.org/ocdDataStores/tripleStore/data/codices.nt")
	defer inFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Decode the existing triples
	var inoutFormat rdf.Format
	inoutFormat = rdf.FormatNT // FormatTTL
	dec := rdf.NewTripleDecoder(inFile, inoutFormat)
	tr, err := dec.DecodeAll()

	// call mongo and lookup the redirection to use...
	session, err := GetMongoCon()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Find all expedire - features
	collection := session.DB("expedire").C("features")
	var collDoc []ocdCommons.Expedition
	err = collection.Find(nil).All(&collDoc)
	if err != nil {
		fmt.Printf("Mongo search error: %v \n", err)
	}

	for _, item := range collDoc {
		newsub, _ := rdf.NewIRI(item.Uri)


// I need one or MORE type predicates for these resources to type them
// to geolink feature?  cruise?   and also as type resource?   
// Look at the geolink graphs I have made to see what I did there.

		newpred1, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/uri")
		newpred2, _ := rdf.NewIRI("http://www.w3.org/2003/01/geo/wgs84_pos#lat")
		newpred3, _ := rdf.NewIRI("http://www.w3.org/2003/01/geo/wgs84_pos#long")
		newpred4, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/hole")
		newpred5, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/leg")
		newpred6, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/site")
		newpred7, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/program")
		newpred8, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/waterdepth")
		newpred9, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/coreCount")
		newpred10, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/initialreportvolume")
		newpred11, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/coredata")
		newpred12, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/logdata")
		newpred13, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/geom")
		newpred14, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/scientificprospectus")
		newpred15, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/coreRecovery")
		newpred16, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/penetration")
		newpred17, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/scientificreportvolume")
		newpred18, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/expeditionsite")
		newpred19, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/preliminaryreport")
		newpred20, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/coreInterval")
		newpred21, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/percentRecovery")
		newpred22, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/drilled")
		newpred23, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/vcdata")
		newpred24, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/note")
		newpred25, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/prcoeedingreport")

		newobj1, _ := rdf.NewLiteral(item.Uri)
		newobj2, _ := rdf.NewLiteral(item.Lat)
		newobj3, _ := rdf.NewLiteral(item.Long)
		newobj4, _ := rdf.NewLiteral(item.Hole)
		newobj5, _ := rdf.NewLiteral(item.Expedition)
		newobj6, _ := rdf.NewLiteral(item.Site)
		newobj7, _ := rdf.NewLiteral(item.Program)
		newobj8, _ := rdf.NewLiteral(item.Waterdepth)
		newobj9, _ := rdf.NewLiteral(item.CoreCount)
		newobj10, _ := rdf.NewLiteral(item.Initialreportvolume)
		newobj11, _ := rdf.NewLiteral(item.Coredata)
		newobj12, _ := rdf.NewLiteral(item.Logdata)
		newobj13, _ := rdf.NewLiteral(item.Geom)
		newobj14, _ := rdf.NewLiteral(item.Scientificprospectus)
		newobj15, _ := rdf.NewLiteral(item.CoreRecovery)
		newobj16, _ := rdf.NewLiteral(item.Penetration)
		newobj17, _ := rdf.NewLiteral(item.Scientificreportvolume)
		newobj18, _ := rdf.NewLiteral(item.Expeditionsite)
		newobj19, _ := rdf.NewLiteral(item.Preliminaryreport)
		newobj20, _ := rdf.NewLiteral(item.CoreInterval)
		newobj21, _ := rdf.NewLiteral(item.PercentRecovery)
		newobj22, _ := rdf.NewLiteral(item.Drilled)
		newobj23, _ := rdf.NewLiteral(item.Vcdata)
		newobj24, _ := rdf.NewLiteral(item.Note)
		newobj25, _ := rdf.NewLiteral(item.Prcoeedingreport)

		newtriple1 := rdf.Triple{Subj: newsub, Pred: newpred1, Obj: newobj1}
		newtriple2 := rdf.Triple{Subj: newsub, Pred: newpred2, Obj: newobj2}
		newtriple3 := rdf.Triple{Subj: newsub, Pred: newpred3, Obj: newobj3}
		newtriple4 := rdf.Triple{Subj: newsub, Pred: newpred4, Obj: newobj4}
		newtriple5 := rdf.Triple{Subj: newsub, Pred: newpred5, Obj: newobj5}
		newtriple6 := rdf.Triple{Subj: newsub, Pred: newpred6, Obj: newobj6}
		newtriple7 := rdf.Triple{Subj: newsub, Pred: newpred7, Obj: newobj7}
		newtriple8 := rdf.Triple{Subj: newsub, Pred: newpred8, Obj: newobj8}
		newtriple9 := rdf.Triple{Subj: newsub, Pred: newpred9, Obj: newobj9}
		newtriple10 := rdf.Triple{Subj: newsub, Pred: newpred10, Obj: newobj10}
		newtriple11 := rdf.Triple{Subj: newsub, Pred: newpred11, Obj: newobj11}
		newtriple12 := rdf.Triple{Subj: newsub, Pred: newpred12, Obj: newobj12}
		newtriple13 := rdf.Triple{Subj: newsub, Pred: newpred13, Obj: newobj13}
		newtriple14 := rdf.Triple{Subj: newsub, Pred: newpred14, Obj: newobj14}
		newtriple15 := rdf.Triple{Subj: newsub, Pred: newpred15, Obj: newobj15}
		newtriple16 := rdf.Triple{Subj: newsub, Pred: newpred16, Obj: newobj16}
		newtriple17 := rdf.Triple{Subj: newsub, Pred: newpred17, Obj: newobj17}
		newtriple18 := rdf.Triple{Subj: newsub, Pred: newpred18, Obj: newobj18}
		newtriple19 := rdf.Triple{Subj: newsub, Pred: newpred19, Obj: newobj19}
		newtriple20 := rdf.Triple{Subj: newsub, Pred: newpred20, Obj: newobj20}
		newtriple21 := rdf.Triple{Subj: newsub, Pred: newpred21, Obj: newobj21}
		newtriple22 := rdf.Triple{Subj: newsub, Pred: newpred22, Obj: newobj22}
		newtriple23 := rdf.Triple{Subj: newsub, Pred: newpred23, Obj: newobj23}
		newtriple24 := rdf.Triple{Subj: newsub, Pred: newpred24, Obj: newobj24}
		newtriple25 := rdf.Triple{Subj: newsub, Pred: newpred25, Obj: newobj25}

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
		tr = append(tr, newtriple12)
		tr = append(tr, newtriple13)
		tr = append(tr, newtriple14)
		tr = append(tr, newtriple15)
		tr = append(tr, newtriple16)
		tr = append(tr, newtriple17)
		tr = append(tr, newtriple18)
		tr = append(tr, newtriple19)
		tr = append(tr, newtriple20)
		tr = append(tr, newtriple21)
		tr = append(tr, newtriple22)
		tr = append(tr, newtriple23)
		tr = append(tr, newtriple24)
		tr = append(tr, newtriple25)
	}

	// Find all expedire - expeditions
	collection2 := session.DB("expedire").C("expeditions")
	var expDocs []CruiseGL
	err = collection2.Find(nil).All(&expDocs)
	if err != nil {
		fmt.Printf("Mongo search error: %v \n", err)
	}

	for _, item := range expDocs {
		newsub, _ := rdf.NewIRI(item.Uri)

		newpred1, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/expedition")
		newpred2, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/cruisetype")
		newpred3, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/endportcall")
		newpred4, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/operator")
		newpred5, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/participant")
		newpred6, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/program")
		newpred7, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/scheduler")
		newpred8, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/startportcall")
		newpred9, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/legsitehole")
		newpred10, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/track")
		newpred11, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/vessel")
		newpred12, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/note")
		newpred13, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/uri")

		newobj1, _ := rdf.NewLiteral(item.Expedition)
		newobj2, _ := rdf.NewLiteral(item.Cruisetype)
		newobj3, _ := rdf.NewLiteral(item.Endportcall)
		newobj4, _ := rdf.NewLiteral(item.Operator)
		newobj5, _ := rdf.NewLiteral(item.Participant)
		newobj6, _ := rdf.NewLiteral(item.Program)
		newobj7, _ := rdf.NewLiteral(item.Scheduler)
		newobj8, _ := rdf.NewLiteral(item.Startportcall)
		newobj9, _ := rdf.NewLiteral(item.Legsitehole)
		newobj10, _ := rdf.NewLiteral(item.Track)
		newobj11, _ := rdf.NewLiteral(item.Vessel)
		newobj12, _ := rdf.NewLiteral(item.Note)
		newobj13, _ := rdf.NewLiteral(item.Uri)

		newtriple1 := rdf.Triple{Subj: newsub, Pred: newpred1, Obj: newobj1}
		newtriple2 := rdf.Triple{Subj: newsub, Pred: newpred2, Obj: newobj2}
		newtriple3 := rdf.Triple{Subj: newsub, Pred: newpred3, Obj: newobj3}
		newtriple4 := rdf.Triple{Subj: newsub, Pred: newpred4, Obj: newobj4}
		newtriple5 := rdf.Triple{Subj: newsub, Pred: newpred5, Obj: newobj5}
		newtriple6 := rdf.Triple{Subj: newsub, Pred: newpred6, Obj: newobj6}
		newtriple7 := rdf.Triple{Subj: newsub, Pred: newpred7, Obj: newobj7}
		newtriple8 := rdf.Triple{Subj: newsub, Pred: newpred8, Obj: newobj8}
		newtriple9 := rdf.Triple{Subj: newsub, Pred: newpred9, Obj: newobj9}
		newtriple10 := rdf.Triple{Subj: newsub, Pred: newpred10, Obj: newobj10}
		newtriple11 := rdf.Triple{Subj: newsub, Pred: newpred11, Obj: newobj11}
		newtriple12 := rdf.Triple{Subj: newsub, Pred: newpred12, Obj: newobj12}
		newtriple13 := rdf.Triple{Subj: newsub, Pred: newpred13, Obj: newobj13}

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
		tr = append(tr, newtriple12)
		tr = append(tr, newtriple13)

	}

	// // create a new triple.
	// newsub, _ := rdf.NewIRI("http://www.example.org/subj")
	// newpred, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/")
	// newobj, _ := rdf.NewLiteral("12")
	// newtriple := rdf.Triple{Subj: newsub, Pred: newpred, Obj: newobj}

	// // append new triple
	// tr = append(tr, newtriple)

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

func GetMongoCon() (*mgo.Session, error) {
	host := os.Getenv("MONGO_HOST")
	return mgo.Dial(host)
}
