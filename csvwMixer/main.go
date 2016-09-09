package main

import (
	// "encoding/json"
	//"github.com/jeffail/gabs"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	// "strconv"
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"text/template"
)

// schema.org Dataset metadata structs
type SchemaOrgMetadata struct {
	Context             []interface{} `json:"@context"`
	Type                string        `json:"@type"`
	Author              Author        `json:"author"`
	Description         string        `json:"description"`
	Distribution        Distribution  `json:"distribution"`
	GlviewDataset       string        `json:"glview:dataset"`
	GlviewKeywords      string        `json:"glview:keywords"`
	OpenCoreLeg         string        `json:"opencore:leg"`
	OpenCoreSite        string        `json:"opencore:site"`
	OpenCoreHole        string        `json:"opencore:hole"`
	OpenCoreMeasurement string        `json:"opencore:measurement"`
	Keywords            string        `json:"keywords"`
	Name                string        `json:"name"`
	Spatial             Spatial       `json:"spatial"`
	URL                 string        `json:"url"`
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

// W3c csvw metadata structs
// 	context := `["http://www.w3.org/ns/csvw", {"@language": "en"}]`
type CSVWMeta struct {
	Context      CSVContext   `json:"@context"`
	Dc_license   Dc_license   `json:"dc:license"`
	Dc_modified  Dc_modified  `json:"dc:modified"`
	Dc_publisher Dc_publisher `json:"dc:publisher"`
	Dc_title     string       `json:"dc:title"`
	Dcat_keyword []string     `json:"dcat:keyword"`
	TableSchema  TableSchema  `json:"tableSchema"`
	URL          string       `json:"url"`
}

type CSVContext struct {
	Vocab    string `json:"@vocab"`
	Language string `json:"@language"`
}

type Dc_license struct {
	Id string `json:"@id"`
}

type Dc_modified struct {
	Type  string `json:"@type"`
	Value string `json:"@value"`
}

type Dc_publisher struct {
	Schema_name string     `json:"schema:name"`
	Schema_url  Schema_url `json:"schema:url"`
}

type Schema_url struct {
	Id string `json:"@id"`
}

type TableSchema struct {
	AboutURL   string    `json:"aboutUrl"`
	Columns    []Columns `json:"columns"`
	PrimaryKey string    `json:"primaryKey"`
}

type Columns struct {
	Datatype       string   `json:"datatype"`
	Dc_description string   `json:"dc:description"`
	Name           string   `json:"name"`
	Required       bool     `json:"required"`
	Titles         []string `json:"titles"`
}

type TemplateData struct {
	URL string
	Data []map[int]string
	ColInfo map[int]string
	ColType map[int]string
}

func main() {
	// get the mongo connection
	mgoconn, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer mgoconn.Close()
	c := mgoconn.DB("test").C("csvwmeta")

	// Get the CSVW metadata datasets (could do this with a service call to ocdServices now too)
	URI := "http://opencoredata.org/id/dataset/fb78f3bd-ce92-4bb3-b72a-0de66e50152a"
	result := CSVWMeta{}
	err = c.Find(bson.M{"url": URI}).One(&result)
	if err != nil {
		log.Printf("URL lookup error: %v", err)
	}

	// range over the metadata data info I need as a prelude to placing in a map
	ms := make(map[string]string)
	colinfo := make(map[int]string)
	coltype := make(map[int]string)
	var a []string
	// datatype will be one of string, int or float (decimal)
	for index, item := range result.TableSchema.Columns {
		log.Printf(" %s  %s  %s  %v  %s", item.Datatype, item.Dc_description, item.Name, item.Required, item.Titles)
		ms[item.Name] = item.Datatype
		colinfo[index] = item.Name
		coltype[index] = item.Datatype
		a = append(a, item.Name)
	}

	log.Printf("%v\n", a)

	csvdata := getFile(result.URL, mgoconn)
	reader := csv.NewReader(bytes.NewBuffer(csvdata))
	reader.FieldsPerRecord = -1
	reader.Comma = '\t' // Use tab-delimited instead of comma <---- here!
	rawCSVData, err := reader.ReadAll()
	if err != nil {
		log.Printf("Error here %s\n", err)
	}

	for _, each := range rawCSVData {
		fmt.Printf("%v\n", each)
	}

	var dataAray []map[int]string
	for _, each := range rawCSVData {
		data := make(map[int]string)
		for i, _ := range a {
			data[i] = each[i]
		}
		dataAray = append(dataAray, data)
	}

	for _, row := range dataAray {
		fmt.Printf("%v\n", row)
	}

	qrytmp, err := ioutil.ReadFile("jsontemplate.txt")
	check(err)

	var buff = bytes.NewBufferString("")
	t, err := template.New(" func template").Parse(string(qrytmp))
	if err != nil {
		log.Printf(" func template creation failed: %s", err)
	}
	
	tdata := TemplateData{URL: "http://opencoredata/id/ID", Data: dataAray, ColInfo: colinfo, ColType: coltype}
	
	err = t.Execute(buff, tdata)
	check(err)
	qry := string(buff.Bytes())

	log.Print(qry)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getFile(URI string, session *mgo.Session) []byte {

	c := session.DB("test").C("schemaorg")
	result := SchemaOrgMetadata{} // need this struct  (it's everywhere.   what can I do about that?   move only ocdServices from ocdWeb?)
	err := c.Find(bson.M{"url": URI}).One(&result)
	if err != nil {
		log.Printf("URL lookup error: %v", err)
	}

	filename := result.Name // the file name
	//  can I just 303 now to the download?  Perhaps I shouldn't in case some client don't follow
	mongodb := session.DB("test")
	file, _ := mongodb.GridFS("fs").Open(filename)
	buf := make([]byte, file.Size())
	file2, err := file.Read(buf)
	if err != nil {
		log.Printf("Error calling aggregation_janusURLSet : %v  length %d", err, file2)
	}

	return buf
}
