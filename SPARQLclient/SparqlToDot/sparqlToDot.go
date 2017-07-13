package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/awalterschulze/gographviz"
	"github.com/knakk/sparql"
)

// sparql "opencoredata.org/ocdCommons/sparqlclient"

const queries = `
# Comments are ignored, except those tagging a query.

# tag: cdfviztest
PREFIX schemaorg: <http://schema.org/>
SELECT DISTINCT ?repository  ?name ?memberOfName ?endpoint_url ?endpoint_description ?endpoint_method
WHERE {
  ?repository rdf:type <http://schema.org/Organization>   .
  ?repository schemaorg:name ?name   .
  ?repository schemaorg:memberOf ?mo  .
  ?mo   schemaorg:programName ?memberOfName   .
    ?repository schemaorg:potentialAction [ schemaorg:target ?action ] .
    ?action schemaorg:urlTemplate ?endpoint_url .
    ?action schemaorg:description ?endpoint_description .
    ?action schemaorg:httpMethod ?endpoint_method .

}
ORDER BY ?name
`

type FacilityServices struct {
	Repository          string
	Name                string
	MemberOfName        string
	EndpointURL         string
	EndpointDescription string
	EndpointMethod      string
}

func main() {
	fs := call()
	serviceDot(fs)
	// fmt.Println(fs)
}

// Take a struct and make a graphvix string in it
func serviceDot(fs []FacilityServices) {
	graph := gographviz.NewGraph()
	if err := graph.SetName("G"); err != nil {
		panic(err)
	}
	if err := graph.SetDir(true); err != nil {
		panic(err)
	}

	// Loop on results to build graphviz graph
	// TODO  this FAILS here since the string are NOT being escaped.  I need to look into
	// reducing them to a valid string and using a map to align them to a label in a map
	// though.. use MD5 hashes to ID the strings..  then be sure to assign lables to ALL nodes
	for _, v := range fs {
		// n := map[string]string{"label": "http://this.is and here I am"}
		graph.AddNode("G", md5hash(v.MemberOfName), map[string]string{"label": v.MemberOfName})
		graph.AddNode("G", md5hash(v.Name), map[string]string{"label": v.Name})
		graph.AddNode("G", md5hash(v.EndpointURL), map[string]string{"label": v.EndpointURL, "url": v.EndpointURL})
		graph.AddEdge(md5hash(v.MemberOfName), md5hash(v.Name), true, nil)
		graph.AddEdge(md5hash(v.Name), md5hash(v.EndpointURL), true, nil)
	}

	output := graph.String()
	output = strings.Replace(output, "->", " -> ", -1)
	fmt.Println(output)
}

func md5hash(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func stripChars(str, chr string) string {
	return strings.Map(func(r rune) rune {
		if strings.IndexRune(chr, r) < 0 {
			return r
		}
		return -1
	}, str)
}

func call() []FacilityServices {
	repo, err := sparql.NewRepo("http://localhost:9999/blazegraph/namespace/ecrwg/sparql")
	if err != nil {
		log.Printf("query make repo: %v\n", err)
	}

	f := bytes.NewBufferString(queries)
	bank := sparql.LoadBank(f)

	// q, err := bank.Prepare("CSDCOHoleID", struct{ HOLEID string }{"http://opencoredata/id/resource/csdco/project/aafblp-llb06-2a"})
	q, err := bank.Prepare("cdfviztest")
	if err != nil {
		log.Printf("query bank prepair: %v\n", err)
	}

	res, err := repo.Query(q)

	if err != nil {
		log.Printf("query call: %v\n", err)
	}

	var fs []FacilityServices

	solutionsTest := res.Solutions() // map[string][]rdf.Term
	// fmt.Println("res.Solutions():")
	for _, i := range solutionsTest {
		// fmt.Printf("At postion %v with %v \n\n", k, i)
		data := FacilityServices{}
		data.Repository = i["repository"].String()
		data.Name = i["name"].String()
		data.MemberOfName = i["memberOfName"].String()
		data.EndpointURL = i["endpoint_url"].String()
		data.EndpointDescription = i["endpoint_description"].String()
		data.EndpointMethod = i["endpoint_method"].String()
		fs = append(fs, data)
	}

	return fs
}
