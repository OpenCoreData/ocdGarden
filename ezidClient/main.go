package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// go run main.go -username=apitest -password=password
func main() {
	username := flag.String("username", "apitest", "EZID username")
	passwd := flag.String("password", "password", "EZID password")
	flag.Parse()

	posttest(*username, *passwd)
}

func posttest(username, passwd string) {
	url := "https://ezid.cdlib.org/shoulder/doi:10.5072/FK2"
	fmt.Println("URL:>", url)

	b, err := ioutil.ReadFile("datacite-example-dataset-v3.0.xml") // just pass the file name
	// b, err := ioutil.ReadFile("test.anvl") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	s := string(b)
	packageString := fmt.Sprintf("datacite: %s", strings.Replace(s, "\n", " ", -1)) // note % must also be escaped

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(packageString)))
	req.Header.Set("Content-Type", "text/plain")
	req.SetBasicAuth(username, passwd)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err) // Might as well panic..  things have gone dreadfully wrong....
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
