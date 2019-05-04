package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime"
	"os"

	"github.com/minio/sha256-simd"
	"github.com/piprate/json-gold/ld"
)

// MimeByType matches file extensions to mimetype
func MimeByType(e string) string {
	t := mime.TypeByExtension(e)
	if t == "" {
		t = "application/octet-stream"
	}
	return t
}

// NQToJSONLD converts NQ RDF to JSON-LD
func NQToJSONLD(triples string) ([]byte, error) {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	// add the processing mode explicitly if you need JSON-LD 1.1 features
	options.ProcessingMode = ld.JsonLd_1_1

	doc, err := proc.FromRDF(triples, options)
	if err != nil {
		panic(err)
	}

	// ld.PrintDocument("JSON-LD output", doc)
	b, err := json.MarshalIndent(doc, "", " ")

	return b, err
}

// SHAFile returns the sha256 of the string
func SHAFile(s string) string {
	h := sha256.New()

	f, err := os.Open(s)
	if err != nil {
		log.Print(err)
	}
	if _, err := io.Copy(h, f); err != nil { // leverage io.Copy to steam build the hash
		log.Print(err)
	}
	f.Close()

	shavalue := fmt.Sprintf("%x", h.Sum(nil))
	return shavalue
}

// WriteRDF save the RDF graph to a file
func WriteRDF(rdf string) (int, error) {
	// for now just append to a file..   later I will send to a triple store
	// If the file doesn't exist, create it, or append to the file

	RunDir := "."

	// check if our graphs directory exists..  make it if it doesn't
	path := fmt.Sprintf("%s/output", RunDir)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}

	f, err := os.OpenFile(fmt.Sprintf("%s/output/objectGraph.nq", RunDir), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}

	fl, err := f.Write([]byte(rdf))
	if err != nil {
		log.Println(err)
	}

	if err := f.Close(); err != nil {
		log.Println(err)
	}

	return fl, err
}
