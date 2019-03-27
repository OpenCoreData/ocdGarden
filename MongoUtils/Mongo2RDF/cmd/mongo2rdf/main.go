package main

import (
	"log"

	// "opencoredata.org/ocdGarden/CSDCO/VaultWalker/pkg/utils"

	"opencoredata.org/ocdGarden/MongoUtils/Mongo2RDF/internal/common"
	"opencoredata.org/ocdGarden/MongoUtils/Mongo2RDF/internal/pipes"
)

func main() {
	// call mongo and lookup the redirection to use...
	session, err := common.GetMongoCon()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	var cab common.Buffer
	l := pipes.NewAbstracts(session, &cab)
	log.Println(l)
	log.Println(cab.Len())
	common.WriteToName(cab.String(), "CSDCOAbstracts.nq")

	// New version abstract call
	var jab common.Buffer
	l = pipes.NewFeatureAbsGeoJSON(session, &jab)
	log.Println(l)
	log.Println(jab.Len())
	common.WriteToName(jab.String(), "JanusExp.nq")

	// New version abstract call
	// The CSVW stuff should be incorporated into this one
	var sdo common.Buffer
	l = pipes.Schemaorg(session, &sdo)
	log.Println(l)
	log.Println(sdo.Len())
	common.WriteToName(sdo.String(), "JanusSDO.nq")

	// New version abstract call
	// The CSVW stuff should be incorporated into this one
	// var nsdo common.Buffer
	// l = pipes.NewSchemaorg(session, &nsdo)
	// log.Println(l)
	// log.Println(nsdo.Len())
	// common.WriteToName(nsdo.String(), "TESTING.nq")

}
