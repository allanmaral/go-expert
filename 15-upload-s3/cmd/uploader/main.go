package main

import (
	"io"
	"log"
	"os"
	"path"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	s3Client *s3.S3
	s3Bucket string
	wg       sync.WaitGroup
)

func init() {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String("us-east-1"),
			Credentials: credentials.NewStaticCredentials(
				"---",
				"---",
				"",
			),
		},
	)
	if err != nil {
		log.Fatalf("failed to open session: %v", err)
	}

	s3Client = s3.New(sess)
	s3Bucket = "goexpert-example-bucket"
}

func main() {
	dir, err := os.Open("./tmp")
	if err != nil {
		log.Fatalf("failed to open tmp directory: %v", err)
	}
	defer dir.Close()

	uploadControl := make(chan struct{}, 100)

	for {
		files, err := dir.ReadDir(1)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("failed to read directory content: %v", err)
			continue
		}

		wg.Add(1)
		uploadControl <- struct{}{}
		go uploadFile(files[0].Name(), uploadControl)
	}
	wg.Wait()

}

func uploadFile(filename string, uploadControl <-chan struct{}) {
	defer wg.Done()
	log.Printf("uploading file %s to bucket %s\n", filename, s3Bucket)

	path := path.Join("tmp", filename)
	f, err := os.Open(path)
	if err != nil {
		log.Printf("failed to open file %s: %v\n", filename, err)
		<-uploadControl
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
		<-uploadControl
		return
	}
	log.Printf("successfully uploaded file %s\n", filename)
	<-uploadControl
}
