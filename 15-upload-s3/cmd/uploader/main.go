package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	s3Client *s3.S3
	s3Bucket string
)

func init() {
	session, err := session.NewSession(
		&aws.Config{
			Region:      aws.String("us-east-1"),
			Credentials: &credentials.Credentials{},
		},
	)
	if err != nil {
		log.Fatalf("failed to open a session: %v", err)
	}

	s3Client = s3.New(session)
	s3Bucket = "goexpert-example-bucket"
}

func main() {
	dir, err := os.Open("./tmp")
	if err != nil {
		log.Fatalf("failed to open tmp directory: %v", err)
	}
	defer dir.Close()

}

func uploadFile(filename string) {
	log.Printf("uploading file %s to bucket %s\n", filename, s3Bucket)

	path := path.Join("tmp", filename)
	f, err := os.Open(path)
	if err != nil {
		log.Printf("failed to open file %s: %v\n", filename, err)
		return
	}
	defer f.Close()

	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(filename),
		Body:   f,
	})
	if err != nil {
		log.Printf("failed to upload file %s: %v\n", filename, err)
		return
	}
	log.Printf("successfully uploaded file %s\n", filename)
}
