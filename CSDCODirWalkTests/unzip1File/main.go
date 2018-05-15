package main

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"log"

	minio "github.com/minio/minio-go"
)

func main() {
	// /media/fils/seagate/datavolume/e946a52c6c955450184e72bc3a2208242ec905e1.zip
	// fmt.Println("vim-go")
	// Unzip("/media/fils/seagate/datavolume/e946a52c6c955450184e72bc3a2208242ec905e1.zip", ".")

	mc := miniConnection() // minio connection

	fo, err := mc.GetObject("packages", "1e62a870ca7423ae175018e6c68e19afd3aa7f11.zip", minio.GetObjectOptions{})
	if err != nil {
		fmt.Println(err)
	}

	oi, err := fo.Stat()
	if err != nil {
		log.Println("Issue with reading an object..  should I just fatal on this to make sure?")
	}

	fmt.Println(oi)

	b, err := unzip(fo, oi.Size, ".")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(b))
}

// func unzip(src, dest string) error {
func unzip(fo *minio.Object, offset int64, dest string) ([]byte, error) {

	// r, err := zip.OpenReader(src) // can I pass a minio resouce here?  Use readerat from minio (get size from info call)
	r, err := zip.NewReader(fo, offset)
	if err != nil {
		return nil, err
	}

	// defer func() {
	// 	if err := r.Close(); err != nil {
	// 		panic(err)
	// 	}
	// }()

	// os.MkdirAll(dest, 0755)

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

	// extractAndWriteFile := func(f *zip.File) error {
	// 	rc, err := f.Open()
	// 	if err != nil {
	// 		return err
	// 	}
	// 	defer func() {
	// 		if err := rc.Close(); err != nil {
	// 			panic(err)
	// 		}
	// 	}()

	// 	path := filepath.Join(dest, f.Name)

	// 	if f.FileInfo().IsDir() {
	// 		os.MkdirAll(path, f.Mode())
	// 	} else {
	// 		os.MkdirAll(filepath.Dir(path), f.Mode())
	// 		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	// 		if err != nil {
	// 			return err
	// 		}
	// 		defer func() {
	// 			if err := f.Close(); err != nil {
	// 				panic(err)
	// 			}
	// 		}()

	// 		_, err = io.Copy(f, rc)
	// 		if err != nil {
	// 			return err
	// 		}
	// 	}
	// 	return nil
	// }

	//err := extractAndWriteFile(f)
	//if err != nil {
	//		return err
	//	}

	var b []byte

	for _, f := range r.File {
		// fmt.Println(f.Name)
		// if f.Name == "metadata/schemaorg.json" {
		if f.Name == "datapackage.json" {

			// err := extractAndWriteFile(f)
			b, err = extractToBytes(f)

			if err != nil {
				return nil, err
			}
		}
	}

	return b, nil
}

// Set up minio and initialize client
func miniConnection() *minio.Client {
	endpoint := "localhost:9000"
	accessKeyID := "AKIAIOSFODNN7EXAMPLE"
	secretAccessKey := "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
	useSSL := false
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	return minioClient
}
