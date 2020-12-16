package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"
	"reflect"

	"github.com/piprate/json-gold/ld"
)

func main() {
	fmt.Println("Schema.org JSON-LD concat")

	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")

	jsonld := example()
	var myInterface []map[string]interface{}
	err := json.Unmarshal([]byte(jsonld), &myInterface)
	if err != nil {
		log.Println("Error when transforming JSON-LD document to interface:", err)
	}

	a := toadd()
	var myInterface2 []map[string]interface{}
	err = json.Unmarshal([]byte(a), &myInterface2)
	if err != nil {
		log.Println("Error when transforming JSON-LD append document to interface:", err)
	}

	m := catMaps(myInterface[len(myInterface)-1], myInterface2[len(myInterface2)-1])

	// m := mergeMaps(myInterface[len(myInterface)-1], myInterface2[len(myInterface2)-1])
	// log.Println(m, len(m))

	// m := mergeFragToGraph(myInterface, myInterface2)
	// log.Println(m, len(m))

	r, err := proc.Expand(m, options)
	if err != nil {
		fmt.Print(err)
	}

	b, _ := json.MarshalIndent(r, "", "  ")

	fmt.Println("-----------")
	fmt.Print(string(b))

	// fmt.Print(expanded)

	// ld.PrintDocument("JSON-LD expansion succeeded", expanded)
}

// overwriting duplicate keys, you should handle that if there is a need
func catMaps(mp1, mp2 map[string]interface{}) map[string]interface{} {

	mp3 := make(map[string]interface{})
	for k, v := range mp1 {
		if _, ok := mp1[k]; ok {
			mp3[k] = v
		}
	}

	for k, v := range mp2 {
		if _, ok := mp2[k]; ok {
			mp3[k] = v
		}
	}

	// mp3["@id"] = "https://example.org/test"

	return mp3
}

// overwriting duplicate keys, you should handle that if there is a need
func mergeMaps(maps ...map[string]interface{}) map[string]interface{} {

	result := make(map[string]interface{})
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

func add(a, b interface{}) interface{} {
	return reflect.AppendSlice(reflect.ValueOf(a), reflect.ValueOf(b)).Interface()
}

// need to pass @id
func toadd() string {

	// TODO use text/tempate to add in ID

	const jsldold = `
	 [  {
	  "@context": {
	    "@vocab": "https://schema.org/"
	  },
	  "@id": "{{.}}",
	  "@type": "https://www.schema.org/DigitalDocument",
	  "provider": {
		  "@id": "https://www.csdco.org",
		  "@type": "Organization",
		  "legalName": "Continental Scientific Drilling Coordinating Office",
		  "name": "CSDCO",
		  "sameAs": "https://www.re3data.org/repository/r3d100012874",
		  "url": "https://www.csdco.org"
		},
		"publisher": {
		  "@id": "https://www.csdco.org"
		}
	  }]
	  `

	t := template.Must(template.New("t2").Parse(jsldold))

	t.Execute(os.Stdout, "test")

	return jsldold

}

func example() string {

	const jsldold = `[
		{
		 "@graph": [
		  {
		   "@id": "_:bbj6s49vtr9b9lop9n40g",
		   "@type": [
			"https://schema.org/PropertyValue"
		   ],
		   "https://schema.org/propertyID": [
			{
			 "@value": "SHA256"
			}
		   ],
		   "https://schema.org/value": [
			{
			 "@value": "e88c1fa53004f75e2f1a761284c42e22d69e1750e70904c02c0b726ff42250ad"
			}
		   ]
		  },
		  {
		   "@id": "https://opencoredata.org/id/do/bj6s49vtr9b9lop9n40g",
		   "@type": [
			"https://www.schema.org/DigitalDocument"
		   ],
		   "https://schema.org/additionType": [
			{
			 "@id": "https://opencoredata.org/voc/csdco/v1/ICDFiles"
			}
		   ],
		   "https://schema.org/dateCreated": [
			{
			 "@value": "2009-11-02"
			}
		   ],
		   "https://schema.org/description": [
			{
			 "@value": "Digital object of type ICD named YNP3-ABD09-1A-1L-1.pdf for CSDCO project YNP3"
			}
		   ],
		   "https://schema.org/encodingFormat": [
			{
			 "@value": "application/pdf"
			}
		   ],
		   "https://schema.org/identifier": [
			{
			 "@id": "_:bbj6s49vtr9b9lop9n40g"
			}
		   ],
		   "https://schema.org/isRelatedTo": [
			{
			 "@value": "YNP3"
			}
		   ],
		   "https://schema.org/license": [
			{
			 "@id": "https://example.com/cc0.html"
			}
		   ],
		   "https://schema.org/name": [
			{
			 "@value": "YNP3-ABD09-1A-1L-1.pdf"
			}
		   ],
		   "https://schema.org/url": [
			{
			 "@id": "https://opencoredata.org/id/do/e88c1fa53004f75e2f1a761284c42e22d69e1750e70904c02c0b726ff42250ad"
			}
		   ]
		  }
		 ],
		 "@id": "https://opencoredata.org/objectgraph/id/e88c1fa53004f75e2f1a761284c42e22d69e1750e70904c02c0b726ff42250ad"
		}
	   ]`

	return jsldold

}
