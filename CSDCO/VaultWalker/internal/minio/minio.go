package minio

import (
	"fmt"
	"log"
	"strings"

	minio "github.com/minio/minio-go"
	"opencoredata.org/ocdGarden/CSDCO/VaultWalker/pkg/utils"
	// faster simd version of sha256 https://github.com/minio/sha256-simd
)

// Do a load to IPFS version?

// LoadToMinio loads jsonld into the specified bucket
// filename: original name of the file
// bucketName:  bucket to put the object in
// objectName: what to name the object that is stored (might be the hash value too)
// project:  A metadata element about a project this file is associated with
// class: The class or type of file ..  might be a scientific measurement class or a type of image
// ext:  the classic .xyz file extension..  used to lookup a mime type to associate with the file
// sv:  the SHA256 hash value
// mc:  A valid minio connection client
func LoadToMinio(filename, bucketName, objectName, project, class, ext, sv, guid string, mc *minio.Client) (int64, error) {
	contentType := utils.MimeByType(ext)

	usermeta := make(map[string]string) // what do I want to know?
	usermeta["filename"] = objectName   //  removeSpace(objectName)
	usermeta["sha256"] = sv
	usermeta["dgraph"] = fmt.Sprintf("http://opencoredata.org/id/do/%s", guid)
	usermeta["project"] = project
	usermeta["class"] = class
	usermeta["resuri"] = fmt.Sprintf("http://opencoredata.org/id/do/%s", sv)

	// Upload the file with FPutObject with objectName or sha256 value
	n, err := mc.FPutObject(bucketName, sv, filename, minio.PutObjectOptions{ContentType: contentType, UserMetadata: usermeta})
	if err != nil {
		log.Printf("bucket: %s   objectname: %s \n", bucketName, objectName)
		log.Println(err)
	}

	return n, err
}

// Note:  this does not replace new lines and tabs..  for that a
// more detailed approach using Fields and Join may be needed
func removeSpace(s string) string {
	return strings.Replace(s, " ", "", -1)
}

// MinioConnection return a connection to the Minio S3 server
func MinioConnection(minioVal, portVal, accessVal, secretVal string) *minio.Client {
	endpoint := fmt.Sprintf("%s:%s", minioVal, portVal)
	accessKeyID := accessVal
	secretAccessKey := secretVal
	useSSL := false
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Println(err)
	}
	return minioClient
}
