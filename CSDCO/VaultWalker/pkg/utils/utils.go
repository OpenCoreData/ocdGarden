package utils

import (
	"fmt"
	"io"
	"log"
	"mime"
	"os"

	"github.com/minio/sha256-simd"
)

func MimeByType(e string) string {
	t := mime.TypeByExtension(e)
	if t == "" {
		t = "application/octet-stream"
	}
	return t
}

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
