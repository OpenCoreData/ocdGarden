package main

import (
	"flag"
	"log"

	"github.com/minio/minio-go"
)

func main() {

	namePtr := flag.String("name", "", "Name for the object, might be a hash to ID the file with.")
	pathPtr := flag.String("path", "", "Path to file to upload")
	typePtr := flag.String("type", "", "Valid mimetype entry")

	// if *objectNamePtr == "" || *filePathPtr == "" || *contentTypePtr == "" {
	// 	log.Fatalln("please fill out all options")
	// }

	flag.Parse()

	objectName := *namePtr
	filePath := *pathPtr
	contentType := *typePtr

	log.Printf("Load %s from %s at type %s \n", objectName, filePath, contentType)

	endpoint := "oss.opencoredata.org"
	accessKeyID := "AKIAIOSFODNN7JASUINM"
	secretAccessKey := "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"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	// Make a new bucked called coreimages.
	bucketName := "tier0"
	location := "us-east-1" // not uses in minio..  an Amazon S3 item

	err = minioClient.MakeBucket(bucketName, location)
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, err := minioClient.BucketExists(bucketName)
		if err == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Println("Fatal error..   bad place to be")
			log.Fatalln(err)
		}
	}
	log.Printf("Successfully created or connected to %s\n", bucketName)

	// Upload the zip file with FPutObject
	n, err := minioClient.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Println(err)
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, n)
}
