package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/kazarena/json-gold/ld"
)

func main() {
	fmt.Println("Convert NQ to JSONLD")

	ld := rdfToJSONLD(dgraph3)

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

const dgraph3 = `<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://purl.org/dc/terms/created> "1959-10-08"^^<http://www.w3.org/2001/XMLSchema#date> .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://purl.org/dc/terms/creator> <http://sample.igsn.org/soilarchive/CDS-NSW> .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://purl.org/dc/terms/issued> "2017-01-03"^^<http://www.w3.org/2001/XMLSchema#date> .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://purl.org/dc/terms/spatial> _:b1 .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://purl.org/dc/terms/title> "ANZ soil sample" .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://purl.org/dc/terms/type> <http://pid.geoscience.gov.au/def/voc/igsn-codelists/PhysicalSample> .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://purl.org/dc/terms/type> <http://pid.geoscience.gov.au/def/voc/igsn-codelists/soil> .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://schema.org/additionalType> <http://pid.geoscience.gov.au/def/voc/igsn-codelists/PhysicalSample> .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://schema.org/additionalType> <http://pid.geoscience.gov.au/def/voc/igsn-codelists/soil> .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://schema.org/creator> <http://sample.igsn.org/soilarchive/CDS-NSW> .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://schema.org/dateCreated> "1959-10-08"^^<http://www.w3.org/2001/XMLSchema#date> .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://schema.org/identifier> "soil_specimen_199.CAN.C410.1.2.1" .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://schema.org/title> "ANZ soil sample" .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://schema.org/url> <http://www.anzsoil.org/data/csiro-natsoil/anzsoilml201/soilspecimen/soil_specimen_199.CAN.C410.1.2.1> .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/2002/07/owl#Thing> .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/ns/sosa/Sample> .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://www.w3.org/2000/01/rdf-schema#label> "ANZ soil sample" .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://www.w3.org/ns/dcat#landingPage> <http://www.anzsoil.org/data/csiro-natsoil/anzsoilml201/soilspecimen/soil_specimen_199.CAN.C410.1.2.1> .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://www.w3.org/ns/sosa/isResultOf> _:b0 .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://www.w3.org/ns/sosa/isSampleOf> <http://www.anzsoil.org/data/csiro-natsoil/anzsoilml201/soil/soil_199.CAN.C410> .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://www.w3.org/ns/sosa/isSampleOf> <http://www.anzsoil.org/data/csiro-natsoil/anzsoilml201/soilhorizon/soil_horizon_199.CAN.C410.1.2> .
_:b0 <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/ns/sosa/Sampling> .
_:b0 <http://www.w3.org/ns/sosa/usedProcedure> <http://www.anzsoil.org/def/au/soil/observation-method/soil-pit> .
_:b1 <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://purl.org/dc/terms/Location> .
_:b1 <http://www.w3.org/ns/dcat#centroid> "POINT(146.067917 -34.79847)"^^<http://www.opengis.net/ont/geosparql#asWKT> .`

const dgraph2 = `<http://sample.igsn.org/soilarchive/CDS-NSW> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/ns/prov#Organization> .
<http://sample.igsn.org/soilarchive/CDS-NSW> <http://www.w3.org/2000/01/rdf-schema#label> "CSIRO Division of Soils (NSW)" .
<http://sample.igsn.org/soilarchive/CLW> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/ns/prov#Organization> .
<http://sample.igsn.org/soilarchive/CLW> <http://www.w3.org/2000/01/rdf-schema#label> "CSIRO Land and Water" .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://purl.org/dc/terms/created> "1959-10-08"^^<http://www.w3.org/2001/XMLSchema#date> .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://purl.org/dc/terms/creator> <http://sample.igsn.org/soilarchive/CDS-NSW> .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://purl.org/dc/terms/issued> "2017-01-03"^^<http://www.w3.org/2001/XMLSchema#date> .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://purl.org/dc/terms/spatial> _:b1 .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://purl.org/dc/terms/title> "ANZ soil sample" .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://purl.org/dc/terms/type> <http://pid.geoscience.gov.au/def/voc/igsn-codelists/PhysicalSample> .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://purl.org/dc/terms/type> <http://pid.geoscience.gov.au/def/voc/igsn-codelists/soil> .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://schema.org/additionalType> <http://pid.geoscience.gov.au/def/voc/igsn-codelists/PhysicalSample> .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://schema.org/additionalType> <http://pid.geoscience.gov.au/def/voc/igsn-codelists/soil> .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://schema.org/creator> <http://sample.igsn.org/soilarchive/CDS-NSW> .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://schema.org/dateCreated> "1959-10-08"^^<http://www.w3.org/2001/XMLSchema#date> .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://schema.org/identifier> "soil_specimen_199.CAN.C410.1.2.1" .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://schema.org/title> "ANZ soil sample" .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://schema.org/url> <http://www.anzsoil.org/data/csiro-natsoil/anzsoilml201/soilspecimen/soil_specimen_199.CAN.C410.1.2.1> .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/2002/07/owl#Thing> .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/ns/sosa/Sample> .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://www.w3.org/2000/01/rdf-schema#label> "ANZ soil sample" .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://www.w3.org/ns/dcat#landingPage> <http://www.anzsoil.org/data/csiro-natsoil/anzsoilml201/soilspecimen/soil_specimen_199.CAN.C410.1.2.1> .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://www.w3.org/ns/dcat#qualifiedRelation> _:b4 .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://www.w3.org/ns/prov#qualifiedAttribution> _:b2 .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://www.w3.org/ns/prov#qualifiedAttribution> _:b3 .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://www.w3.org/ns/prov#qualifiedAttribution> _:b5 .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://www.w3.org/ns/sosa/isResultOf> _:b0 .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://www.w3.org/ns/sosa/isSampleOf> <http://www.anzsoil.org/data/csiro-natsoil/anzsoilml201/soil/soil_199.CAN.C410> .
<http://sample.igsn.org/soilarchive/CSAZSXXXXX> <http://www.w3.org/ns/sosa/isSampleOf> <http://www.anzsoil.org/data/csiro-natsoil/anzsoilml201/soilhorizon/soil_horizon_199.CAN.C410.1.2> .
_:b0 <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/ns/sosa/Sampling> .
_:b0 <http://www.w3.org/ns/sosa/usedProcedure> <http://www.anzsoil.org/def/au/soil/observation-method/soil-pit> .
_:b1 <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://purl.org/dc/terms/Location> .
_:b1 <http://www.w3.org/ns/dcat#centroid> "POINT(146.067917 -34.79847)"^^<http://www.opengis.net/ont/geosparql#asWKT> .
_:b2 <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/ns/prov#Attribution> .
_:b2 <http://www.w3.org/ns/dcat#hadRole> <http://registry.it.csiro.au/def/isotc211/CI_RoleCode/owner> .
_:b2 <http://www.w3.org/ns/prov#agent> <http://sample.igsn.org/soilarchive/CLW> .
_:b3 <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/ns/prov#Attribution> .
_:b3 <http://www.w3.org/ns/dcat#hadRole> <http://registry.it.csiro.au/def/isotc211/CI_RoleCode/custodian> .
_:b3 <http://www.w3.org/ns/prov#agent> <http://sample.igsn.org/soilarchive/CLW> .
_:b4 <http://purl.org/dc/terms/relation> <http://www.anzsoil.org/data/csiro-natsoil/anzsoilml201/soilprofile/soil_profile_199.CAN.C410.1> .
_:b4 <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/ns/dcat#Relationship> .
_:b4 <http://www.w3.org/ns/dcat#hadRole> <http://pid.geoscience.gov.au/def/voc/igsn-codelists/hasSamplingFeature> .
_:b5 <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/ns/prov#Attribution> .
_:b5 <http://www.w3.org/ns/dcat#hadRole> <http://registry.it.csiro.au/def/isotc211/CI_RoleCode/originator> .
_:b5 <http://www.w3.org/ns/prov#agent> <http://sample.igsn.org/soilarchive/CDS-NSW> .`

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
