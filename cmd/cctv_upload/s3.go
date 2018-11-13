package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/satori/go.uuid"
)

func UploadToS3(region, bucket, filepath string) (string, error) {
	s, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	if err != nil {
		return "", err
	}

	// Change the filename to a UUID
	newFilename := fmt.Sprintf("%s.jpg", uuid.Must(uuid.NewV4()))

	// Upload it
	err = s3Put(s, bucket, filepath, newFilename)
	if err != nil {
		return "", err
	}

	// Return the URL
	return fileURL(region, bucket, newFilename), nil
}

func s3Put(s *session.Session, bucket, filepath, filename string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	_, err = s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:             aws.String(bucket),
		Key:                aws.String(filename),
		ACL:                aws.String("public-read"),
		Body:               bytes.NewReader(buffer),
		ContentLength:      aws.Int64(size),
		ContentType:        aws.String(http.DetectContentType(buffer)),
		ContentDisposition: aws.String("attachment"),
	})

	return err
}

func fileURL(region, bucket, filename string) string {
	return fmt.Sprintf("https://s3-%s.amazonaws.com/%s/%s", region, bucket, filename)
}
