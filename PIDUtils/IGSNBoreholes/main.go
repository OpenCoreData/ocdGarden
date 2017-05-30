package main

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/knakk/sparql"
	"opencoredata.org/ocdGarden/PIDUtils/ExpDOIBuilder/structures"
)

type IGSN struct {
	XMLName                 xml.Name `xml:"sesar:samples"`
	NS0                     string   `xml:"xmlns,attr,omitempty"`
	NS1                     string   `xml:"xmlns:igsn,attr,omitempty"`
	NS2                     string   `xml:"xmlns:sesar,attr,omitempty"`
	NS3                     string   `xml:"xmlns:xsi,attr,omitempty"`
	SchemLocation           string   `xml:"xsi:schemaLocation,attr,omitempty"`
	UserCode                string   `xml:"sesar:sample>sesar:user_code"`
	Sampletype              string   `xml:"sesar:sample>sesar:sample_type"`
	Name                    string   `xml:"sesar:sample>sesar:name"`
	Material                string   `xml:"sesar:sample>sesar:material"`
	Igsn                    string   `xml:"sesar:sample>sesar:igsn,omitempty"`
	Description             string   `xml:"sesar:sample>sesar:description,omitempty"`
	AgeMin                  string   `xml:"sesar:sample>sesar:age_min,omitempty"`
	AgeMax                  string   `xml:"sesar:sample>sesar:age_max,omitempty"`
	AgeUnit                 string   `xml:"sesar:sample>sesar:age_unit,omitempty"`
	Collectionmethod        string   `xml:"sesar:sample>sesar:collection_method,omitempty"`
	Latitude                string   `xml:"sesar:sample>sesar:latitude,omitempty"`
	Longitude               string   `xml:"sesar:sample>sesar:longitude,omitempty"`
	Elevation               string   `xml:"sesar:sample>sesar:elevation,omitempty"`
	PrimaryLocationName     string   `xml:"sesar:sample>sesar:primary_location_name,omitempty"`
	Country                 string   `xml:"sesar:sample>sesar:country,omitempty"`
	Province                string   `xml:"sesar:sample>sesar:province,omitempty"`
	County                  string   `xml:"sesar:sample>sesar:county,omitempty"`
	Collector               string   `xml:"sesar:sample>sesar:collector,omitempty"`
	CollectionStartDate     string   `xml:"sesar:sample>sesar:collection_start_date,omitempty"`
	CollectionDatePrecision string   `xml:"sesar:sample>sesar:collection_date_precision,omitempty"`
	PublishDate             string   `xml:"sesar:sample>sesar:publish_date,omitempty"`
	// Externalurl                 string   `xml:"sesar:sample>sesar:external_urls>sesar:external_url,omitempty"`
	// ExURL                       string   `xml:"sesar:sample>sesar:external_url>sesar:url,omitempty"`
	// ExDescription               string   `xml:"sesar:sample>sesar:external_url>sesar:description,omitempty"`
	// ExURLType                   string   `xml:"sesar:sample>sesar:external_url>sesar:url_type,omitempty"`
	CollectionMethodDescription string `xml:"sesar:sample>sesar:collection_method_descr,omitempty"`
	Size                        string `xml:"sesar:sample>sesar:size,omitempty"`
	CruiseFieldProgram          string `xml:"sesar:sample>sesar:cruise_field_prgrm,omitempty"`
	PlatformType                string `xml:"sesar:sample>sesar:platform_type,omitempty"`
	PlatformName                string `xml:"sesar:sample>sesar:platform_name,omitempty"`
	CollectorDetail             string `xml:"sesar:sample>sesar:collector_detail,omitempty"`
	CurrentArchive              string `xml:"sesar:sample>sesar:current_archive,omitempty"`
	CurrentArchiveContact       string `xml:"sesar:sample>sesar:current_archive_contact,omitempty"`
	OriginalArchive             string `xml:"sesar:sample>sesar:original_archive,omitempty"`
}

type CSDCO struct {
	LocationName           string
	LocationType           string
	Project                string
	LocationID             string
	Site                   string
	Hole                   string
	SiteHole               string
	OriginalID             string
	HoleID                 string
	Platform               string
	Date                   string
	WaterDepthM            string
	Country                string
	State_Province         string
	County_Region          string
	PI                     string
	Lat                    string
	Long                   string
	Elevation              string
	Position               string
	StorageLocationWorking string
	StorageLocationArchive string
	SampleType             string
	Comment                string
	MblfT                  string
	MblfB                  string
	MetadataSource         string
}

const queries = ` 
# Comments are ignored, except those tagging a query.
#tag: csdcoholeid
SELECT ?uri ?project ?lat ?long
WHERE  
{    
  ?uri rdf:type <http://opencoredata.org/id/voc/csdco/v1/CSDCOProject> .    
  ?uri <http://opencoredata.org/id/voc/csdco/v1/holeid> "{{.HoleID}}" . 
  ?uri <http://opencoredata.org/id/voc/csdco/v1/project> ?project . 
  ?uri <http://www.w3.org/2003/01/geo/wgs84_pos#lat> ?lat .
  ?uri <http://www.w3.org/2003/01/geo/wgs84_pos#long> ?long .
}
`

func main() {
	structTest()
	// writeIGSNforCSDCOFeature("CAPA-GBA91-1A")
}

func structTest() {

	var igsn IGSN

	// xmlns:igsn="http://schema.igsn.org/description/1.0"
	// xmlns:sesar="http://app.geosamples.org/description/1.0"
	// ref:  http://www.geosamples.org/profile?igsn=odp000001
	// igsn.NS1 = "http://schema.igsn.org/description/1.0" //igsn
	igsn.NS2 = "http://app.geosamples.org"
	// "http://app.geosamples.org/description/1.0" //sesar
	igsn.NS0 = "http://app.geosamples.org"
	igsn.NS3 = "http://www.w3.org/2001/XMLSchema-instance"
	igsn.SchemLocation = "http://app.geosamples.org/samplev2.xsd"
	igsn.UserCode = "CDO"    // IGSN = ODP000001
	igsn.Sampletype = "Hole" // "http://vocabulary.odm2.org/samplingfeaturetype/borehole/"
	igsn.Name = "CAPA-MBLG91-1A"
	igsn.Material = "Not applicable" // "nil:notApplicable" //http://www.opengis.net/def/nil/OGC/0/missing
	igsn.PublishDate = "2017-05-15"
	igsn.Collectionmethod = "Coring"
	igsn.CollectionMethodDescription = "Coring method varies along hole length"
	// igsn.Size =
	igsn.Latitude = "7.324"
	igsn.Longitude = "13.739"
	igsn.Elevation = "1105"
	igsn.CruiseFieldProgram = "CAPA" // ask about this
	igsn.PlatformType = "unknown"
	// igsn.PlatformName =
	igsn.Collector = "PI"
	igsn.CollectorDetail = "CSDCO PO Box 90250"
	igsn.CurrentArchive = "Continental Scientific Drilling Coordinating Office (CSDCO)"
	igsn.CurrentArchiveContact = "Curator"
	igsn.OriginalArchive = "CSDCO PO Box 90250"

	output, err := xml.MarshalIndent(igsn, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	os.Stdout.Write(output)

}

func writeIGSNforCSDCOFeature(project string) {
	repo, err := sparql.NewRepo("http://opencore.dev/blazegraph/namespace/opencore/sparql")
	if err != nil {
		log.Printf("query make repo: %v\n", err)
	}

	f := bytes.NewBufferString(queries)
	bank := sparql.LoadBank(f)

	q, err := bank.Prepare("csdcoholeid", struct{ HoleID string }{project})
	if err != nil {
		log.Printf("query bank prepair: %v\n", err)
	}

	res, err := repo.Query(q)

	if err != nil {
		log.Printf("query call: %v\n", err)
	}

	data := structures.DataCite{}
	latLongs := []structures.GeoPoint{}
	bindingsTest2 := res.Bindings() // map[string][]rdf.Term

	//  Had to hack this up to deal with
	for key := range bindingsTest2["lat"] {
		latLong := structures.GeoPoint{}
		latLong.Lat = bindingsTest2["lat"][key].String()
		latLong.Long = bindingsTest2["long"][key].String()
		latLongs = append(latLongs, latLong)
	}

	data.CreatorName = "Continental Scientific Drilling Coordination Office"
	data.ExpURI = fmt.Sprintf("http://opencoredata.org/id/expedition/%s", project)
	data.ResourceType = "Field_expedition"
	data.ContributorName = "Open Core Data"
	data.ContributorDOI = "10.17616/R37936"
	data.Title = project
	data.GeoPoint = latLongs
	data.Abstract = "Abstract value here"
	data.Version = "1"
	data.Publisher = "Interdisciplinary Earth Data Applications (IEDA)"
	data.PubYear = "2017"

	// blend with the XML template and return the text

	ht, err := template.New("some template").Parse(IGSNTemplate)
	if err != nil {
		log.Printf("template parse failed: %s", err)
	}

	var buff = bytes.NewBufferString("")
	err = ht.Execute(buff, data)
	if err != nil {
		log.Printf("RDF template execution failed: %s", err)
	}

	// fmt.Println(string(buff.Bytes()))

	writeFile(fmt.Sprintf("./output/csdco_%s.xml", project), string(buff.Bytes()))

}

func writeFile(name string, xmldata string) {
	// Create the output file
	outFile, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	w := bufio.NewWriter(outFile)

	_, err = w.WriteString(xmldata)
	w.Flush()

	if err != nil {
		log.Fatal(err)
	}
}

const IGSNTemplate = `<?xml version="1.0"?>
<igsn:resource 
    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" 
    xmlns:igsn="http://schema.igsn.org/description/1.0" xsi:schemaLocation="http://schema.igsn.org/description/1.0 https://raw.githubusercontent.com/IGSN/metadata/dev/description/resource.xsd" registedObjectType="http://schema.igsn.org/vocab/PhysicalSample">
    <igsn:identifier type="IGSN">http://igsn.org/AU239</igsn:identifier>
    <igsn:title>Sample igsn:AU239</igsn:title>
    <igsn:alternateIdentifiers>
        <igsn:alternateIdentifier>2000362089A</igsn:alternateIdentifier>
    </igsn:alternateIdentifiers>
    <igsn:description/>
    <igsn:registrant>
        <igsn:name>Geoscience Australia</igsn:name>
        <igsn:affiliation>
            <igsn:identifier type="URI">http://pid.geoscience.gov.au/org/ga</igsn:identifier>
            <igsn:name>Geoscience Australia</igsn:name>
        </igsn:affiliation>
    </igsn:registrant>
    <igsn:locations>
        <igsn:geometry type="Point" sridType="4326">SRID=8311;POINTZ(137.8563726 -33.7108293 51)</igsn:geometry>
    </igsn:locations>
    <igsn:resourceTypes>
        <igsn:resourceType>http://vocabulary.odm2.org/specimentype/core/</igsn:resourceType>
    </igsn:resourceTypes>
    <igsn:materials>
        <igsn:material>http://vocabulary.odm2.org/medium/rock/</igsn:material>
    </igsn:materials>
    <igsn:collectionMethods>
        <igsn:collectionMethod>http://www.opengis.net/def/nil/OGC/0/missing</igsn:collectionMethod>
    </igsn:collectionMethods>
    <igsn:isMetadataPublic>Public</igsn:isMetadataPublic>
    <igsn:contributors>
        <igsn:contributor type="HostingInstitution">
            <igsn:identifier type="URI">http://pid.geoscience.gov.au/org/ga</igsn:identifier>
            <igsn:name>Geoscience Australia</igsn:name>
        </igsn:contributor>
    </igsn:contributors>
</igsn:resource>
`
