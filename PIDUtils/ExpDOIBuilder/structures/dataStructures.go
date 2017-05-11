package structures

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

// bring in the DataCite style struct to test serlizing to struct the SPARQL results
type DataCite struct {
	ExpDOI          string   // Is this the ID of the expedition or something else
	ExpURI          string   // something like http://data.rvdata.us/id/cruise/TN272 for R2R
	ResourceType    string   // Field_expedition
	CreatorName     string   // Open Core Data
	ContributorDOI  string   // re3data DOI  static   10.17616/R37936
	Title           string   // Expedition XXX on Joides Resoultion or CSDCO
	Abstract        string   // * abstract here...
	DateCollected   string   // ** Really a data of a specific format 2011-11-05/2011-12-17
	ContributorName string   // Joides Resolution Science Office || Continental Scientific Drilling Corrdinating Office
	RelatedDOIs     []string // 1 or more related DOI's
	GeoPoint        []GeoPoint
	Publisher       string // Rolling Deck to Repository (R2R) Program
	Version         string // 1, 2, 3, etc
	PubYear         string // 2016
}

type GeoPoint struct {
	Long string // longitude
	Lat  string // latitude
}

const DataCitev4Template = `<?xml version="1.0"?>
<resource xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://datacite.org/schema/kernel-4" xsi:schemaLocation="http://datacite.org/schema/kernel-4 http://schema.datacite.org/meta/kernel-4/metadata.xsd">
   <identifier identifierType="DOI">{{.ExpDOI}}</identifier>
  <alternateIdentifiers>
     <alternateIdentifier alternateIdentifierType="URL">{{.ExpURI}}</alternateIdentifier>
  </alternateIdentifiers>
  <resourceType resourceTypeGeneral="Event">{{.ResourceType}}</resourceType>
  <creators>
    <creator>
      <creatorName>{{.CreatorName}}</creatorName>
    </creator>
  </creators>
  <titles>
    <title>{{.Title}}</title>
  </titles>
  <descriptions>
    <description descriptionType="Abstract">{{.Abstract}}</description>
  </descriptions>
  <dates>
    <date dateType="Collected">{{.DateCollected}}</date>
  </dates>
  <language>en</language>
  <contributors>
    <contributor contributorType="Producer">
      <contributorName>{{.ContributorName}}</contributorName>
      <nameIdentifier nameIdentifierScheme="DOI">{{.ContributorDOI}}</nameIdentifier>
    </contributor>
    <contributor contributorType="Sponsor">
      <contributorName>National Science Foundation</contributorName>
      <nameIdentifier nameIdentifierScheme="DOI">10.13039/100000001</nameIdentifier>
    </contributor>
  </contributors>
  <relatedIdentifiers>
                {{range $ITEMS := .RelatedDOIs}}
    <relatedIdentifier relatedIdentifierType="DOI" relationType="IsReferencedBy">{{.}}</relatedIdentifier>  
                {{end}}
  </relatedIdentifiers>
  <geoLocations>
	  {{range $latlong := .GeoPoint}} <geoLocation> <geoLocationPoint>
        <pointLongitude> {{.Long}}</pointLongitude>
        <pointLatitude>{{.Lat}}</pointLatitude>
    </geoLocationPoint> </geoLocation>
			{{end}}
  </geoLocations>
  <publisher>{{.Publisher}}</publisher>
  <version>{{.Version}}</version>
  <publicationYear>{{.PubYear}}</publicationYear>
</resource>
`
