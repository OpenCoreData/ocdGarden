package facilities

import (
	"bytes"
	"fmt"
	"html/template"
	"log"

	"opencoredata.org/ocdGarden/ExpDOIBuilder/structures"

	sparql "github.com/knakk/sparql"
)

const jrsoqueries = ` 
# Comments are ignored, except those tagging a query.

#tag: jrsoallproj
SELECT *
WHERE  
{    
  	?uri rdf:type  <http://schema.geolink.org/1.0/base/main#Cruise> 
}

#tag: jrsoproj
prefix skos: <http://www.w3.org/2004/02/skos/core#>
SELECT DISTINCT ?lat ?long ?program ?leg
WHERE  
{    
  <{{.Project}}> skos:narrower ?narrow  .
  ?narrow <http://www.w3.org/2003/01/geo/wgs84_pos#lat> ?lat .
  ?narrow <http://www.w3.org/2003/01/geo/wgs84_pos#long> ?long .
  ?narrow <http://opencoredata.org/id/voc/janus/v1/program> ?program .     
  ?narrow   <http://opencoredata.org/id/voc/janus/v1/leg> ?leg .
}

`

func JRSOProj() {
	repo, err := sparql.NewRepo("http://opencore.dev/blazegraph/namespace/opencore/sparql")
	if err != nil {
		log.Printf("NIL at make repo: %v\n", err)
	}

	f := bytes.NewBufferString(jrsoqueries)
	bank := sparql.LoadBank(f)

	q, err := bank.Prepare("jrsoallproj")
	if err != nil {
		log.Printf("query bank prepair: %v\n", err)
	}

	res, err := repo.Query(q)
	// fmt.Print(res)
	bindingsTest2 := res.Bindings() // map[string][]rdf.Term

	for _, uri := range bindingsTest2["uri"] {
		fmt.Println(uri)
		JRSOProjMetadata(uri.String())
	}
}

func JRSOProjMetadata(project string) {
	repo, err := sparql.NewRepo("http://opencore.dev/blazegraph/namespace/opencore/sparql")
	if err != nil {
		log.Printf("query make repo: %v\n", err)
	}

	f := bytes.NewBufferString(jrsoqueries)
	bank := sparql.LoadBank(f)

	q, err := bank.Prepare("jrsoproj", struct{ Project string }{project})
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

	data.CreatorName = "JOIDES Resulution Science Operator"
	data.ExpURI = fmt.Sprintf("http://opencoredata.org/id/expedition/%s", bindingsTest2["leg"][0].String())
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

	ht, err := template.New("some template").Parse(structures.JRSOTemplate)
	if err != nil {
		log.Printf("template parse failed: %s", err)
	}

	var buff = bytes.NewBufferString("")
	err = ht.Execute(buff, data)
	if err != nil {
		log.Printf("RDF template execution failed: %s", err)
	}

	if len(bindingsTest2) == 0 {
		fmt.Printf("No expedition data found for %s\n", project)
	} else {
		// fmt.Println(string(buff.Bytes()))
		writeFile(fmt.Sprintf("./output/jrso_%s.xml", bindingsTest2["leg"][0].String()), string(buff.Bytes()))
	}

}
