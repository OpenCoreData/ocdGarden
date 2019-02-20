package main

import (
	"flag"
	"fmt"
	"os"

	myminio "../internal/minio"
	minio "github.com/minio/minio-go"
)

var minioVal, portVal, accessVal, secretVal, bucketVal string
var uploadVal, sslVal bool

func init() {
	akey := os.Getenv("MINIO_ACCESS_KEY")
	skey := os.Getenv("MINIO_SECRET_KEY")

	flag.StringVar(&minioVal, "address", "192.168.2.131", "FQDN for server")
	flag.StringVar(&portVal, "port", "9000", "Port for minio server, default 9000")
	flag.StringVar(&accessVal, "access", akey, "Access Key ID")
	flag.StringVar(&secretVal, "secret", skey, "Secret access key")
	flag.StringVar(&bucketVal, "bucket", "csdco", "The configuration bucket")
	flag.BoolVar(&sslVal, "ssl", false, "Use SSL boolean")
	flag.BoolVar(&uploadVal, "upload", false, "Upload files to object store")
}

func main() {
	fmt.Println("Build packages from metadata")

	flag.Parse()
	mc := myminio.MinioConnection(minioVal, portVal, accessVal, secretVal)

	//scanAll(mc)
	pkgFiles(mc, "CAHO")

}

func pkgFiles(mc *minio.Client, proj string) {
	// Create a done channel to control 'ListObjectsV2' go routine.
	doneCh := make(chan struct{})

	// Indicate to our routine to exit cleanly upon return.
	defer close(doneCh)

	isRecursive := true
	objectCh := mc.ListObjectsV2("csdco", "", isRecursive, doneCh)
	for object := range objectCh {
		if object.Err != nil {
			fmt.Println(object.Err)
			return
		}

		objInfo, err := mc.StatObject("csdco", object.Key, minio.StatObjectOptions{})
		if err != nil {
			fmt.Println(err)
			return
		}

		pim := objInfo.Metadata["X-Amz-Meta-Project"][0]
		if pim == proj {
			fmt.Println(object.Key)
			fmt.Println(objInfo.Metadata)
		}
	}

	// TODO   not take this collection of files (place in an array or struct) and generate
	// a Frictonless Data package package.json (?) manifest from it.  (a virtual package)

}

func scanAll(mc *minio.Client) {
	// Create a done channel to control 'ListObjectsV2' go routine.
	doneCh := make(chan struct{})

	// Indicate to our routine to exit cleanly upon return.
	defer close(doneCh)

	isRecursive := true
	objectCh := mc.ListObjectsV2("csdco", "", isRecursive, doneCh)
	for object := range objectCh {
		if object.Err != nil {
			fmt.Println(object.Err)
			return
		}

		fmt.Println(object.Key)
		objInfo, err := mc.StatObject("csdco", object.Key, minio.StatObjectOptions{})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(objInfo.Metadata["X-Amz-Meta-Project"][0])
	}
}
