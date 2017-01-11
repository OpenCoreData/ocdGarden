package main

import (
	"log"

	"github.com/minio/minio-go"
)

func main() {
	endpoint := "172.20.42.161:9000"
	accessKeyID := "E48OPWF0ICVNMF4E3UHN"
	secretAccessKey := "qJdWwMTN4ZyO/jwrmueQRUODM51C+WLzW3efr88U"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	// Make a new bucked called coreimages.
	bucketName := "coreimages"
	location := "us-east-1"

	err = minioClient.MakeBucket(bucketName, location)
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, err := minioClient.BucketExists(bucketName)
		if err == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Println("here we are")
			log.Fatalln(err)
		}
	}
	log.Printf("Successfully created or connected to %s\n", bucketName)

	// Upload the zip file
	objectName := "coreref2.png"
	filePath := "/Users/dfils/Google Drive/Documents/Images/coreref.png"
	contentType := "image/png"

	// Upload the zip file with FPutObject
	n, err := minioClient.FPutObject(bucketName, objectName, filePath, contentType)
	if err != nil {
		log.Println(err)
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, n)
}
