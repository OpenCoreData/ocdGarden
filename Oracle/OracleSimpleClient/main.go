package main

// Simple call to Oracle in a self contained packge
// be sure to do:
// export JANUS_USERNAME=X
// export JANUS_PASSWORD=X
// export JANUS_HOST=X
// export JANUS_SERVICENAME=X
//
// also needs Oracle client SDK

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/jmoiron/sqlx"
	"github.com/kisielk/sqlstruct"
	"gopkg.in/rana/ora.v3"
)

type JanusPaleoSamplecVSW struct {
	Tables []JanusPaleoSampletable `json:"tables"`
}

type JanusPaleoSampletable struct {
	URL string                     `json:"url"`
	Row []JanusPaleoSamplejanusRow `json:"row"`
}

type JanusPaleoSamplejanusRow struct {
	URL       string             `json:"url"`
	Rownum    int                `json:"rownum"`
	Describes []JanusPaleoSample `json:"describes"`
}

// make name generic  How to load the body of struct
type JanusPaleoSample struct {
	Leg                  int64          `json:"Leg"`
	Site                 int64          `json:"Site"`
	Hole                 string         `json:"Hole"`
	Core                 int64          `json:"Core"`
	Core_type            string         `json:"Core_type"`
	Section_number       int64          `json:"Section_number"`
	Section_type         string         `json:"Section_type"`
	Top_cm               float64        `json:"Top_cm"`
	Bot_cm               float64        `json:"Bot_cm"`
	Depth_mbsf           float64        `json:"Depth_mbsf"`
	Sample_id            int64          `json:"Sample_id"`
	Location             string         `json:"Location"`
	Fossil_group         int64          `json:"Fossil_group"`
	Fossil_group_name    string         `json:"Fossil_group_name"`
	Geologic_age_old     sql.NullString `json:"Geologic_age_old"`
	Geologic_age_young   sql.NullString `json:"Geologic_age_young"`
	Zone_old             sql.NullString `json:"Zone_old"`
	Zone_young           sql.NullString `json:"Zone_young"`
	Group_abundance_name sql.NullString `json:"Group_abundance_name"`
	Preservation_name    sql.NullString `json:"Preservation_name"`
	Scientist_id         int64          `json:"Scientist_id"`
	Scientist_last_name  string         `json:"Scientist_last_name"`
	Scientist_first_name sql.NullString `json:"Scientist_first_name"`
	Paleo_sample_comment sql.NullString `json:"Paleo_sample_comment"`
}

// Not needed?
func JanusPaleoSampleModel() *JanusPaleoSample {
	return &JanusPaleoSample{}
}

// init function to register oracle driver for SQL driver
func init() {
	ora.Register(nil)
}

func check(e error) {
	if e != nil {
		log.Print(e)
	}
}

// Populate the file prefix
const namespace_prefix = `@prefix dc: <http://purl.org/dc/elements/1.1/> .
@prefix rdfs: <http://www.w3.org/2000/01/rdf-schema#> .
@prefix rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#> .
@prefix geo: <http://www.w3.org/2003/01/geo/wgs84_pos#> .
@prefix datacite: <http://purl.org/spar/datacite/> .
@prefix glperson: <http://schema.geolink.org/dev/view#Person> .
@prefix foaf: <http://xmlns.com/foaf/0.1/> .
@prefix vcard: <http://www.w3.org/2006/vcard/ns#> .
@prefix rdfa: <http://www.w3.org/ns/rdfa#> .
@prefix xsd: <http://www.w3.org/2001/XMLSchema#> .
@prefix chronos: <http://chronos.org/voc/1/> .
@prefix mlgs: <http://data.oceandrilling.org/core/MLGS/> .
@prefix iodp: <http://data.oceandrilling.org/core/iodp/> .
@prefix glview: <http://schema.geolink.org/dev/view#> .
@prefix ocd: <http://opencoredata.org/voc/1/> .
@prefix ocdjanus: <http://opencoredata.org/voc/janus/1/> .

`

const sampleTemplate = `
<http://opencoredata.org/id/expedition/{{.Leg}}> a glview:PhysicalSample ;   
      rdfs:label "Sample resource for {{.Leg}}"@en ;
      ocdjanus:expedition "{{.Leg}}" ; 
    {{if .Leg}} ocdjanus:leg  "{{.Leg}}"  ; {{end}}
	{{if .Site}} ocdjanus:site  "{{.Site}}"  ; {{end}}
	{{if .Hole}} ocdjanus:hole  "{{.Hole}}"  ; {{end}}
	{{if .Core}} ocdjanus:core  "{{.Core}}"  ; {{end}}
	{{if .Core_type}} ocdjanus:core_type  "{{.Core_type}}"  ; {{end}}
	{{if .Section_number}} ocdjanus:section_number  "{{.Section_number}}"  ; {{end}}
	{{if .Section_type}} ocdjanus:section_type  "{{.Section_type}}"  ; {{end}}
	{{if .Top_cm}} ocdjanus:top_cm  "{{.Top_cm}}"  ; {{end}}
	{{if .Bot_cm}} ocdjanus:bot_cm  "{{.Bot_cm}}"  ; {{end}}
	{{if .Depth_mbsf}} ocdjanus:depth_mbsf  "{{.Depth_mbsf}}"  ; {{end}}
	{{if .Sample_id}} ocdjanus:sample_id  "{{.Sample_id}}"  ; {{end}}
	{{if .Location}} ocdjanus:location  "{{.Location}}"  ; {{end}}
	{{if .Fossil_group}} ocdjanus:fossil_group  "{{.Fossil_group}}"  ; {{end}}
	{{if .Fossil_group_name}} ocdjanus:fossil_group_name  "{{.Fossil_group_name}}"  ; {{end}}
	{{if .Geologic_age_old}} ocdjanus:geologic_age_old  "{{.Geologic_age_old.String}}"  ; {{end}}
	{{if .Geologic_age_young}} ocdjanus:geologic_age_young  "{{.Geologic_age_young.String}}"  ; {{end}}
	{{if .Zone_old}} ocdjanus:zone_old  "{{.Zone_old.String}}"  ; {{end}}
	{{if .Zone_young}} ocdjanus:zone_young  "{{.Zone_young.String}}"  ; {{end}}
	{{if .Group_abundance_name}} ocdjanus:group_abundance_name  "{{.Group_abundance_name.String}}"  ; {{end}}
	{{if .Preservation_name}} ocdjanus:preservation_name  "{{.Preservation_name.String}}"  ; {{end}}
	{{if .Scientist_id}} ocdjanus:scientist_id  "{{.Scientist_id}}"  ; {{end}}
	{{if .Scientist_last_name}} ocdjanus:scientist_last_name  "{{.Scientist_last_name}}"  ; {{end}}
	{{if .Scientist_first_name}} ocdjanus:scientist_first_name  "{{.Scientist_first_name.String}}"  ; {{end}}
	{{if .Paleo_sample_comment}} ocdjanus:paleo_sample_comment  "{{.Paleo_sample_comment.String}}"  ; {{end}}
      ocdjanus:leg "{{.Leg}}" .  

`

// main simple Oracle call test
func main() {
	err := JanusPaleoSampleFunc()
	if err != nil {
		log.Printf(`Error: "%s"`, err)
	}
}

// GetJanusCon returns a Oracle DB connection
func GetJanusCon() (*sqlx.DB, error) {
	username := os.Getenv("JANUS_USERNAME")
	password := os.Getenv("JANUS_PASSWORD")
	host := os.Getenv("JANUS_HOST")
	servicename := os.Getenv("JANUS_SERVICENAME")
	connectionString := fmt.Sprintf("%s/%s@(DESCRIPTION=(ADDRESS_LIST=(ADDRESS=(PROTOCOL=TCP)(HOST=%s)(PORT=1521)))(CONNECT_DATA=(SERVICE_NAME=%s)))", username, password, host, servicename)
	return sqlx.Open("ora", connectionString)
}

// JanusPaleoSampleFunc calls to database and issues the select from ocd_paleo_sample call
func JanusPaleoSampleFunc() error {

	// the query
	// qry := "SELECT * FROM ocd_paleo_sample WHERE rownum < 101"
	qry := "SELECT * FROM ocd_paleo_sample"

	// get the Oracle connection
	conn, err := GetJanusCon()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Make the calls the get the rows
	rows, err := conn.Query(qry)
	if err != nil {
		log.Printf(`Error with "%s": %s`, qry, err)
	}

	// struct to hold all the results
	// allResults := []JanusPaleoSamplejanusRow{}
	i := 1 // rows is not a enumeration..  sigh..  have to do it myself..

	filepath := fmt.Sprintf("./output/%s", "sampleData.rdf")
	f, err := os.Create(filepath)
	check(err)

	_, err = f.Write([]byte(namespace_prefix))
	check(err)

	defer f.Close()

	// go through the rows
	for rows.Next() {
		d := []JanusPaleoSample{}
		var t JanusPaleoSample
		err := sqlstruct.Scan(&t, rows)
		if err != nil {
			log.Print(err)
		}
		d = append(d, t)

		t.Paleo_sample_comment.String = strings.Replace(t.Paleo_sample_comment.String, "\"", " ", -1)

		// rowURL := fmt.Sprintf("%s/%s#row=%v", "URL_pattern/", "filename", i)
		// aRow := JanusPaleoSamplejanusRow{rowURL, i, d}
		// allResults = append(allResults, aRow)
		i = i + 1
		//log.Printf("row: %v\n\n", aRow) // UGLY print statement in here...   I will try and clean this up and make a more formated output
		//log.Printf("%s\n\n", ToRDF(t))
		_, err = f.Write([]byte(ToRDF(t)))
		check(err)
	}

	f.Close() // really no need with the defer
	return nil
}

// ToRDF convert struct to RDF
func ToRDF(row JanusPaleoSample) string {

	ct, err := template.New("RDF template").Parse(sampleTemplate)
	if err != nil {
		log.Printf("RDF template creation failed for sample: %s", err)
	}

	var buff = bytes.NewBufferString("")
	err = ct.Execute(buff, row)
	if err != nil {
		log.Printf("RDF template execution failed: %s", err)
	}

	return string(buff.Bytes())
}
