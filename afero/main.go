package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"log"

	"github.com/spf13/afero"
)

func main() {
	fmt.Println("afero testing")

	var AppFs = afero.NewMemMapFs()

	// create test files and directories
	appFS.MkdirAll("src/a", 0755)
	afero.WriteFile(appFS, "src/a/b", []byte("file b"), 0644)
	afero.WriteFile(appFS, "src/c", []byte("file c"), 0644)

	// I can marshal the files into the above, but can I build
	// out the package?  If I can build the zip in memory maybe I
	//can just verify it afterwards

	// Create a buffer to write our archive to.
	buf := new(bytes.Buffer)

	// Create a new zip archive.
	w := zip.NewWriter(buf)

	// Add some files to the archive.
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	}
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	// Make sure to check the error on Close.
	err := w.Close()
	if err != nil {
		log.Fatal(err)
	}

	// now buf has our zip file
	// we can write it to disk then

}
