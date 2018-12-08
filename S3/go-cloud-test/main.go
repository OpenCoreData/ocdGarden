package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/google/go-cloud/blob"
	"github.com/google/go-cloud/blob/gcsblob"
	"github.com/google/go-cloud/blob/s3blob"
	"github.com/google/go-cloud/gcp"
)

// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Command upload saves files to blob storage on GCP and AWS.

func main() {
	// Define our input.
	cloud := flag.String("cloud", "", "Cloud storage to use")
	bucketName := flag.String("bucket", "go-cloud-bucket", "Name of bucket")
	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatal("Failed to provide file to upload")
	}
	file := flag.Arg(0)

	ctx := context.Background()
	// Open a connection to the bucket.
	b, err := setupBucket(context.Background(), *cloud, *bucketName)
	if err != nil {
		log.Fatalf("Failed to setup bucket: %s", err)
	}

	// Prepare the file for upload.
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	w, err := b.NewWriter(ctx, file, nil)
	if err != nil {
		log.Fatalf("Failed to obtain writer: %s", err)
	}
	_, err = w.Write(data)
	if err != nil {
		log.Fatalf("Failed to write to bucket: %s", err)
	}
	if err = w.Close(); err != nil {
		log.Fatalf("Failed to close: %s", err)
	}
}

// setupBucket creates a connection to a particular cloud provider's blob storage.
func setupBucket(ctx context.Context, cloud, bucket string) (*blob.Bucket, error) {
	switch cloud {
	case "aws":
		return setupAWS(ctx, bucket)
	case "gcp":
		return setupGCP(ctx, bucket)
	default:
		return nil, fmt.Errorf("invalid cloud provider: %s", cloud)
	}
}

// setupGCP creates a connection to Google Cloud Storage (GCS).
func setupGCP(ctx context.Context, bucket string) (*blob.Bucket, error) {
	// DefaultCredentials assumes a user has logged in with gcloud.
	// See here for more information:
	// https://cloud.google.com/docs/authentication/getting-started
	creds, err := gcp.DefaultCredentials(ctx)
	if err != nil {
		return nil, err
	}
	c, err := gcp.NewHTTPClient(gcp.DefaultTransport(), gcp.CredentialsTokenSource(creds))
	if err != nil {
		return nil, err
	}
	return gcsblob.OpenBucket(ctx, bucket, c)
}

// setupAWS creates a connection to Simple Cloud Storage Service (S3).
func setupAWS(ctx context.Context, bucket string) (*blob.Bucket, error) {
	c := &aws.Config{
		// Either hard-code the region or use AWS_REGION.
		Region:   aws.String("us-east-2"),
		Endpoint: aws.String("http://localhost:9000"),
		// Need `S3ForcePathStyle: true` and `DisableSSL : true`   (thanks Minio slack!)
		S3ForcePathStyle: aws.Bool(true),
		DisableSSL:       aws.Bool(true),
		// credentials.NewEnvCredentials assumes two environment variables are
		// present:
		// 1. AWS_ACCESS_KEY_ID, and
		// 2. AWS_SECRET_ACCESS_KEY.
		Credentials: credentials.NewEnvCredentials(),
	}
	s := session.Must(session.NewSession(c))
	return s3blob.OpenBucket(ctx, s, bucket)
}
