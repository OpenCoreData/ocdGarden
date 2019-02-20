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
// The only real goal here is to associate the UID of a datapackage with the
// UID of a CSDCO project.
// The project graph will hold the project level metadata and the (and holeid level)
// and the schema.org JSON-LD holds the dataset level structured data.
// These triples are simply to connect the two.  This could really be placed in the
// schema.org JSON-LD too and avoid this step.
func BuildGraph(f []kv.FileMeta) string {
	log.Println("In graph builder")

	// TODO  look at ocdBulk and see how to leverage this!
	// The above is where the old CSDCO graph is made

	// make a graph
	// build the triples, simple triples connecting package URI to project and holeid graphs.
	// at this level nothing more needed, (some prov like stuff) since the triples for the dataset
	// are in the package JSON-LD.
	// connect to the main CSDCO graph  (from excel file)

	// TODO  a simple graph connecting a package URI to the project URI  (nothing fancy needed here)
	// Note the package will have a schame.org based JSON-LD with more details in it.
	// DO I even need this one if I harvest from the package JSON-LD?

	for i := range f {
		log.Println("A random string for the file resource uri, UUID?")
		log.Println(f[i].File) //  just the name as given
		// log.Println(computerSha256(f[i].File))

		log.Println(f[i].Measurement) // use schma.org mesTech ?
		//  http://schema.org/measurementTechnique  "TEXT"

		log.Println(f[i].ProjName) // a UUID for the project (already have one?)  connection predicate?
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
