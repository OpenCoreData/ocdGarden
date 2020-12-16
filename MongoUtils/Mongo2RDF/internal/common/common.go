package common

import (
	"fmt"
	"log"
	"os"

	"github.com/knakk/rdf"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	mgo "gopkg.in/mgo.v2"
)

// StripCtlAndExtFromUnicode ref:  https://rosettacode.org/wiki/Strip_control_codes_and_extended_characters_from_a_string#Go
func StripCtlAndExtFromUnicode(str string) string {
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

// WriteToName save the RDF graph to a file
func WriteToName(rdf, filename string) (int, error) {
	// for now just append to a file..   later I will send to a triple store
	// If the file doesn't exist, create it, or append to the file

	RunDir := "."

	// check if our graphs directory exists..  make it if it doesn't
	path := fmt.Sprintf("%s/output", RunDir)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}

	f, err := os.OpenFile(fmt.Sprintf("%s/output/%s", RunDir, filename), os.O_CREATE|os.O_WRONLY, 0644)
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

// WriteFile file name
func WriteFile(name string, tr []rdf.Triple) {
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

// GetMongoCon provides a mongo session connection
func GetMongoCon() (*mgo.Session, error) {
	host := os.Getenv("machine.dev")
	return mgo.Dial(host)
}
