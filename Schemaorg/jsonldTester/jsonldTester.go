package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/kazarena/json-gold/ld"
)

// Notes
// 1) This checks all the JSON-LD on a page, which may not be what you want. So if there are multiple JSON-LD scripts they are all checked
// 2) This does not validate the use of context in the JSON-LD, only that it is well formed.

func main() {
	fmt.Println("Simple testing for JSON-LD in a web page")

	resPtr := flag.String("res", "", "A URL to check the JSON-LD content of")
	filePtr := flag.String("file", "", "A sitemap file of URLs to check the JSON-LD content of")
	urlPtr := flag.String("url", "", "A URL to a sitemap of URLs to check the JSON-LD content of")
	// maxNbConcurrentGoroutines := flag.Int("maxNbConcurrentGoroutines", 5, "the number of goroutines that are allowed to run concurrently")
	flag.Parse()

	// Dummy channel to coordinate the number of concurrent goroutines.
	// concurrentGoroutines := make(chan struct{}, *maxNbConcurrentGoroutines)
	// for i := 0; i < *maxNbConcurrentGoroutines; i++ {
	// 	concurrentGoroutines <- struct{}{}
	// }
	// done := make(chan bool)           // The done channel indicates when a single goroutine has finished
	// waitForAllJobs := make(chan bool) // The waitForAllJobs channel allows the main program to wait until we have indeed done all the jobs.

	var urls []string

	if *resPtr != "" {
		urls = append(urls, *resPtr)
		fmt.Println(urls)
	}

	if *filePtr != "" {
		urls = siteMapFile(*filePtr)
	}

	if *urlPtr != "" {
		urls = siteMapURL(*urlPtr)
	}

	start := time.Now()

	// Collect all the jobs, and since the job is finished, we can
	// release another spot for a goroutine.
	// nbJobs := len(urls)
	// go func() {
	// 	for i := 0; i < nbJobs; i++ {
	// 		<-done
	// 		// Say that another goroutine can now start.
	// 		concurrentGoroutines <- struct{}{}
	// 	}
	// 	// We have collected all the jobs, the program
	// 	// can now terminate
	// 	waitForAllJobs <- true
	// }()

	// for i := 1; i <= nbJobs; i++ {
	// 	fmt.Printf("ID: %v: waiting to launch!\n", i)
	// 	// Try to receive from the concurrentGoroutines channel. When we have something,
	// 	// it means we can start a new goroutine because another one finished.
	// 	// Otherwise, it will block the execution until an execution
	// 	// spot is available.
	// 	<-concurrentGoroutines
	// 	fmt.Printf("ID: %v: it's my turn!\n", i)
	// 	go func(id int) {
	// 		// DoWork()
	// 		url := urls[id-1]
	// 		err := parseTest(url)
	// 		if err != nil {
	// 			fmt.Printf("ERROR:  URL %s  at index %d has error: %v\n", url, i, err)
	// 		} else {
	// 			fmt.Printf("URL processed %d/%d: %s \n", i, nbJobs, url)
	// 		}
	// 		fmt.Printf("ID: %v: all done!\n", id)

	// 		done <- true
	// 	}(i)
	// }

	// Old loop
	for index := range urls {
		url := urls[index]
		err := parseTest(url)
		if err != nil {
			fmt.Printf("URL %s  at index %d has error: %v\n", url, index, err)
		} else {
			fmt.Printf("URL processed %d/%d: %s \n", index+1, len(urls), url)
		}
	}

	elapsed := time.Since(start)
	log.Printf("P418 indexer took %s", elapsed)
}

func siteMapURL(sourceurl string) []string {
	urls := []string{}

	// Get the data
	resp, err := http.Get(sourceurl)
	if err != nil {
		fmt.Printf("Error fetching sitemap at URL: %v\n", err)
		log.Fatal(err) // go ahead and make errors here fatal...  we have to have this file
	}
	if resp.StatusCode != 200 {
		b, _ := ioutil.ReadAll(resp.Body)
		log.Fatal(string(b))
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		urls = append(urls, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return urls
}

func siteMapFile(filename string) []string {
	urls := []string{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		urls = append(urls, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return urls
}

func parseTest(url string) error {
	// http2 thoughts
	// Need to make this up higher and pass along the client to the functions, otherwise
	// it's pointless to make N number http 2.0 connections...   the goal is to have ONE
	// Look at connection pooling as a easier way to approach concurency in this code?
	// transport := &http2.Transport()
	// client := &http.Client{
	// 	Transport: t,
	// }

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Print(err) // not even being able to make a req instance..  might be a fatal thing?
	}

	req.Header.Set("User-Agent", "EarthCube_DataBot/1.0")

	resp, err := client.Do(req)
	if err != nil {
		log.Print(err) // TODO..  make better.. but recall these errors should NOT be fatal
	}
	if resp.StatusCode != 200 {
		b, _ := ioutil.ReadAll(resp.Body)
		log.Print(err)
		log.Print(string(b))
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		fmt.Printf("error %v", err)
	}

	// Version that just looks for script type application/ld+json
	// this will look for ALL nodes in the doc that match, there may be more than one
	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		// s.Has()
		val, _ := s.Attr("type")
		if val == "application/ld+json" {
			// fmt.Printf("%s\n", s.Text()) //  or send off to a scheme.org parser (JSONLD parser)
			err = isValid(s.Text())
		}
	})

	return err
}

func isValid(jsonld string) error {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	options.Format = "application/nquads"

	var myInterface interface{}
	err := json.Unmarshal([]byte(jsonld), &myInterface)
	if err != nil {
		log.Println("Error when transforming JSON-LD document to interface:", err)
		return err
	}

	_, err = proc.ToRDF(myInterface, options) // returns triples but toss them, we just want to see if this processes with no err
	if err != nil {
		log.Println("Error when transforming JSON-LD document to RDF:", err)
		return err
	}

	return err
}
