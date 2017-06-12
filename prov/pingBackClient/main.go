package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type urilist struct {
	uri      string
	listitem string
}

func main() {
	fmt.Println("Simple prov-aq demo client")

	body := getAResource("http://127.0.0.1:9900/rdf/graph/void.ttl")
	// do something with body....   save to file, workflow, etc...
	getURIList(body) // go get the links from this...
	// getProvEntry()
	// doPingBack()
}

func getURIList(body []byte) {

	fmt.Println(len(body))

	// parse and print the text/uri-list entry

}

func getAResource(urlstring string) []byte {
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

	var b bytes.Buffer
	err = res.Header.Write(&b) // W3C comments...   multiple LINK items not easy with res.Get.header.
	if err != nil {
		log.Println(err)
	}

	//  Need to split/scan our header to locate all LINKS and look for pingback
	reader := bytes.NewReader(b.Bytes())
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	return body
}
