package facilities

import (
	"bytes"
	"fmt"
	"log"
	"text/template"

	"opencoredata.org/ocdGarden/ExpDOIBuilder/structures"

	sparql "github.com/knakk/sparql"
)

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

#tag: csdcoproj
SELECT ?uri ?lat ?long
WHERE  
{    
  ?uri rdf:type <http://opencoredata.org/id/voc/csdco/v1/CSDCOProject> .    
  ?uri <http://opencoredata.org/id/voc/csdco/v1/project> "{{.Project}}" . 
  ?uri <http://www.w3.org/2003/01/geo/wgs84_pos#lat> ?lat .
  ?uri <http://www.w3.org/2003/01/geo/wgs84_pos#long> ?long .
}

#tag: allcsdcoproj
SELECT DISTINCT ?project
WHERE  
{    
  ?uri rdf:type <http://opencoredata.org/id/voc/csdco/v1/CSDCOProject> .    
  ?uri <http://opencoredata.org/id/voc/csdco/v1/project> ?project . 
  ?uri <http://www.w3.org/2003/01/geo/wgs84_pos#lat> ?lat .
  ?uri <http://www.w3.org/2003/01/geo/wgs84_pos#long> ?long .
}
`

func CSDCOProj() {
	repo, err := sparql.NewRepo("http://opencore.dev/blazegraph/namespace/opencore/sparql")
	if err != nil {
		log.Printf("query make repo: %v\n", err)
	}

	f := bytes.NewBufferString(queries)
	bank := sparql.LoadBank(f)

	q, err := bank.Prepare("allcsdcoproj")
	// q, err := bank.Prepare("csdcoproj")
	if err != nil {
		log.Printf("query bank prepair: %v\n", err)
	}

	res, err := repo.Query(q)
	bindingsTest2 := res.Bindings() // map[string][]rdf.Term

	for _, proj := range bindingsTest2["project"] {
		fmt.Println(proj)
		CSDCOProjMetadata(proj.String())
	}

}

func CSDCOProjMetadata(project string) {
	repo, err := sparql.NewRepo("http://opencore.dev/blazegraph/namespace/opencore/sparql")
	if err != nil {
		log.Printf("query make repo: %v\n", err)
	}

	f := bytes.NewBufferString(queries)
	bank := sparql.LoadBank(f)

	q, err := bank.Prepare("csdcoproj", struct{ Project string }{project})
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

	ht, err := template.New("some template").Parse(structures.CSDCOtemplate)
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
