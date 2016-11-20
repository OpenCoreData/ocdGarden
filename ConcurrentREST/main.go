// concurrent.go
//
// A simple program to demonstract some basic concurrance (no sync issues)
// for calling an REST API endpoint
//
//  Some references worth looking over
// http://blog.narenarya.in/concurrent-http-in-go.html
// http://stackoverflow.com/questions/33104192/how-to-run-10000-goroutines-in-parallel-where-each-routine-calls-an-api
// https://medium.com/golangspec/synchronized-goroutines-part-i-4fbcdd64a4ec#.qav8o43uj
// https://medium.com/golangspec/synchronized-goroutines-part-ii-b1130c815c9d#.mximrcqzv
// http://nomad.so/2016/01/interesting-ways-of-using-go-channels/

package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type Candidates struct {
	GivenName  string
	FamilyName string
	EmailFrag  string
}

type OrcidIdentifier struct {
	Value interface{} `json:"value,omitempty"`
	URI   string      `json:"uri,omitempty"`
	Path  string      `json:"path,omitempty"`
	Host  string      `json:"host,omitempty"`
}

func main() {
	start := time.Now()
	ch := make(chan string)

	// read in a CSV file with the URL's
	csvdata := readMetaData()

	// get the token from the command lines
	tokenPtr := flag.String("token", "xxx-xxx-xxx-xxx", "A valid Orcid token")
	flag.Parse()

	// send of the calls based on the CSV file
	for _, person := range csvdata {
		go MakeRequest(*tokenPtr, strings.TrimSpace(person.GivenName), strings.TrimSpace(person.FamilyName), strings.TrimSpace(person.EmailFrag), ch)
	}

	// read the channels
	for range csvdata {
		// fmt.Println(<-ch)
		printOrcid(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func printOrcid(data string) {
	dec := json.NewDecoder(strings.NewReader(data))
	for {
		t, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error %v \n", err)
		}

		if t == "orcid-identifier" {
			// fmt.Println("found ID")
			// fmt.Println(t)

			//for dec.More() {
			var m OrcidIdentifier
			// decode an array value (Message)
			err := dec.Decode(&m)
			if err != nil {
				log.Printf("Error %v \n", err)
				return
			}
			fmt.Printf("Values %v  %v  %v: %v\n", m.Value, m.URI, m.Path, m.Host)
			//}
		}
	}
}

func MakeRequest(token, givenname, familyname, emailfrag string, ch chan<- string) {
	// start := time.Now()
	urlstring := "https://pub.orcid.org/v1.2/search/orcid-bio/?q=family-name%3AFils%20AND%20given-names%3ADoug*%20OR%20email%3A*%40iodp.org&rows=10&start=0"

	u, err := url.Parse(urlstring)
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	q.Set("q", fmt.Sprintf("family-name:%s+AND+given-names:%s*+OR+email:*@%s", familyname, givenname, emailfrag))
	// u.RawQuery = q.Encode()  // this should work, but Orcid doesn't like the way the URL is being encoded
	u.RawQuery = fmt.Sprintf("q=family-name:%s+AND+given-names:%s*+OR+email:*@%s&rows=10&start=0", familyname, givenname, emailfrag)

	req, _ := http.NewRequest("GET", u.String(), nil)

	// req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")                     // oddly the content-type is ignored for the accept header...
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token)) // make a var to hide key
	req.Header.Set("Cache-Control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	// secs := time.Since(start).Seconds()
	body, _ := ioutil.ReadAll(res.Body)

	// ch <- fmt.Sprintf("%.2f elapsed with response length: %d %s", secs, len(body), u.String())
	ch <- string(body)
}

func readMetaData() []Candidates {
	csvFile, err := os.Open("./candidates.csv")
	defer csvFile.Close()
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(csvFile)
	//r.Comma = '\t' // Use tab-delimited instead of comma

	r.FieldsPerRecord = -1
	lines, err := r.ReadAll()
	if err != nil {
		log.Fatalf("Error reading all lines: %v", err)
	}

	commentLines := 0
	callstoMake := make([]Candidates, len(lines)-commentLines)

	for i, entry := range lines {
		ob := Candidates{GivenName: entry[0], FamilyName: entry[1], EmailFrag: entry[2]} //  later can split this for more complex CSV lines with many items
		callstoMake[i] = ob
	}

	return callstoMake

}
