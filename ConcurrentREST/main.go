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
//

package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Calls struct {
	URL string
}

func main() {
	start := time.Now()
	ch := make(chan string)

	// read in a CSV file with the URL's
	csvdata := readMetaData()

	// for _, url := range os.Args[1:] {
	for _, url := range csvdata {
		go MakeRequest(strings.TrimSpace(url.URL), ch)
	}

	// for range os.Args[1:] {
	for range csvdata {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func readMetaData() []Calls {
	csvFile, err := os.Open("./sites.csv")
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
	callstoMake := make([]Calls, len(lines)-commentLines)

	for i, entry := range lines {
		ob := Calls{URL: entry[0]} //  later can split this for more complex CSV lines with many items
		callstoMake[i] = ob
	}

	return callstoMake

}

func MakeRequest(url string, ch chan<- string) {
	start := time.Now()
	resp, _ := http.Get(url)

	secs := time.Since(start).Seconds()
	body, _ := ioutil.ReadAll(resp.Body)
	ch <- fmt.Sprintf("%.2f elapsed with response length: %d %s", secs, len(body), url)
}
