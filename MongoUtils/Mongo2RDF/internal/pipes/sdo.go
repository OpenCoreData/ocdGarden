package pipes

import (
	"bytes"
	"crypto/sha1"
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"mime"
	"strconv"
	"strings"

	"github.com/knakk/rdf"
	"github.com/minio/minio-go"
	"github.com/piprate/json-gold/ld"

	_ "github.com/mattn/go-sqlite3"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"opencoredata.org/ocdGarden/MongoUtils/Mongo2RDF/internal/common"
	ocdstructs "opencoredata.org/ocdGarden/MongoUtils/Mongo2RDF/internal/structs"
)

// DigitalObject builds a data graph and pulls the byte stream needed to populate a DOC
func DigitalObject(session *mgo.Session) int {
	// connect and get documents
	csvw := session.DB("test").C("schemaorg")
	var schemaDocs []ocdstructs.SchemaOrgMetadata
	err := csvw.Find(nil).All(&schemaDocs)
	if err != nil {
		fmt.Printf("this is error %v \n", err)
	}

	c2 := session.DB("test").C("csvwmeta")

	// set up minio connection
	mc := minioConnection()

	// Open the SQLlite file for the resolving some terms
	// later this should be replaced with SPARQL in a bootstrapped vocabulary
	packageBase := "/home/fils/src/Projects/stonesoup/external/JRSODataDictionary"
	db, err := sql.Open("sqlite3", fmt.Sprintf("%s/data/db/dataDictionary.sqlite3", packageBase))
	if err != nil {
		log.Panic(err)
	}

	// Loop on documents
	for _, item := range schemaDocs {
		// Mongo to Minio Byte Stream
		ab := getMongoBytes(session, item.Name)
		h := sha256.New()
		h.Write(ab)
		s := h.Sum(nil)
		// fmt.Printf("%x", s) // use this SHA in the data graph and to write the object

		// set up the PID for the x-do-meta object..  it needs to be the dgraph meta entry in minio too
		// guid := xid.New()
		// pid := guid.String()
		pid := fmt.Sprintf("%x.jsonld", s)

		resIRI := fmt.Sprintf("https://opencoredata.org/id/jrso/do/%s", pid) // what is here should become the SDO @ID, item.URL

		// write file (digital object) to minio
		n, err := writeToMinio(mc, "jrso", fmt.Sprintf("%x", s), item.Name, ".csv", pid, ab)
		if err != nil {
			log.Println(err)
		}
		log.Printf("Wrote %s to jrso  length %d\n", item.Name, n)

		// Note item.URL changes in this process (from mongo do s3 backed DOC)
		// from item.URL to id/do/[1]

		// Make a context for this graph
		c, _ := rdf.NewIRI(fmt.Sprintf("https://opencoredata.org/id/jrso/objectgraph/%s", fmt.Sprintf("%x", s))) // was pa[3]
		ctx := rdf.Context(c)

		var b strings.Builder

		err = common.IITriple(resIRI, "http://www.w3.org/1999/02/22-rdf-syntax-ns#type", "https://schema.org/Dataset", ctx, &b)
		if err != nil {
			log.Println(err)
		}
		err = common.IITriple(resIRI, "https://schema.org/additionalType", "https://opencoredata.org/voc/janus/v1/Dataset", ctx, &b)
		if err != nil {
			log.Println(err)
		}
		// TODO add subclass of dataset

		// Build a better description
		_, mesDesc, err := dbCall(db, fmt.Sprintf("jrso:%s", item.OpenCoreMeasurement)) // database call
		if err != nil {
			log.Println(err)
		}
		desc := fmt.Sprintf("Leg %s Site %s Hole %s ( %s )  %s", item.OpenCoreLeg,
			item.OpenCoreSite, item.OpenCoreHole,
			getJanusKeyword(item.OpenCoreMeasurement), mesDesc)

		err = common.ILTriple(resIRI, "https://schema.org/description", desc, ctx, &b)
		if err != nil {
			log.Println(err)
		}
		err = common.ILTriple(resIRI, "https://schema.org/keywords", item.Keywords, ctx, &b)
		if err != nil {
			log.Println(err)
		}
		err = common.ILTriple(resIRI, "https://schema.org/license", "https://creativecommons.org/publicdomain/zero/1.0/", ctx, &b)
		if err != nil {
			log.Println(err)
		}
		err = common.ILTriple(resIRI, "https://schema.org/name", item.Name, ctx, &b)
		if err != nil {
			log.Println(err)
		}
		err = common.ILTriple(resIRI, "https://schema.org/url", fmt.Sprintf("https://opencoredata.org/id/jrso/do/%s", fmt.Sprintf("%x", s)), ctx, &b) // was item.URL
		if err != nil {
			log.Println(err)
		}

		// TODO These next three should point to the URI for these
		err = common.ILTriple(resIRI, "https://opencoredata.org/voc/janus/v1/hasLeg", item.OpenCoreLeg, ctx, &b)
		if err != nil {
			log.Println(err)
		}
		err = common.ILTriple(resIRI, "https://opencoredata.org/voc/janus/v1/hasSite", item.OpenCoreSite, ctx, &b)
		if err != nil {
			log.Println(err)
		}
		err = common.ILTriple(resIRI, "https://opencoredata.org/voc/janus/v1/hasHole", item.OpenCoreHole, ctx, &b)
		if err != nil {
			log.Println(err)
		}
		// TODO this should point the IRI of the Measrement (then have a literal triple too?)
		err = common.IITriple(resIRI, "https://schema.org/additionalType", fmt.Sprintf("https://opencoredata.org/voc/janus/v1/%s", item.OpenCoreMeasurement), ctx, &b)
		if err != nil {
			log.Println(err)
		}
		err = common.ILTriple(resIRI, "https://opencoredata.org/voc/janus/v1/measurement", item.OpenCoreMeasurement, ctx, &b)
		if err != nil {
			log.Println(err)
		}

		// LDN Inbox link
		err = common.IITriple(resIRI, "http://www.w3.org/ns/ldp#inbox", fmt.Sprintf("https://opencoredata.org/id/ldn/%s/inbox", fmt.Sprintf("%x", s)), ctx, &b)
		if err != nil {
			log.Println(err)
		}
		// Branches
		err = common.IBTriple(resIRI, "https://schema.org/distribution", "b0", ctx, &b)
		if err != nil {
			log.Println(err)
		}
		// err = common.IBTriple(resIRI, "https://schema.org/publisher", "b1", ctx, &b)
		err = common.IBTriple(resIRI, "https://schema.org/spatialCoverage", "b2", ctx, &b)
		if err != nil {
			log.Println(err)
		}
		err = common.IITriple(resIRI, "https://schema.org/provider", "https://jrso.org", ctx, &b)
		if err != nil {
			log.Println(err)
		}
		err = common.IITriple(resIRI, "https://schema.org/publisher", "https://jrso.org", ctx, &b)
		if err != nil {
			log.Println(err)
		}

		err = common.BITriple("b0", "http://www.w3.org/1999/02/22-rdf-syntax-ns#type", "https://schema.org/DataDownload", ctx, &b)
		if err != nil {
			log.Println(err)
		}
		err = common.BLTriple("b0", "https://schema.org/contentUrl", fmt.Sprintf("https://opencoredata.org/id/jrso/do/%s", fmt.Sprintf("%x", s)), ctx, &b)
		if err != nil {
			log.Println(err)
		}

		// err = common.BITriple("b1", "http://www.w3.org/1999/02/22-rdf-syntax-ns#type", "https://schema.org/Organization", ctx, &b)
		// err = common.BLTriple("b1", "https://schema.org/description", "NSF funded International Ocean Discovery Program operated by JRSO", ctx, &b)
		// err = common.BLTriple("b1", "https://schema.org/name", "International Ocean Discovery Program", ctx, &b)
		// err = common.BLTriple("b1", "https://schema.org/url", "https://iodp.org", ctx, &b)

		err = common.BITriple("b2", "http://www.w3.org/1999/02/22-rdf-syntax-ns#type", "https://schema.org/Place", ctx, &b)
		if err != nil {
			log.Println(err)
		}
		err = common.BBTriple("b2", "https://schema.org/geo", "b3", ctx, &b)
		if err != nil {
			log.Println(err)
		}

		err = common.BITriple("b3", "http://www.w3.org/1999/02/22-rdf-syntax-ns#type", "https://schema.org/GeoCoordinates", ctx, &b)
		if err != nil {
			log.Println(err)
		}
		err = common.BLTriple("b3", "https://schema.org/latitude", item.Spatial.Geo.Latitude, ctx, &b) // TODO update latitude to float type
		if err != nil {
			log.Println(err)
		}
		err = common.BLTriple("b3", "https://schema.org/longitude", item.Spatial.Geo.Longitude, ctx, &b) // TODO update longitutde to float type
		if err != nil {
			log.Println(err)
		}

		// JRSO and publisher and provider
		err = common.IITriple("https://jrso.org", "http://www.w3.org/1999/02/22-rdf-syntax-ns#type", "https://schema.org/Organization", ctx, &b)
		if err != nil {
			log.Println(err)
		}
		err = common.ILTriple("https://jrso.org", "https://schema.org/legalNme", "JOIDES Resoultion Science Operator", ctx, &b)
		if err != nil {
			log.Println(err)
		}
		err = common.ILTriple("https://jrso.org", "https://schema.org/name", "JRSO", ctx, &b)
		if err != nil {
			log.Println(err)
		}
		err = common.IITriple("https://jrso.org", "https://schema.org/sameAs", "https://re3data.org/repository/r3d100010267", ctx, &b)
		if err != nil {
			log.Println(err)
		}
		err = common.ILTriple("https://jrso.org", "https://schema.org/url", "https://iodp.tamu.edu", ctx, &b)
		if err != nil {
			log.Println(err)
		}

		result2 := ocdstructs.CSVWMeta{}
		err = c2.Find(bson.M{"url": item.URL}).One(&result2)
		if err != nil {
			log.Printf("URL lookup error: %v", err)
		}

		cols := result2.TableSchema.Columns
		for x := range cols {
			err = common.IBTriple(resIRI, "https://schema.org/variableMeasured", fmt.Sprintf("vm%d", x), ctx, &b)
			if err != nil {
				log.Println(err)
			}
			err = common.BITriple(fmt.Sprintf("vm%d", x), "http://www.w3.org/1999/02/22-rdf-syntax-ns#type", "https://schema.org/PropertyValue", ctx, &b)
			if err != nil {
				log.Println(err)
			}
			err = common.BLTriple(fmt.Sprintf("vm%d", x), "https://schema.org/name", cols[x].Name, ctx, &b)
			if err != nil {
				log.Println(err)
			}
			err = common.BLTriple(fmt.Sprintf("vm%d", x), "https://schema.org/unitText", cols[x].Datatype, ctx, &b)
			if err != nil {
				log.Println(err)
			}

			_, t2 := thingName(cols[x].Name)
			vmu, vmd, err := dbCall(db, fmt.Sprintf("jrso:%s", t2)) // database call
			if err != nil {
				log.Println(err)
			}
			// log.Printf("Got %s with description:\n%s\n", vmu, vmd)

			// was a BLTriple with just vmu
			err = common.BITriple(fmt.Sprintf("vm%d", x), "https://schema.org/url", exapndJRSO(vmu), ctx, &b) // TODO needs to go to measurement class
			if err != nil {
				log.Println(err)
			}
			// err = common.BLTriple(fmt.Sprintf("vm%d", x), "https://schema.org/description", cols[x].Dc_description, ctx, &b)
			err = common.BLTriple(fmt.Sprintf("vm%d", x), "https://schema.org/description", vmd, ctx, &b)
			if err != nil {
				log.Println(err)
			}

			err = common.BLTriple(fmt.Sprintf("vm%d", x), "https://opencoredata.org/voc/janus/v1/hasColPos", strconv.Itoa(x), ctx, &b)
			if err != nil {
				log.Println(err)
			}
		}

		dg, err := nqToJSONLD(b.String())
		if err != nil {
			log.Println(err)
		}

		// write the data graph to minio
		log.Println(pid)
		n, err = writeToMinio(mc, "jrso", pid, item.Name, ".jsonld", pid, dg)
		if err != nil {
			log.Println(err)
		}
		log.Printf("Wrote %s to jrso metadata length %d\n", item.Name, n)

	}

	// TODO
	// to validate these I would need to get a list of all graphs
	// pull each and convert to JSON-LD and look for errors
	db.Close()
	return 0 //  this is a dumb return....
}

func exapndJRSO(s string) string {
	ss := strings.Split(s, ":")
	return fmt.Sprintf("https://opencoredata.org/voc/janus/v1/%s", ss[len(ss)-1])
}

func getMongoBytes(session *mgo.Session, filename string) []byte {

	file, err := session.DB("test").GridFS("fs").Open(filename)
	if err != nil {
		log.Println(err)
	}

	log.Println(file.Name())

	buf := new(bytes.Buffer)
	buf.ReadFrom(file)

	return buf.Bytes()

}

func nqToJSONLD(triples string) ([]byte, error) {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	// add the processing mode explicitly if you need JSON-LD 1.1 features
	options.ProcessingMode = ld.JsonLd_1_1

	doc, err := proc.FromRDF(triples, options)
	if err != nil {
		log.Println("ERROR: converting from RDF/NQ to JSON-LD")
		log.Println(err)
	}

	// ld.PrintDocument("JSON-LD output", doc)
	b, err := json.MarshalIndent(doc, "", " ")

	return b, err
}

func writeToMinio(mc *minio.Client, bucketName, objectName, origName, ext, guid string, b []byte) (int64, error) {
	// use sha for object and just put object in the metadata?
	// remember looking up on metadata is N where the object name is an index

	// TODO get this from ext on objectName (will always be .csv initially (likely))
	// ext := ".csv"

	h := sha1.New()
	h.Write(b)
	bs := h.Sum(nil)
	bss := fmt.Sprintf("%x", bs) // better way to convert bs hex string to string?

	bb := bytes.NewBuffer(b)
	contentType := mimeByType(ext)
	usermeta := make(map[string]string) // what do I want to know?
	usermeta["url"] = "notDefined"
	usermeta["dgraph"] = fmt.Sprintf("https://opencoredata.org/id/jrso/do/%s", guid)
	usermeta["sha1"] = bss
	usermeta["filename"] = origName

	n, err := mc.PutObject(bucketName, objectName, bb, int64(bb.Len()),
		minio.PutObjectOptions{ContentType: contentType, UserMetadata: usermeta})

	return n, err
}

func mimeByType(e string) string {
	t := mime.TypeByExtension(e)
	if t == "" {
		t = "application/octet-stream"
	}
	return t
}

func minioConnection() *minio.Client {
	// Set up minio and initialize client
	endpoint := "192.168.86.45:32768" // localhost:9000
	accessKeyID := "AKIAIOSFODNN7EXAMPLE"
	secretAccessKey := "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
	useSSL := false
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	return minioClient
}
