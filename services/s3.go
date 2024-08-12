package services

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Serivce struct {
	s3Client *s3.Client
}

type S3Details struct {
	BucketName string
	Key        string
	FileName   string
}

func (s S3Serivce) UploadFile(bucketName, key, fileName string) error {
	log.Println("Upload DB")
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = s.s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   file,
	})
	return err
}

func (s S3Serivce) DownloadFile(bucketName, key, fileName string) error {
	log.Println("Download DB")
	s3Obj, err := s.s3Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		return err
	}
	defer s3Obj.Body.Close()
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(s3Obj.Body)
	if err != nil {
		return err
	}
	_, err = file.Write(body)
	return err
}

func NewS3Service(config aws.Config) *S3Serivce {
	return &S3Serivce{
		s3.NewFromConfig(config),
	}
}

func ConfigAws() (*S3Serivce, *S3Details) {
	awsConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatal("Config Failure")
	}
	s3Svc := NewS3Service(awsConfig)
	// bucketName := os.Getenv("BUCKET_NAME")
	bucketName := "david-christie-enterprise-tf-state-us"
	if bucketName == "" {
		log.Fatal("BUCKET_NAME ENV REQUIRED")
	}
	details := &S3Details{
		BucketName: bucketName,
		Key:        "tmp/davidchristie-dev.db",
		FileName:   "/tmp/davidchristie-dev.db",
	}
	return s3Svc, details
}
