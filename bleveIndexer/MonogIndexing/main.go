package main

import (
	"gopkg.in/mgo.v2"
	// "encoding/json"
	"fmt"
	"os"
	// "gopkg.in/mgo.v2/bson"
	"github.com/blevesearch/bleve"
	ocdServices "opencoredata.org/ocdServices/documents"
)

func main() {
	// call mongo and lookup the redirection to use...
	session, err := GetMongoCon()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	IndexCSVW(session)
	// IndexSchema(session)

}

func IndexCSVW(session *mgo.Session) {
	// Optional. Switch the session to a monotonic behavior.
	csvw := session.DB("test").C("csvwmeta")

	// Find all Documents CSVW
	var csvwDocs []ocdServices.CSVWMeta
	err := csvw.Find(nil).All(&csvwDocs)
	if err != nil {
		fmt.Printf("this is error %v \n", err)
	}

	// open a new index
	mapping := bleve.NewIndexMapping()
	// analyzer := mapping.Ad

	index, berr := bleve.New("csvw.bleve", mapping)
	if berr != nil {
		fmt.Printf("Bleve error making index %v \n", berr)
	}

	// index some data
	for i, item := range csvwDocs {
		berr = index.Index(item.URL, item)
		fmt.Printf("Indexed item %d with URL %s\n", i, item.URL)
		if berr != nil {
			fmt.Printf("Bleve error indexing %v \n", berr)
		}
	}

}

func IndexSchema(session *mgo.Session) {
	// Optional. Switch the session to a monotonic behavior.
	schemaorg := session.DB("test").C("schemaorg")

	// Find all Documents schema.org
	var schemaDocs []ocdServices.SchemaOrgMetadata
	err := schemaorg.Find(nil).All(&schemaDocs)
	if err != nil {
		fmt.Printf("this is error %v \n", err)
	}

	// open a new index
	mapping := bleve.NewIndexMapping()
	index, berr := bleve.New("schema.bleve", mapping)
	if berr != nil {
		fmt.Printf("Bleve error making index %v \n", berr)
	}

	// index some data
	for i, item := range schemaDocs {
		berr = index.Index(item.URL, item)
		fmt.Printf("Indexed item %d with URL %s\n", i, item.URL)
		if berr != nil {
			fmt.Printf("Bleve error indexing %v \n", berr)
		}
	}
}

func GetMongoCon() (*mgo.Session, error) {
	host := os.Getenv("MONGO_HOST")
	return mgo.Dial(host)
}
