package main

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"log"
	"mime"
	"os"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"

	rdf "github.com/knakk/rdf"
	minio "github.com/minio/minio-go"
	"gopkg.in/mgo.v2"
	ocdstructs "opencoredata.org/ocdGarden/MongoUtils/Mongo2Minio/internal/structs"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile) // show file and line number
}

func main() {
	// call mongo and lookup the redirection to use...
	session, err := getMongoCon()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	mc := minioConnection()
	docRunner(session, mc)
}

// update with material from jsongold
func docRunner(session *mgo.Session, mc *minio.Client) {
	// connect and get metadata documents
	csvw := session.DB("scratch").C("csvwmeta")
	var csvwDocs []ocdstructs.CSVWMeta
	err := csvw.Find(nil).All(&csvwDocs)
	if err != nil {
		log.Printf("this is error %v \n", err)
	}

	// Loop on documents in gridfs
	for _, item := range csvwDocs {
		log.Println(item.URL) // Sprintf a correct URI here
		file, err := session.DB("scratch").GridFS("fs").Open(item.Dc_title)
		if err != nil {
			log.Println(err)
		}

		log.Println(file.Name())

		buf := new(bytes.Buffer)
		buf.ReadFrom(file)
		body := buf.String() // Does a complete copy of the bytes in the buffer.

		log.Println(len(body))

		// Write to minio
		// n, err := writeToMinio(mc, "test", item.Dc_title, body)
		// if err != nil {
		// 	log.Println(err)
		// }
		// log.Println(n)
	}
}

func writeToMinio(mc *minio.Client, bucketName, objectName, body string) (int64, error) {
	// use sha for object and just put object in the metadata?
	// remember looking up on metadata is N where the object name is an index

	// TODO get this from ext on objectName (will always be .csv initially (likely))
	ext := ".csv"

	h := sha1.New()
	h.Write([]byte(body))
	bs := h.Sum(nil)
	bss := fmt.Sprintf("%x", bs) // better way to convert bs hex string to string?

	b := bytes.NewBufferString(body)
	contentType := mimeByType(ext)
	usermeta := make(map[string]string) // what do I want to know?
	usermeta["url"] = ""                // TODO   set this...
	usermeta["sha1"] = bss
	usermeta["filename"] = objectName

	n, err := mc.PutObject(bucketName, objectName, b, int64(b.Len()),
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
	endpoint := "192.168.2.131:9000"
	accessKeyID := "AKIAIOSFODNN7EXAMPLE"
	secretAccessKey := "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
	useSSL := false
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	return minioClient
}

// ref:  https://rosettacode.org/wiki/Strip_control_codes_and_extended_characters_from_a_string#Go
func stripCtlAndExtFromUnicode(str string) string {
	isOk := func(r rune) bool {
		return r < 32 || r >= 127
	}
	// The isOk filter is such that there is no need to chain to norm.NFC
	t := transform.Chain(norm.NFKD, transform.RemoveFunc(isOk))
	// This Transformer could also trivially be applied as an io.Reader
	// or io.Writer filter to automatically do such filtering when reading
	// or writing data anywhere.
	str, _, _ = transform.String(t, str)
	return str
}

func writeFile(name string, tr []rdf.Triple) {
	// Create the output file
	outFile, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	// Write triples to a file
	var inoutFormat rdf.Format
	inoutFormat = rdf.NTriples // Turtle NQuads Ntriples
	enc := rdf.NewTripleEncoder(outFile, inoutFormat)
	err = enc.EncodeAll(tr)
	// err = enc.Encode(newtriple)
	enc.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func getMongoCon() (*mgo.Session, error) {
	host := os.Getenv("machine.dev")
	return mgo.Dial(host)
}
