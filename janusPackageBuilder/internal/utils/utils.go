package utils

import (
	"bytes"
	"log"

	minio "github.com/minio/minio-go"
)

// Set up minio and initialize client
func MinioConnection(ep, ak, sk string) *minio.Client {
	endpoint := ep
	accessKeyID := ak
	secretAccessKey := sk
	useSSL := false
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}
	return minioClient
}

// LoadToMinio loads jsonld into the specified bucket
func LoadToMinio(object, bucketName, objectName string, mc *minio.Client) (int64, error) {

	// set up some elements for PutObject
	contentType := "application/ld+json"
	usermeta := make(map[string]string) // what do I want to know?
	b := bytes.NewBufferString(object)
	// usermeta["url"] = urlloc
	// usermeta["sha1"] = bss

	// Upload the zip file with FPutObject
	n, err := mc.PutObject(bucketName, objectName, b, int64(b.Len()), minio.PutObjectOptions{ContentType: contentType, UserMetadata: usermeta})
	if err != nil {
		log.Printf("%s", objectName)
		log.Fatalln(err)
	}

	// log.Printf("#%d Uploaded Bucket:%s File:%s Size %d\n", i, bucketName, objectName, n)

	return n, nil
}
