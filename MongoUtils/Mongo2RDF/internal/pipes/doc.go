package pipes

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"

	"github.com/fatih/camelcase"
	"github.com/knakk/rdf"

	// _ "github.com/mattn/go-sqlite3"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"opencoredata.org/ocdGarden/MongoUtils/Mongo2RDF/internal/common"
	ocdstructs "opencoredata.org/ocdGarden/MongoUtils/Mongo2RDF/internal/structs"
)

// Schemaorg update with material from jsongold
func Schemaorg(session *mgo.Session, ub *common.Buffer) int {
	var b strings.Builder
	// connect and get documents
	csvw := session.DB("test").C("schemaorg")
	var schemaDocs []ocdstructs.SchemaOrgMetadata
	err := csvw.Find(nil).All(&schemaDocs)
	if err != nil {
		fmt.Printf("this is error %v \n", err)
	}

	c2 := session.DB("test").C("csvwmeta")

	// Open the SQLlite file for the resolving some terms
	// later this should be replaced with SPARQL in a bootstrapped vocabulary
	packageBase := "/home/fils/src/Projects/stonesoup/external/JRSODataDictionary"
	db, err := sql.Open("sqlite3", fmt.Sprintf("%s/data/db/dataDictionary.sqlite3", packageBase))
	if err != nil {
		log.Panic(err)
	}

	// Loop on documents
	for _, item := range schemaDocs {
		// make resource IRI
		resIRI := fmt.Sprintf("%s", item.URL) // what is here should become the SDO @ID

		// Make a context for this graph
		u, _ := url.Parse(resIRI)
		pa := strings.Split(u.Path, "/")
		c, _ := rdf.NewIRI(fmt.Sprintf("http://opencoredata.org/id/objectgraph/%s", pa[3]))
		ctx := rdf.Context(c)

		_ = common.IITriple(resIRI, "http://www.w3.org/1999/02/22-rdf-syntax-ns#type", "http://schema.org/Dataset", ctx, &b)
		_ = common.IITriple(resIRI, "http://www.w3.org/1999/02/22-rdf-syntax-ns#type", "http://opencoredata.org/voc/janus/v1/Dataset", ctx, &b)
		// TODO add subclass of dataset
		_ = common.ILTriple(resIRI, "http://schema.org/description", item.Description, ctx, &b)
		_ = common.ILTriple(resIRI, "http://schema.org/keywords", item.Keywords, ctx, &b)
		_ = common.ILTriple(resIRI, "http://schema.org/license", "https://creativecommons.org/publicdomain/zero/1.0/", ctx, &b)
		_ = common.ILTriple(resIRI, "http://schema.org/name", item.Name, ctx, &b)
		_ = common.ILTriple(resIRI, "http://schema.org/url", item.URL, ctx, &b)
		_ = common.IBTriple(resIRI, "http://schema.org/distribution", "b0", ctx, &b)
		_ = common.IBTriple(resIRI, "http://schema.org/publisher", "b1", ctx, &b)
		_ = common.IBTriple(resIRI, "http://schema.org/spatialCoverage", "b2", ctx, &b)

		// TODO These next three should point to the URI for these
		_ = common.ILTriple(resIRI, "http://opencoredata.org/voc/janus/v1/hasLeg", item.OpenCoreLeg, ctx, &b)
		_ = common.ILTriple(resIRI, "http://opencoredata.org/voc/janus/v1/hasSite", item.OpenCoreSite, ctx, &b)
		_ = common.ILTriple(resIRI, "http://opencoredata.org/voc/janus/v1/hasHole", item.OpenCoreHole, ctx, &b)

		// TODO this should point the IRI of the Measrement (then have a literal triple too?)
		_ = common.IITriple(resIRI, "http://www.w3.org/1999/02/22-rdf-syntax-ns#type", fmt.Sprintf("http://opencoredata.org/voc/janus/v1/%s", item.OpenCoreMeasurement), ctx, &b)
		_ = common.ILTriple(resIRI, "http://opencoredata.org/voc/janus/v1/measurement", item.OpenCoreMeasurement, ctx, &b)

		// LDN Inbox link
		_ = common.IITriple(resIRI, "http://www.w3.org/ns/ldp#inbox", fmt.Sprintf("http://opencoredata.org/id/ldn/%s/inbox", pa[3]), ctx, &b)

		_ = common.BITriple("b0", "http://www.w3.org/1999/02/22-rdf-syntax-ns#type", "http://schema.org/DataDownload", ctx, &b)
		_ = common.BLTriple("b0", "http://schema.org/contentUrl", "http://opencoredata.org/api/v1/documents/download/204_1244B_JanusThermalConductivity_rkfjjNYV.csv", ctx, &b)

		_ = common.BITriple("b1", "http://www.w3.org/1999/02/22-rdf-syntax-ns#type", "http://schema.org/Organization", ctx, &b)
		_ = common.BLTriple("b1", "http://schema.org/description", "NSF funded International Ocean Discovery Program operated by JRSO", ctx, &b)
		_ = common.BLTriple("b1", "http://schema.org/name", "International Ocean Discovery Program", ctx, &b)
		_ = common.BLTriple("b1", "http://schema.org/url", "http://iodp.org", ctx, &b)

		_ = common.BITriple("b2", "http://www.w3.org/1999/02/22-rdf-syntax-ns#type", "http://schema.org/Place", ctx, &b)
		_ = common.BLTriple("b2", "http://schema.org/geo", "b3", ctx, &b)

		_ = common.BITriple("b3", "http://www.w3.org/1999/02/22-rdf-syntax-ns#type", "http://schema.org/GeoCoordinates", ctx, &b)
		_ = common.BLTriple("b3", "http://schema.org/latitude", "44.59", ctx, &b)
		_ = common.BLTriple("b3", "http://schema.org/longitude", "-125.12", ctx, &b)

		result2 := ocdstructs.CSVWMeta{}
		err = c2.Find(bson.M{"url": resIRI}).One(&result2)
		if err != nil {
			log.Printf("URL lookup error: %v", err)
		}

		cols := result2.TableSchema.Columns
		for x := range cols {
			_ = common.IBTriple(resIRI, "http://schema.org/variableMeasured", fmt.Sprintf("vm%d", x), ctx, &b)
			_ = common.BITriple(fmt.Sprintf("vm%d", x), "http://www.w3.org/1999/02/22-rdf-syntax-ns#type", "http://schema.org/PropertyValue", ctx, &b)
			_ = common.BLTriple(fmt.Sprintf("vm%d", x), "http://schema.org/name", cols[x].Name, ctx, &b)
			_ = common.BLTriple(fmt.Sprintf("vm%d", x), "http://schema.org/unitText", cols[x].Datatype, ctx, &b)

			_, t2 := thingName(cols[x].Name)
			vmu, vmd, err := dbCall(db, fmt.Sprintf("jrso:%s", t2))
			if err != nil {
				log.Println(err)
			}
			// log.Printf("Got %s with description:\n%s\n", vmu, vmd)

			// TODO should not use jrso: in quads workflow..   need to expand this this prefix
			_ = common.BLTriple(fmt.Sprintf("vm%d", x), "http://schema.org/url", vmu, ctx, &b) // TODO needs to go to measurement class
			// _ = common.BLTriple(fmt.Sprintf("vm%d", x), "http://schema.org/description", cols[x].Dc_description, ctx, &b)
			err = common.BLTriple(fmt.Sprintf("vm%d", x), "http://schema.org/description", vmd, ctx, &b)
			if err != nil {
				log.Println(err)
			}

			_ = common.BLTriple(fmt.Sprintf("vm%d", x), "http://opencoredata.org/voc/janus/v1/hasColPos", strconv.Itoa(x), ctx, &b)
		}

	}

	len, err := ub.Write([]byte(b.String()))
	if err != nil {
		log.Printf("error in the buffer write... %v\n", err)
	}

	// TODO
	// to validate these I would need to get a list of all graphs
	// pull each and convert to JSON-LD and look for errors
	db.Close()
	return len
}

func dbCall(db *sql.DB, term string) (string, string, error) {
	statement, _ := db.Prepare("SELECT URI, description as pdesc FROM classes where URI = ? LIMIT 1")
	rows, err := statement.Query(term)
	if err != nil {
		log.Println(err)
	}
	var URI, pdesc string

	for rows.Next() {
		err = rows.Scan(&URI, &pdesc)
		if err != nil {
			log.Println(err)
		}
	}

	rows.Close() //good habit to close
	return URI, pdesc, err
}

// add errors..  trim, spaces to _, check and handle len 1 or len 2
func thingName(name string) (string, string) {
	if strings.Contains(name, ":") {
		sn := strings.Split(name, ":")
		return sn[0], sn[1]
	} else {
		return "", name
	}
}

// split the camel case string into various words
func getJanusKeyword(s string) string {
	ssplit := strings.Split(s, ",")
	var targetString string
	for _, element := range ssplit {
		if strings.Contains(strings.ToLower(element), "janus") {
			// targetString = strings.ToLower(element)

			splitted := camelcase.Split(element)          // hackish split of terms
			targetString = strings.Join(splitted[:], " ") //

		}
	}
	return strings.TrimSpace(targetString)
}
