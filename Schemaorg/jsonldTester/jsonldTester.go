package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
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
	fmt.Println("Simple testing for jsonld in a web page")

	resPtr := flag.String("res", "", "A URL to check the JSON-LD content of")
	filePtr := flag.String("file", "", "A sitemap file of URLs to check the JSON-LD content of")
	urlPtr := flag.String("url", "", "A URL to a sitemap of URLs to check the JSON-LD content of")

	flag.Parse()

	var urls []string

	if *resPtr != "" {
		urls = append(urls, *resPtr)
	}

	if *filePtr != "" {
		urls = siteMapFile(*filePtr)
	}

	if *urlPtr != "" {
		urls = siteMapURL(*urlPtr)
	}

	start := time.Now()
	for index := range urls {
		url := urls[index]
		err := parseTest(url)
		if err != nil {
			fmt.Printf("URL %s  at index %d has error: %v\n", url, index, err)
		} else {
			fmt.Printf("URL processed %d/%d: %s \n", index, len(urls), url)
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
	// res, err := http.Get(url)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", "EarthCube_DataBot/1.0")

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	doc, err := goquery.NewDocumentFromResponse(res)
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
