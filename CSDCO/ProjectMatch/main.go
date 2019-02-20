package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"text/template"

	sparql "github.com/knakk/sparql"
)

const queries = `
# Comments are ignored, except those tagging a query.

# tag: test1
SELECT DISTINCT *
WHERE 
{ 
  ?uri rdf:type <http://opencoredata.org/id/voc/csdco/v1/CSDCOProject> . 
  ?uri <http://opencoredata.org/id/voc/csdco/v1/project> "{{.}}" . 
  ?uri ?p ?o . 
}

# tag: urionly
SELECT DISTINCT ?uri
WHERE 
{ 
  ?uri rdf:type <http://opencoredata.org/id/voc/csdco/v1/CSDCOProject> . 
  ?uri <http://opencoredata.org/id/voc/csdco/v1/project> "{{.}}" . 
  ?uri ?p ?o . 
}

`

// The XML I sadly have to marshal
// <?xml version='1.0' encoding='UTF-8'?>
// <sparql xmlns='http://www.w3.org/2005/sparql-results#'>
//         <head>
//                 <variable name='uri'/>
//         </head>
//         <results>
//                 <result>
//                         <binding name='uri'>
//                                 <uri>http://opencoredata/id/resource/csdco/project/aafblp</uri>
//                         </binding>
//                 </result>
//         </results>
// </sparql>

type SparqlResult struct {
	XMLName xml.Name `xml:"sparql"`
	Results Results  `xml:"results"`
}

type Results struct {
	Result Result `xml:"result"`
}

type Result struct {
	Binding Binding `xml:"binding"`
}

type Binding struct {
	Uri string `xml:"uri"`
}

func main() {
	content, err := ioutil.ReadFile("testSet.txt") // testSet.txt  or projectFolderList.txt
	if err != nil {
		fmt.Printf("Error with ioutils %s\n", err)
	}
	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		splits := strings.Split(line, " ")
		ss := strings.TrimSpace(splits[0])
		res := blazeCall(ss)
		fmt.Printf("Project %s with URL %s \n", ss, res)
	}

	// for _, line := range lines {
	// 	splits := strings.Split(line, " ")
	// 	res := callHack(strings.TrimSpace(splits[0]))

	// 	v := SparqlResult{}

	// 	// fmt.Println(res)

	// 	err := xml.Unmarshal([]byte(res), &v)
	// 	if err != nil {
	// 		fmt.Printf("error: %v", err)
	// 		return
	// 	}

	// 	fmt.Printf("Matching %s to URI: %#v\n", strings.TrimSpace(splits[0]), v.Results.Result.Binding.Uri)
	// }

}

// try blaze again with SPARQL library now that it's fixed to address the JSON issue
func blazeCall(project string) string {
	repo, err := sparql.NewRepo("http://localhost:9999/blazegraph/namespace/csdco/sparql")
	// repo, err := sparql.NewRepo("http://opencoredata.org/sparql")

	if err != nil {
		log.Printf("query make repo: %v\n", err)
	}

	f := bytes.NewBufferString(queries)
	bank := sparql.LoadBank(f)

	q, err := bank.Prepare("urionly", project)
	if err != nil {
		log.Printf("query bank prepair: %v\n", err)
	}

	res, err := repo.Query(q)

	if err != nil {
		log.Printf("query call: %v\n", err)
	}

	bindingsTest := res.Results.Bindings // map[string][]rdf.Term
	var URI string
	for _, i := range bindingsTest {
		URI = fmt.Sprintf("%v", i["uri"].Value)
	}
	return URI
}

// crappy hack call when I was having issues getting blazegraph to return JSON
func callHack(project string) string {

	// Example SPARQL call used
	// SELECT DISTINCT *
	// WHERE
	// {
	//   ?uri rdf:type <http://opencoredata.org/id/voc/csdco/v1/CSDCOProject> .
	//   ?uri <http://opencoredata.org/id/voc/csdco/v1/project> "AAFBLP" .
	//   ?uri ?p ?o .
	// }

	const url = "http://localhost:9999/bigdata/namespace/csdco/sparql?query=SELECT%20DISTINCT%20%3Furi%20WHERE%20%20%7B%20%20%20%20%3Furi%20rdf%3Atype%20%3Chttp%3A%2F%2Fopencoredata.org%2Fid%2Fvoc%2Fcsdco%2Fv1%2FCSDCOProject%3E%20.%20%20%20%20%3Furi%20%3Chttp%3A%2F%2Fopencoredata.org%2Fid%2Fvoc%2Fcsdco%2Fv1%2Fproject%3E%20%22{{.}}%22%20.%20%20%7D"

	dt, err := template.New("RDF template").Parse(url)
	if err != nil {
		log.Printf("RDF template creation failed for hole data: %s", err)
	}

	var buff = bytes.NewBufferString("")
	err = dt.Execute(buff, project)
	if err != nil {
		log.Printf("RDF template execution failed: %s", err)
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", string(buff.Bytes()), nil)

	//  Not working, so I have to do crappy XML decoding
	// req.Header.Add("Accept", "application/sparql-results+json")
	// req.Header.Add("cache-control", "no-cache")
	// req.Header.Add("accept-encoding", "gzip, deflate")

	// fmt.Println(req.Header)

	res, _ := client.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)

}
