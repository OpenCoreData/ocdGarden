package pkggrapher

import (
	"archive/zip"
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	minio "github.com/minio/minio-go"
	"github.com/piprate/json-gold/ld"
	"github.com/rs/xid"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"opencoredata.org/ocdGarden/CSDCO/GraphBuilder/internal/connectors"
)

type Manifest struct {
	Profile   string `json:"profile"`
	Resources []struct {
		Encoding string `json:"encoding"`
		Name     string `json:"name"`
		Path     string `json:"path"`
		Profile  string `json:"profile"`
	} `json:"resources"`
}

func PKGGrapher() {
	fmt.Println("Build graph of FDP packages in S3 Bucket")
	mc := connectors.MinioConnection() // minio connection
	multiCall(mc, "packages")
}

func multiCall(mc *minio.Client, bucketname string) {

	// Create a done channel to control 'ListObjectsV2' go routine.
	doneCh := make(chan struct{})

	// Indicate to our routine to exit cleanly upon return.
	defer close(doneCh)

	isRecursive := true
	objectCh := mc.ListObjectsV2(bucketname, "", isRecursive, doneCh)
	var buffer bytes.Buffer
	for object := range objectCh {
		if object.Err != nil {
			fmt.Println(object.Err)
			return
		}
		fo, err := mc.GetObject("packages", object.Key, minio.GetObjectOptions{})
		if err != nil {
			fmt.Println(err)
		}

		oi, err := fo.Stat()
		if err != nil {
			log.Println("Issue with reading an object..  should I just fatal on this to make sure?")
		}

		b, err := getsdo(fo, oi.Size, ".", "metadata/schemaorg.json")
		if err != nil {
			fmt.Println(err)
		}

		nq, err := jsonLDToNQ(string(b), oi.Key)
		if err != nil {
			fmt.Println(err)
		}

		_, err = buffer.WriteString(globalUniqueBNodes(nq))
		if err != nil {
			fmt.Println(err)
		}

		// fmt.Printf("%d bytes written to buffer \n", n)

	}

	writeRDF(buffer.String())
}

func writeRDF(rdf string) (int, error) {
	// for now just append to a file..   later I will send to a triple store
	// If the file doesn't exist, create it, or append to the file

	f, err := os.OpenFile("./data/output/pkgGraph.nt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fl, err := f.Write([]byte(rdf))
	if err != nil {
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	return fl, err // always nil,  we will never get here with FATAL..   leave for test..  but later remove to log only
}

func getsdo(fo *minio.Object, offset int64, dest, target string) ([]byte, error) {
	r, err := zip.NewReader(fo, offset)
	if err != nil {
		return nil, err
	}

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractToBytes := func(f *zip.File) ([]byte, error) {
		rc, err := f.Open()
		if err != nil {
			return nil, err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		b, err := ioutil.ReadAll(rc)
		if err != nil {
			return nil, err
		}

		return b, nil
	}

	var b []byte

	for _, f := range r.File {
		// fmt.Println(f.Name)
		if f.Name == target {
			b, err = extractToBytes(f)
			if err != nil {
				return nil, err
			}
		}
	}

	//TODO  while here, get the []string of files in the zip archive and return that too.
	return b, nil
}

func jsonLDToNQ(jsonld, urlval string) (string, error) {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	options.Format = "application/nquads"

	// --- hack warning ----
	value := gjson.Get(jsonld, "description")
	// fmt.Printf("------------------- %s \n", value.String())
	s := value.String()
	s = strings.TrimPrefix(s, "A CSDCO data package for  project ")
	sa := strings.Split(s, "(")
	fs := strings.Trim(sa[0], " ")

	// kw := gjson.Get(jsonld, "keywords")
	jsonld, _ = sjson.Set(jsonld, "keywords", fmt.Sprintf("%s", fs)) // NOTE   just overwrite keywords for now..  I'll fix the JSON later in the package build process and REMOVE this
	// --- end hack warning ----

	// --- second hack warning ---
	// Need to read the manafest file, then put in all the info about the various files here, including description and short description
	//  this should be done IN PACKAGE CREATION!  (sigh....)

	// get the type string..    append to it based on the

	var myInterface interface{}
	err := json.Unmarshal([]byte(jsonld), &myInterface)
	if err != nil {
		log.Println("Error when transforming %s JSON-LD document to interface:", urlval, err)
		return "", err
	}

	nq, err := proc.ToRDF(myInterface, options) // returns triples but toss them, we just want to see if this processes with no err
	if err != nil {
		log.Println("Error when transforming %s  JSON-LD document to RDF:", urlval, err)
		return "", err
	}

	return nq.(string), err
}

func globalUniqueBNodes(nq string) string {

	scanner := bufio.NewScanner(strings.NewReader(nq))
	// make a map here to hold our old to new map
	m := make(map[string]string)

	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		// parse the line
		split := strings.Split(scanner.Text(), " ")
		sold := split[0]
		oold := split[2]

		if strings.HasPrefix(sold, "_:") { // we are a blank node
			// check map to see if we have this in our value already
			if _, ok := m[sold]; ok {
				// fmt.Printf("We had %s, already\n", sold)
			} else {
				guid := xid.New()
				snew := fmt.Sprintf("_:b%s", guid.String())
				m[sold] = snew
			}
		}

		// scan the object nodes too.. though we should find nothing here.. the above wouldn't
		// eventually find
		if strings.HasPrefix(oold, "_:") { // we are a blank node
			// check map to see if we have this in our value already
			if _, ok := m[oold]; ok {
				// fmt.Printf("We had %s, already\n", oold)
			} else {
				guid := xid.New()
				onew := fmt.Sprintf("_:b%s", guid.String())
				m[oold] = onew
			}
		}
		// triple := tripleBuilder(split[0], split[1], split[3])
		// fmt.Println(triple)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//fmt.Println(m)

	filebytes := []byte(nq)

	for k, v := range m {
		// fmt.Printf("Replace %s with %v \n", k, v)
		filebytes = bytes.Replace(filebytes, []byte(k), []byte(v), -1)
	}

	return string(filebytes)
}
