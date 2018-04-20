package graph

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"

	"opencoredata.org/ocdGarden/CSDCODirWalkTests/godirwalk/kv"
)

// BuildGraph generates the graph of file structured data
func BuildGraph(f []kv.FileMeta) string {
	log.Println("In graph builder")

	// make a graph
	// build the triples
	// connect to the main CSDCO graph  (from excel file)

	for i := range f {
		log.Println("A random string for the file resource uri, UUID?")
		log.Println(f[i].File) //  just the name as given
		// log.Println(computerSha256(f[i].File))

		log.Println(f[i].Measurement) // use schma.org mesTech ?
		//  http://schema.org/measurementTechnique  "TEXT"

		log.Println(f[i].ProjName) // a UUID for the project (already have one?)
		// log.Println(f[i].Valid)  // ignore.. code use only
	}

	return "generated graph"
}

// based on projname, look up resource URI associated with it and return
func projLookup(proj string) string {

	// need to SPARQL or text look up the URI based on the proj name
	//  problem..  there is not a UNIQUE one at this level
	// will need to build one in the main graph.
	// find the code that built this!

	return proj + ": resold URI"
}

func computeSha256(fp string) {
	f, err := os.Open(fp)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%x", h.Sum(nil))
}
