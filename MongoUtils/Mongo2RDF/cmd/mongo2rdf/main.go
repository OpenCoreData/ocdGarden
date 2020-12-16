package main

import (
	"log"

	// "opencoredata.org/ocdGarden/CSDCO/VaultWalker/pkg/utils"

	"opencoredata.org/ocdGarden/MongoUtils/Mongo2RDF/internal/common"
	"opencoredata.org/ocdGarden/MongoUtils/Mongo2RDF/internal/pipes"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile) // show file and line number
}

func main() {
	// call mongo and lookup the redirection to use...
	session, err := common.GetMongoCon()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Janus Digital Objects
	l := pipes.DigitalObject(session)
	log.Println(l)

	/*
		// Janus Digital Objects
		var sdo common.Buffer
		l := pipes.Schemaorg(session, &sdo)
		log.Println(l)
		log.Println(sdo.Len())
		common.WriteToName(sdo.String(), "JanusSDO.nq")
	*/

	// New Schema.org function
	// var nsdo common.Buffer
	// l = pipes.NewSchemaorg(session, &nsdo) // doesn't exist?
	// log.Println(l)
	// log.Println(nsdo.Len())
	// common.WriteToName(nsdo.String(), "TESTING.nq")

	/*
		// Janus Expeditions
		var jab common.Buffer
		l = pipes.NewFeatureAbsGeoJSON(session, &jab)
		log.Println(l)
		log.Println(jab.Len())
		common.WriteToName(jab.String(), "JanusExp.nq")

		// CSDCO Abstracts
		var cab common.Buffer
		l = pipes.NewAbstracts(session, &cab)
		log.Println(l)
		log.Println(cab.Len())
		common.WriteToName(cab.String(), "CSDCOAbstracts.nq")
	*/

	/*

	   root@opencore:~# ./bin/mc cat local/csdco-do-meta/bj6s49vtr9b9lop9n40g
	   [
	    {
	     "@graph": [
	      {
	       "@id": "_:bbj6s49vtr9b9lop9n40g",
	       "@type": [
	        "http://schema.org/PropertyValue"
	       ],
	       "http://schema.org/propertyID": [
	        {
	         "@value": "SHA256"
	        }
	       ],
	       "http://schema.org/value": [
	        {
	         "@value": "e88c1fa53004f75e2f1a761284c42e22d69e1750e70904c02c0b726ff42250ad"
	        }
	       ]
	      },
	      {
	       "@id": "http://opencoredata.org/id/do/bj6s49vtr9b9lop9n40g",
	       "@type": [
	        "http://www.schema.org/DigitalDocument"
	       ],
	       "http://schema.org/additionType": [
	        {
	         "@id": "http://opencoredata.org/voc/csdco/v1/ICDFiles"
	        }
	       ],
	       "http://schema.org/dateCreated": [
	        {
	         "@value": "2009-11-02"
	        }
	       ],
	       "http://schema.org/description": [
	        {
	         "@value": "Digital object of type ICD named YNP3-ABD09-1A-1L-1.pdf for CSDCO project YNP3"
	        }
	       ],
	       "http://schema.org/encodingFormat": [
	        {
	         "@value": "application/pdf"
	        }
	       ],
	       "http://schema.org/identifier": [
	        {
	         "@id": "_:bbj6s49vtr9b9lop9n40g"
	        }
	       ],
	       "http://schema.org/isRelatedTo": [
	        {
	         "@value": "YNP3"
	        }
	       ],
	       "http://schema.org/license": [
	        {
	         "@id": "http://example.com/cc0.html"
	        }
	       ],
	       "http://schema.org/name": [
	        {
	         "@value": "YNP3-ABD09-1A-1L-1.pdf"
	        }
	       ],
	       "http://schema.org/url": [
	        {
	         "@id": "http://opencoredata.org/id/do/e88c1fa53004f75e2f1a761284c42e22d69e1750e70904c02c0b726ff42250ad"
	        }
	       ]
	      }
	     ],
	     "@id": "http://opencoredata.org/objectgraph/id/e88c1fa53004f75e2f1a761284c42e22d69e1750e70904c02c0b726ff42250ad"
	    }
	   ]root@opencore:~#./bin/mc ls local


	*/

}
