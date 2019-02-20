package minio

import (
	"fmt"
	"log"

	minio "github.com/minio/minio-go"
	"opencoredata.org/ocdGarden/CSDCO/VaultWalker/pkg/utils"
	// faster simd version of sha256 https://github.com/minio/sha256-simd
)

// Do a load to IPFS version?

// LoadToMinio loads jsonld into the specified bucket
func LoadToMinio(filename, bucketName, objectName, project, class, ext, sv string, mc *minio.Client) (int64, error) {
	contentType := utils.MimeByType(ext)

	usermeta := make(map[string]string) // what do I want to know?
	usermeta["filename"] = objectName
	usermeta["sha256"] = sv
	usermeta["project"] = project
	usermeta["class"] = class

	// Upload the file with FPutObject with objectName or sha256 value
	n, err := mc.FPutObject(bucketName, sv, filename, minio.PutObjectOptions{ContentType: contentType, UserMetadata: usermeta})
	if err != nil {
		log.Printf("bucket: %s   objectname: %s \n", bucketName, objectName)
		log.Fatalln(err)
	}

	return n, err
}

func MinioConnection(minioVal, portVal, accessVal, secretVal string) *minio.Client {
	endpoint := fmt.Sprintf("%s:%s", minioVal, portVal)
	accessKeyID := accessVal
	secretAccessKey := secretVal
	useSSL := false
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}
	return minioClient
}
