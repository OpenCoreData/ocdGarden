package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// ProvPingBackURLs holds the URLs related to prov-aq pingback and has prov resoruces
type ProvPingBackURLs struct {
	HasProv  string
	PingBack string
}

func main() {
	fmt.Println("Simple prov-aq demo client")

	body, provLinks := getAResource("http://opencoredata.org/rdf/graph/void.ttl")

	fmt.Println(len(body)) // body of the resoruces we started with..   do whatever with it...

	fmt.Println(provLinks)
	// existingProv := getProvEntry(provLinks.HasProv) //  see what prov the resource has
	// existingProv := getProvEntry("http://127.0.0.1:9900/id/rdf/graph/void.ttl/provenance") //  see what prov the resource has
	existingProv := getProvEntry("http://opencoredata.org/id/graph/void.ttl/provenance") //  see what prov the resource has
	fmt.Println(string(existingProv))                                                    // body of existing prov for resources..  do whatever with it...

	// status := doPingBack(provLinks.PingBack)
	status := doPingBack("http://openoredata.org/rdf/graph/void.ttl/pingback")
	fmt.Println(status)
}

// doPingBack will send via POST prov package and expect (hope for) a 204 in response
func doPingBack(urlstring string) string {
	var jsonStr = []byte(`{"title":"Example POST body here as a JSON package"}`)
	req, err := http.NewRequest("POST", urlstring, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	// body, _ := ioutil.ReadAll(resp.Body)  // with 204 we wont even get a body
	// fmt.Println("response Body:", string(body))

	return resp.Status
}

func getProvEntry(urlstring string) []byte {
	u, err := url.Parse(urlstring)
	if err != nil {
		log.Println(err)
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("Accept", "application/json") // oddly the content-type is ignored for the accept header...
	req.Header.Set("Cache-Control", "no-cache")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}

	return body
}

func getAResource(urlstring string) ([]byte, ProvPingBackURLs) {
	u, err := url.Parse(urlstring)
	if err != nil {
		log.Println(err)
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("Accept", "application/json") // oddly the content-type is ignored for the accept header...
	req.Header.Set("Cache-Control", "no-cache")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	var provLinks ProvPingBackURLs

	// Read the headers  (there has got to be a better way)
	links := res.Header["Link"]
	for link := range links {
		linkset := strings.Split(links[link], ";")
		if strings.TrimSpace(linkset[1]) == "rel=\"http://www.w3.org/ns/prov#pingbck\"" {
			fmt.Printf("The pingback is: %s\n", removeBraces(linkset[0]))
			provLinks.PingBack = removeBraces(linkset[0])
		}
		if strings.TrimSpace(linkset[1]) == "rel=\"http://www.w3.org/ns/prov#has_provenance\"" {
			fmt.Printf("The provenance is: %s\n", removeBraces(linkset[0]))
			provLinks.HasProv = removeBraces(linkset[0])
		}
	}

	// Read the body..  but I really don't care or use it in this test...
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}

	return body, provLinks
}

func removeBraces(url string) string {
	url = strings.TrimSpace(url)
	url = strings.Replace(url, "<", "", -1)
	url = strings.Replace(url, ">", "", -1)
	return url
}
