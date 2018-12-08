package connectors

import (
	"log"

	minio "github.com/minio/minio-go"
)

// Set up minio and initialize client
func MinioConnection() *minio.Client {
	endpoint := "localhost:9111"
	accessKeyID := "AKIAIOSFODNN7EXAMPLE"
	secretAccessKey := "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
	useSSL := false
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}
	return minioClient
}
