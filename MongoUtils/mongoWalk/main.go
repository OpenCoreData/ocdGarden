package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"encoding/json"
	"fmt"
	"os"
	"text/template"

	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	sparql "github.com/knakk/sparql"
)

const queries = `
# Comments are ignored, except those tagging a query.

# tag: maxage
prefix rdfs: <http://www.w3.org/2000/01/rdf-schema#> 
prefix iodp: <http://data.oceandrilling.org/core/1/> 
prefix foaf: <http://xmlns.com/foaf/0.1/> 
prefix owl: <http://www.w3.org/2002/07/owl#> 
prefix xsd: <http://www.w3.org/2001/XMLSchema#> 
prefix rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#> 
SELECT (MAX(?age) AS ?maxage) (MIN(?age) AS ?minage)
WHERE {
   ?s iodp:leg "192" .
   ?s iodp:ma ?age   .
}


# tag: query2
PREFIX chronos: <http://www.chronos.org/loc-schema#>
PREFIX geo: <http://www.w3.org/2003/01/geo/wgs84_pos#>
SELECT  ?dataset ?ob  ?age ?depth ?long ?lat
FROM <http://chronos.org/janusamp#>
WHERE {
  ?ob chronos:age ?age . 
  ?ob chronos:depth ?depth . 
  ?ob <http://purl.org/linked-data/cube#dataSet>  ?dataset .
  ?dataset geo:long ?long .
  ?dataset geo:lat ?lat .
   ?dataset chronos:leg "{{.Leg}}" .  
   ?dataset chronos:site "{{.Site}}" .
   ?dataset chronos:hole "{{.Hole}}" . 
}
ORDER BY ?dataset DESC(?age)
LIMIT 1

# tag: csiro
prefix gts: <http://resource.geosciml.org/ontology/timescale/gts#> 
prefix thors: <http://resource.geosciml.org/ontology/timescale/thors#> 
prefix tm: <http://def.seegrid.csiro.au/isotc211/iso19108/2002/temporal>
PREFIX isc: <http://resource.geosciml.org/classifier/ics/ischart/>
PREFIX rdfs: <http://www.w3.org/2000/01/rdf-schema#>
SELECT *
WHERE {
                 ?era gts:rank ?rank .
                 ?era thors:begin/tm:temporalPosition/tm:value ?begin .
                 ?era thors:begin/tm:temporalPosition/tm:frame <http://resource.geosciml.org/classifier/cgi/geologicage/ma> .
                 ?era thors:end/tm:temporalPosition/tm:value ?end .
                 ?era thors:end/tm:temporalPosition/tm:frame <http://resource.geosciml.org/classifier/cgi/geologicage/ma> .
                 ?era rdfs:label ?name .
                 BIND ( "439."^^xsd:decimal AS ?targetAge )
                 FILTER ( ?targetAge > xsd:decimal(?end) )
                 FILTER ( ?targetAge < xsd:decimal(?begin) )
}

`

type CSIROStruct struct {
	Head struct {
		Vars []string `json:"vars"`
	} `json:"head"`
	Results struct {
		Bindings []struct {
			Begin struct {
				Datatype string `json:"datatype"`
				Type     string `json:"type"`
				Value    string `json:"value"`
			} `json:"begin"`
			End struct {
				Datatype string `json:"datatype"`
				Type     string `json:"type"`
				Value    string `json:"value"`
			} `json:"end"`
			Era struct {
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"era"`
			Name struct {
				Type     string `json:"type"`
				Value    string `json:"value"`
				XML_lang string `json:"xml:lang"`
			} `json:"name"`
			Rank struct {
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"rank"`
			TargetAge struct {
				Datatype string `json:"datatype"`
				Type     string `json:"type"`
				Value    string `json:"value"`
			} `json:"targetAge"`
		} `json:"bindings"`
	} `json:"results"`
}

type Expedition struct {
	Expedition string `json:"expedition"`
	Site       string `json:"site"`
	Hole       string `json:"hole"`
}

func main() {
	// call mongo and lookup the redirection to use...
	session, err := GetMongoCon()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	IndexCSVW(session)
	// IndexSchema(session)

}

func IndexCSVW(session *mgo.Session) {
	// Optional. Switch the session to a monotonic behavior.
	csvw := session.DB("expedire").C("features")

	// Find all Documents CSVW
	var csvwDocs []Expedition
	err := csvw.Find(nil).All(&csvwDocs)
	if err != nil {
		fmt.Printf("this is error %v \n", err)
	}

	// index some data
	for _, item := range csvwDocs {
		// fmt.Printf("Indexed item %d with Leg: %s\n", index, item.Expedition)

		res := SPARQLCall(item.Expedition, item.Site, item.Hole, "query2", "http://opencoredata.org/sparql")

		solutionsTest := res.Solutions() // map[string][]rdf.Term
		// fmt.Println("res.Solutions():")
		for _, i := range solutionsTest {
			fmt.Printf("Leg Site Hole Age: %s %s %s %s \n", item.Expedition, item.Site, item.Hole, i["age"])

			CSIROHack(fmt.Sprint(i["age"]))
			// csiroSolutions := csirores.Solutions()
			// fmt.Println(csirores)
			// for _, x := range csiroSolutions {
			// 	fmt.Printf("TEST item %s \n", x)
			// }

		}

	}

	// // Update Example for Mongo
	// 	colQuerier := bson.M{"name": "Ale"}
	// 	change := bson.M{"$set": bson.M{"phone": "+86 99 8888 7777", "timestamp": time.Now()}}
	// 	err = c.Update(colQuerier, change)
	// 	if err != nil {
	// 		panic(err)
	// 	}

}

func SPARQLCall(leg string, site string, hole string, query string, endpoint string) *sparql.Results {
	repo, err := sparql.NewRepo(endpoint,
		sparql.Timeout(time.Millisecond*15000),
	)
	if err != nil {
		log.Printf("%s\n", err)
	}

	f := bytes.NewBufferString(queries)
	bank := sparql.LoadBank(f)

	q, err := bank.Prepare(query, struct{ Leg, Site, Hole string }{leg, site, hole})
	if err != nil {
		log.Printf("%s\n", err)
	}

	res, err := repo.Query(q)
	if err != nil {
		log.Printf("%s\n", err)
	}

	return res
}

func CSIROHack(age string) {

	const url = "http://resource.geosciml.org/sparql/isc2014?query=PREFIX+isc%3A+%3Chttp%3A%2F%2Fresource.geosciml.org%2Fclassifier%2Fics%2Fischart%2F%3E%0D%0APREFIX+gts%3A+%3Chttp%3A%2F%2Fresource.geosciml.org%2Fontology%2Ftimescale%2Fgts%23%3E%0D%0APREFIX+thors%3A+%3Chttp%3A%2F%2Fresource.geosciml.org%2Fontology%2Ftimescale%2Fthors%23%3E%0D%0APREFIX+tm%3A+%3Chttp%3A%2F%2Fdef.seegrid.csiro.au%2Fisotc211%2Fiso19108%2F2002%2Ftemporal%23%3E%0D%0APREFIX+rdfs%3A+%3Chttp%3A%2F%2Fwww.w3.org%2F2000%2F01%2Frdf-schema%23%3E%0D%0A%0D%0ASELECT+%2A%0D%0AWHERE+%7B%0D%0A++++++++++++++++%3Fera+gts%3Arank+%3Frank+.%0D%0A+++++++++++++++++%3Fera+thors%3Abegin%2Ftm%3AtemporalPosition%2Ftm%3Avalue+%3Fbegin+.%0D%0A+++++++++++++++++%3Fera+thors%3Abegin%2Ftm%3AtemporalPosition%2Ftm%3Aframe+%3Chttp%3A%2F%2Fresource.geosciml.org%2Fclassifier%2Fcgi%2Fgeologicage%2Fma%3E+.%0D%0A+++++++++++++++++%3Fera+thors%3Aend%2Ftm%3AtemporalPosition%2Ftm%3Avalue+%3Fend+.%0D%0A+++++++++++++++++%3Fera+thors%3Aend%2Ftm%3AtemporalPosition%2Ftm%3Aframe+%3Chttp%3A%2F%2Fresource.geosciml.org%2Fclassifier%2Fcgi%2Fgeologicage%2Fma%3E+.%0D%0A++++++++++++++++%3Fera+rdfs%3Alabel+%3Fname+.%0D%0A+++++++++++++++++BIND+%28+%22{{.}}%22%5E%5Exsd%3Adecimal+AS+%3FtargetAge+%29%0D%0A+++++++++++++++++FILTER+%28+%3FtargetAge+%3E+xsd%3Adecimal%28%3Fend%29+%29%0D%0A+++++++++++++++++FILTER+%28+%3FtargetAge+%3C+xsd%3Adecimal%28%3Fbegin%29+%29%0D%0A%7D%0D%0A"

	dt, err := template.New("RDF template").Parse(url)
	if err != nil {
		log.Printf("RDF template creation failed for hole data: %s", err)
	}

	var buff = bytes.NewBufferString("")
	err = dt.Execute(buff, age)
	if err != nil {
		log.Printf("RDF template execution failed: %s", err)
	}

	req, _ := http.NewRequest("GET", string(buff.Bytes()), nil)

	req.Header.Add("accept", "application/sparql-results+json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	csiroStruct := CSIROStruct{}
	json.Unmarshal(body, &csiroStruct)

	// loop on csiroStruct.Results.Bindings
	for _, item := range csiroStruct.Results.Bindings {
		fmt.Println(item.Era.Type)
		fmt.Println(item.Era.Value)
		fmt.Println(item.Name.Type)
		fmt.Println(item.Name.Value)
	}
	// fmt.Println(string(body))

}

func CSIROCall() *sparql.Results {
	repo, err := sparql.NewRepo("http://resource.geosciml.org/sparql/isc2014",
		sparql.Timeout(time.Millisecond*15000),
	)
	if err != nil {
		log.Printf("%s\n", err)
	}

	f := bytes.NewBufferString(queries)
	bank := sparql.LoadBank(f)

	q, err := bank.Prepare("csiro", struct{}{})
	if err != nil {
		log.Printf("%s\n", err)
	}

	res, err := repo.Query(q)
	if err != nil {
		log.Printf("%s\n", err)
	}

	return res
}

func GetMongoCon() (*mgo.Session, error) {
	host := os.Getenv("MONGO_HOST")
	return mgo.Dial(host)
}
