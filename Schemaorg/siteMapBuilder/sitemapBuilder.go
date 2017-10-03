package main

import (
	"fmt"
	"os"

	mgo "gopkg.in/mgo.v2"

	ocdstructs "opencoredata.org/ocdGarden/MongoUtils/Mongo2RDF/structs"
)

// SiteMapEntry is a URL that will be registered in a sitemap
type SiteMapEntry struct {
	URL string
}

func main() {
	fmt.Println("Sitemap builder")

	session, err := getMongoCon()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	schemaorg(session)
}

func schemaorg(session *mgo.Session) {
	// connect and get documents
	csvw := session.DB("test").C("schemaorg")
	var dataSetPages []ocdstructs.SchemaOrgMetadata
	err := csvw.Find(nil).All(&dataSetPages)
	if err != nil {
		fmt.Printf("this is error %v \n", err)
	}

	writeFile("sitemap.txt", dataSetPages)

}

func writeFile(name string, dataSetPages []ocdstructs.SchemaOrgMetadata) {
	// Create the output file
	outFile, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	for _, item := range dataSetPages {
		fmt.Println(item.URL)
		outFile.WriteString(fmt.Sprintf("%s\n", item.URL))
	}

}

func getMongoCon() (*mgo.Session, error) {
	host := os.Getenv("MONGO_HOST")
	return mgo.Dial(host)
}
