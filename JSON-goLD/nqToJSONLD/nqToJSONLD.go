package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/kazarena/json-gold/ld"
)

const nquads = `<file:///Users/dfils/src/go/src/lab.esipfed.org/provisium/prov/provenance_response_example.ttl> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/ns/prov#Bundle> .
<file:///Users/dfils/src/go/src/lab.esipfed.org/provisium/prov/provenance_response_example.ttl> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/ns/prov#Entity> .
<file:///Users/dfils/src/go/src/lab.esipfed.org/provisium/prov/provenance_response_example.ttl> <http://www.w3.org/2000/01/rdf-schema#label> "A collection of provenance"^^<http://www.w3.org/2001/XMLSchema#string> .
<file:///Users/dfils/src/go/src/lab.esipfed.org/provisium/prov/provenance_response_example.ttl> <http://www.w3.org/ns/prov#wasAttributedTo> <http://provisium.io#provAQ> .
<file:///Users/dfils/src/go/src/lab.esipfed.org/provisium/prov/provenance_response_example.ttl> <http://www.w3.org/ns/prov#generatedAtTime> "2011-07-16T02:52:02Z"^^<http://www.w3.org/2001/XMLSchema#dateTime> .
<http://provisium.io#usgs> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/ns/prov#Agent> .
<http://provisium.io#usgs> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/ns/prov#Organization> .
<http://provisium.io#usgs> <http://www.w3.org/2000/01/rdf-schema#label> "United States Geological Survey"^^<http://www.w3.org/2001/XMLSchema#string> .
<http://provisium.io#usgs> <http://xmlns.com/foaf/0.1/givenName> "USGS" .
<http://provisium.io#usgs> <http://xmlns.com/foaf/0.1/mbox> <mailto:usgs@example.org> .
<http://provisium.io#dataset> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/ns/prov#Entity> .
<http://provisium.io#dataset> <http://www.w3.org/2000/01/rdf-schema#label> "Some dataset from USGS"^^<http://www.w3.org/2001/XMLSchema#string> .
<http://provisium.io#dataset> <http://www.w3.org/ns/prov#wasGeneratedBy> <http://provisium.io#processingActivity2> .
<http://provisium.io#dataset> <http://www.w3.org/ns/prov#wasDerivedFrom> <http://provisium.io#somewhatProcessedData> .
<http://provisium.io#dataset> <http://www.w3.org/ns/prov#wasAttributedTo> <http://provisium.io#usgs> .
<http://provisium.io#processingActivity2> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/ns/prov#Activity> .
<http://provisium.io#processingActivity2> <http://www.w3.org/2000/01/rdf-schema#label> "A processing activity"^^<http://www.w3.org/2001/XMLSchema#string> .
<http://provisium.io#processingActivity2> <http://www.w3.org/ns/prov#used> <http://provisium.io#somewhatProcessedData> .
<http://provisium.io#processingActivity2> <http://www.w3.org/ns/prov#wasAssociatedWith> <http://provisium.io#usgs> .
<http://provisium.io#somewhatProcessedData> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/ns/prov#Entity> .
<http://provisium.io#somewhatProcessedData> <http://www.w3.org/2000/01/rdf-schema#label> "An intermediate dataset"^^<http://www.w3.org/2001/XMLSchema#string> .
<http://provisium.io#somewhatProcessedData> <http://www.w3.org/ns/prov#wasGeneratedBy> <http://provisium.io#processingActivity1> .
<http://provisium.io#somewhatProcessedData> <http://www.w3.org/ns/prov#wasDerivedFrom> <http://provisium.io#rawSensorMeasurements> .
<http://provisium.io#somewhatProcessedData> <http://www.w3.org/ns/prov#wasAttributedTo> <http://provisium.io#usgs> .
<http://provisium.io#processingActivity1> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/ns/prov#Activity> .
<http://provisium.io#processingActivity1> <http://www.w3.org/2000/01/rdf-schema#label> "First processing activity"^^<http://www.w3.org/2001/XMLSchema#string> .
<http://provisium.io#processingActivity1> <http://www.w3.org/ns/prov#startedAtTime> "2011-07-14T01:01:01Z"^^<http://www.w3.org/2001/XMLSchema#dateTime> .
<http://provisium.io#processingActivity1> <http://www.w3.org/ns/prov#wasAssociatedWith> <http://provisium.io#usgs> .
<http://provisium.io#processingActivity1> <http://www.w3.org/ns/prov#used> <http://provisium.io#rawSensorMeasurements> .
<http://provisium.io#processingActivity1> <http://www.w3.org/ns/prov#used> <http://provisium.io#processingCode> .
<http://provisium.io#processingActivity1> <http://www.w3.org/ns/prov#endedAtTime> "2011-07-14T02:02:02Z"^^<http://www.w3.org/2001/XMLSchema#dateTime> .
<http://provisium.io#rawSensorMeasurements> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/ns/prov#Entity> .
<http://provisium.io#rawSensorMeasurements> <http://www.w3.org/2000/01/rdf-schema#label> "The raw data that came off the sensors"^^<http://www.w3.org/2001/XMLSchema#string> .
<http://provisium.io#rawSensorMeasurements> <http://www.w3.org/ns/prov#wasAttributedTo> <http://provisium.io#usgs> .
<http://provisium.io#processingCode> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/ns/prov#Entity> .
<http://provisium.io#processingCode> <http://www.w3.org/2000/01/rdf-schema#label> "Some processing code"^^<http://www.w3.org/2001/XMLSchema#string> .
<http://provisium.io#processingCode> <http://www.w3.org/ns/prov#wasAttributedTo> <http://provisium.io#usgs> .
<http://provisium.io#pingBacks> <http://www.w3.org/2000/01/rdf-schema#label> "URIs submitted to the pingback service"^^<http://www.w3.org/2001/XMLSchema#string> .
<http://provisium.io#pingBacks> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/ns/prov#Collection> .
<http://provisium.io#pingBacks> <http://www.w3.org/ns/prov#hadMember> <http://provisium.io#Dataset001> .
<http://provisium.io#pingBacks> <http://www.w3.org/ns/prov#hadMember> <http://provisium.io#Dataset002> .
<http://provisium.io#pingBacks> <http://www.w3.org/ns/prov#hadMember> <http://provisium.io#Dataset003> .
<http://provisium.io#pingBacks> <http://www.w3.org/ns/prov#hadMember> <http://provisium.io#Dataset004> .
<http://provisium.io#pingBacks> <http://www.w3.org/ns/prov#hadMember> <http://provisium.io#Dataset005> .
<http://provisium.io#pingBacks> <http://www.w3.org/ns/prov#hadMember> <http://provisium.io#Dataset006> .
<http://provisium.io#pingBacks> <http://www.w3.org/ns/prov#hadMember> <http://provisium.io#Dataset007> .
<http://provisium.io#pingBacks> <http://www.w3.org/ns/prov#hadMember> <http://provisium.io#Dataset008> .
<http://provisium.io#pingBacks> <http://www.w3.org/ns/prov#hadMember> <http://provisium.io#Dataset009> .

`

func main() {
	fmt.Println("Convert NQ to JSONLD")

	ld := rdfToJSONLD(nquads)

	fmt.Print(ld)

}

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
