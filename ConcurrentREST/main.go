// concurrent.go
//
// A simple program to demonstract some basic concurrance (no sync issues)
// for calling an REST API endpoint
//

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

type SearchResults struct {
	Results    string
	GivenName  string
	FamilyName string
}

type OrcidIdentifier struct {
	Value interface{} `json:"value,omitempty"`
	URI   string      `json:"uri,omitempty"`
	Path  string      `json:"path,omitempty"`
	Host  string      `json:"host,omitempty"`
}

func main() {
	start := time.Now()
	ch := make(chan SearchResults) // ch := make(chan string)

	// get the token and input file from the command lines
	tokenPtr := flag.String("token", "xxx-xxx-xxx-xxx", "A valid Orcid token")
	filenamePtr := flag.String("file", "./candidates.csv", "A CSV file of fname, lname, email to feed into the orcid API")
	flag.Parse()

	// read in a CSV file with the URL's
	filenameValue := *filenamePtr // dereference the pointer for my function
	csvdata := readMetaData(filenameValue)

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

func printOrcid(data SearchResults) {
	// get the json string...  indented nicely...   in case we need to inspect it.
	// var out bytes.Buffer
	// err := json.Indent(&out, []byte(data.Results), "", "  ")
	// if err != nil {
	// 	log.Printf("Problem indenting json %v \n", err)
	// 	return
	// }
	// fmt.Println(out.String())

	dec := json.NewDecoder(strings.NewReader(data.Results))
	for {
		t, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Token error %v \n", err)
		}

		if t == "orcid-identifier" {
			var m OrcidIdentifier
			err := dec.Decode(&m)
			if err != nil {
				log.Printf("Decode error %v \n", err)
				return
			}
			fmt.Printf("Results \t%s \t%s\t\t%v  \t%v \t%v\n", data.GivenName, data.FamilyName, m.URI, m.Path, m.Host)
			// At this point I could also send this info off to a function to make up some RDF for this too  (future feature, already have the code for this)
		}
	}
}

func MakeRequest(token, givenname, familyname, emailfrag string, ch chan<- SearchResults) {
	// start := time.Now()
	urlstring := "https://pub.orcid.org/v1.2/search/orcid-bio/?q=query&rows=10&start=0"

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

	results := SearchResults{FamilyName: familyname, GivenName: givenname, Results: string(body)}

	// ch <- fmt.Sprintf("%.2f elapsed with response length: %d %s", secs, len(body), u.String())
	ch <- results //   ch <- string(body)
}

func readMetaData(inputfile string) []Candidates {
	csvFile, err := os.Open(inputfile)
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
