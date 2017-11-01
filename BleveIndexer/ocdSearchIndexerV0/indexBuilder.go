package main

import (
	// "fmt"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/blevesearch/bleve"
	"github.com/fatih/camelcase"
	"github.com/knakk/sparql"
	"gopkg.in/mgo.v2"
)

type Mdoc struct {
	ProfileID      string   `json:"profile_id"`
	GroupID        string   `json:"group_id"`
	LastModified   string   `json:"last_modified"`
	Tags           []string `json:"tags"`
	Read           bool     `json:"read"`
	Starred        bool     `json:"starred"`
	Authored       bool     `json:"authored"`
	Confirmed      bool     `json:"confirmed"`
	Hidden         bool     `json:"hidden"`
	CitationKey    string   `json:"citation_key"`
	SourceType     string   `json:"source_type"`
	Language       string   `json:"language"`
	ShortTitle     string   `json:"short_title"`
	ReprintEdition string   `json:"reprint_edition"`
	Genre          string   `json:"genre"`
	Country        string   `json:"country"`
	Translators    []struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	} `json:"translators"`
	SeriesEditor            string `json:"series_editor"`
	Code                    string `json:"code"`
	Medium                  string `json:"medium"`
	UserContext             string `json:"user_context"`
	PatentOwner             string `json:"patent_owner"`
	PatentApplicationNumber string `json:"patent_application_number"`
	PatentLegalStatus       string `json:"patent_legal_status"`
	Notes                   string `json:"notes"`
	Accessed                string `json:"accessed"`
	FileAttached            bool   `json:"file_attached"`
	Created                 string `json:"created"`
	ID                      string `json:"id"`
	Year                    int    `json:"year"`
	Month                   int    `json:"month"`
	Day                     int    `json:"day"`
	Source                  string `json:"source"`
	Edition                 string `json:"edition"`
	Authors                 []struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	} `json:"authors"`
	Keywords     []string `json:"keywords"`
	Pages        string   `json:"pages"`
	Volume       string   `json:"volume"`
	Issue        string   `json:"issue"`
	Websites     []string `json:"websites"`
	Publisher    string   `json:"publisher"`
	City         string   `json:"city"`
	Institution  string   `json:"institution"`
	Department   string   `json:"department"`
	Series       string   `json:"series"`
	SeriesNumber string   `json:"series_number"`
	Chapter      string   `json:"chapter"`
	Editors      []struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	} `json:"editors"`
	Title    string `json:"title"`
	Revision string `json:"revision"`
	// Identifiers string `json:"identifiers"`
	Identifiers []struct {
		Doi  string `json:"dio"`
		Issn string `json:"issn"`
	} `json:"identifiers"`
	Abstract  string `json:"abstract"`
	Type      string `json:"type"`
	OCDSOURCE string `json:ocdsource`
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
	OCDSOURCE              string `json:"ocdsource"`
}

type SchemaOrgMetadata struct {
	Context             Context      `json:"@context"`
	Type                string       `json:"@type"`
	Author              Author       `json:"author"`
	Description         string       `json:"description"`
	Distribution        Distribution `json:"distribution"`
	GlviewDataset       string       `json:"glview:dataset"`
	GlviewKeywords      string       `json:"glview:keywords"`
	GlviewMD5           string       `json:"glview:md5"`
	OpenCoreLeg         string       `json:"opencore:leg"`
	OpenCoreSite        string       `json:"opencore:site"`
	OpenCoreHole        string       `json:"opencore:hole"`
	OpenCoreProgram     string       `json:"opencore:program"`
	OpenCoreMeasurement string       `json:"opencore:measurement"`
	Keywords            string       `json:"keywords"`
	Name                string       `json:"name"`
	Spatial             Spatial      `json:"spatial"`
	URL                 string       `json:"url"`
	OCDSOURCE           string       `json:"ocdsource"`
	OpenCoreParamters   []Params     `json:"opencore:params"`
}

type Context struct {
	Vocab    string `json:"@vocab"`
	GeoLink  string `json:"glview"`
	OpenCore string `json:"opencore"`
}

type Author struct {
	Type        string `json:"@type"`
	Description string `json:"description"`
	Name        string `json:"name"`
	URL         string `json:"url"`
}

type Distribution struct {
	Type           string `json:"@type"`
	ContentURL     string `json:"contentUrl"`
	DatePublished  string `json:"datePublished"`
	EncodingFormat string `json:"encodingFormat"`
	InLanguage     string `json:"inLanguage"`
}

type Spatial struct {
	Type string `json:"@type"`
	Geo  Geo    `json:"geo"`
}

type Geo struct {
	Type      string `json:"@type"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type Params struct {
	Pname        string
	Pdatatype    string
	Pdescription string
}

const queries = `
# Comments are ignored, except those tagging a query.

# tag: call1
SELECT  ?uri ?name ?type ?column ?desc WHERE {  
?uri <http://example.org/rdf/type> <http://opencoredata.org/id/voc/janus/v1/JanusQuerySet> .   
?uri     <http://opencoredata.org/id/voc/janus/v1/struct_name> "{{.}}" .  
?uri   <http://opencoredata.org/id/voc/janus/v1/go_struct_name> ?name .
?uri  <http://opencoredata.org/id/voc/janus/v1/go_struct_type> ?type .  
?uri    <http://opencoredata.org/id/voc/janus/v1/column_id> ?column  .
?uri    <http://opencoredata.org/id/voc/janus/v1/JanusMeasurement> ?jmes .  
?jmes  <http://opencoredata.org/id/voc/janus/v1/json_descript>  ?desc  
}
ORDER By (xsd:integer(?column))
`

func main() {
	indexSchemaOrg()
	// indexAbstracts()
	// indexCSDCO()
}

// JRSO
func indexCSDCO() {
	// open a new index
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New("csdco.bleve", mapping)

	// Open mongo and read out a record...   then index it.
	session, err := GetMongoCon()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Do schema.org
	d := session.DB("test").C("csdco")

	var results2 []CSDCO
	err = d.Find(nil).All(&results2)
	if err != nil {
		log.Printf("Error calling test schema.org collection : %v", err)
	}

	for _, elem2 := range results2 {
		elem2.OCDSOURCE = "CSDCO"
		err = index.Index(elem2.HoleID, elem2) // TODO:  review if this is really what I want for a UID here?
	}

}

// JRSO
func indexSchemaOrg() {
	// open a new index
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New("jrso.bleve", mapping)

	// Open mongo and read out a record...   then index it.
	session, err := GetMongoCon()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Do schema.org
	d := session.DB("test").C("schemaorg")

	var results2 []SchemaOrgMetadata
	err = d.Find(nil).All(&results2)
	if err != nil {
		log.Printf("Error calling test schema.org collection : %v", err)
	}

	for _, elem2 := range results2 {
		elem2.OCDSOURCE = "JRSO"
		fmt.Printf("Looking for measurement: %s\n", elem2.OpenCoreMeasurement)
		fmt.Println(sparqlCall(elem2.OpenCoreMeasurement))
		elem2.OpenCoreParamters = sparqlCall(elem2.OpenCoreMeasurement)
		splitted := camelcase.Split(elem2.OpenCoreMeasurement)     // hackish split of terms
		elem2.OpenCoreMeasurement = strings.Join(splitted[:], " ") // array to string
		err = index.Index(elem2.URL, elem2)
	}

}

// function to perform SPARQL call to get parameters and descriptions to load into the index
// get the name, type and description
func sparqlCall(measurement string) []Params {
	repo, err := sparql.NewRepo("http://opencoredata.org/blazegraph/namespace/opencore/sparql")

	if err != nil {
		log.Printf("query make repo: %v\n", err)
	}

	f := bytes.NewBufferString(queries)
	bank := sparql.LoadBank(f)

	q, err := bank.Prepare("call1", strings.ToLower(measurement))
	if err != nil {
		log.Printf("query bank prepair: %v\n", err)
	}

	// fmt.Println(q)

	res, err := repo.Query(q)

	if err != nil {
		log.Printf("query call: %v\n", err)
	}

	var fs []Params

	solutionsTest := res.Solutions() // map[string][]rdf.Term
	for _, i := range solutionsTest {
		// fmt.Printf("At postion %v with %v \n\n", k, i)
		data := Params{}
		data.Pname = i["name"].String()
		data.Pdatatype = i["type"].String()
		data.Pdescription = i["desc"].String()
		fs = append(fs, data)
	}
	return fs
}

func indexAbstracts() {
	// open a new index
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New("abstracts.bleve", mapping)

	// Open mongo and read out a record...   then index it.
	session, err := GetMongoCon()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	//Do the abstracts
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("abstracts").C("csdco")

	var results []Mdoc
	err = c.Find(nil).All(&results)
	if err != nil {
		log.Printf("Error calling CSDCO abstract collection : %v", err)
	}

	for _, elem := range results {
		elem.OCDSOURCE = "CSDCO"
		err = index.Index(elem.ID, elem)
	}
}

func GetMongoCon() (*mgo.Session, error) {
	host := os.Getenv("MONGO_HOST")

	return mgo.Dial(host)
}
