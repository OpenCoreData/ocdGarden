package main

import (
	"encoding/xml"
	"fmt"
	"os"

	mgo "gopkg.in/mgo.v2"

	ocdstructs "opencoredata.org/ocdGarden/MongoUtils/Mongo2RDF/structs"
)

type XMLSitemapURLset struct {
	XMLName   xml.Name `xml:"urlset"`
	Namespace string   `xml:"xmlns,attr"`
	URLset    []XMLSiteMapURL
}

type XMLSiteMapURL struct {
	XMLName    xml.Name `xml:"url"`
	Location   string   `xml:"loc"`
	LastMod    string   `xml:"lastmod,omitempty"`    // optional W3C datetime format:  can use YYYY-MM-DD
	ChangeFreq string   `xml:"changefreq,omitempty"` // optional one of: always hourly	daily weekly monthly yearly never
	Priority   float32  `xml:"priority,omitempty"`   // optional 0.0 to 1.0  default 0.5
}

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

	writeFile("sitemap.xml", dataSetPages)
}

func writeFile(name string, dataSetPages []ocdstructs.SchemaOrgMetadata) {
	// Create the output file
	outFile, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	// Loop through our schema.org struct and make a set of XMLSitemap structs
	urlset := []XMLSiteMapURL{}
	for _, item := range dataSetPages {
		entry := XMLSiteMapURL{Location: item.URL}
		urlset = append(urlset, entry)
		// outFile.WriteString(fmt.Sprintf("%s\n", item.URL))
	}

	sitemap := XMLSitemapURLset{URLset: urlset, Namespace: "http://www.sitemaps.org/schemas/sitemap/0.9"}
	output, err := xml.MarshalIndent(sitemap, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	outFile.WriteString("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n")
	outFile.Write(output)
	// os.Stdout.Write(output)
}

func getMongoCon() (*mgo.Session, error) {
	host := os.Getenv("MONGO_HOST")
	return mgo.Dial(host)
}
