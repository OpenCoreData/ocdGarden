package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/kazarena/json-gold/ld"

	yaml "gopkg.in/yaml.v2"
)

type Person struct {
	Name string
}

func main() {
	person := readYAML()
	b, err := jsonldBuilder(person)
	if err != nil {
		fmt.Println("Error when building", err)
	}

	fmt.Println(string(b))
}

func jsonldBuilder(dm Person) ([]byte, error) {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")

	doc := map[string]interface{}{
		"@type":                       "Person",
		"http://schema.org/name":      dm.Name,
		"http://schema.org/jobTitle":  "Professor",
		"http://schema.org/telephone": "(425) 123-4567",
		"http://schema.org/url":       "http://www.janedoe.com",
	}

	context := map[string]interface{}{
		"@context": map[string]interface{}{
			"@vocab":  "http://schema.org/",
			"re3data": "http://example.org/re3data/0.1/",
		},
	}

	compactedDoc, err := proc.Compact(doc, context, options)
	if err != nil {
		fmt.Println("Error when compacting", err)
	}

	return json.MarshalIndent(compactedDoc, "", " ")
}

func readYAML() Person {

	p := Person{}
	data := "name: Douglas Fils"

	err := yaml.Unmarshal([]byte(data), &p)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return p

}
