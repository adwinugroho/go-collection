package services

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"mime/multipart"

	"github.com/adwinugroho/go-aws-s3/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type (
	AWSService struct{}
)

func NewService() *AWSService {
	return &AWSService{}
}

func (service *AWSService) UploadFile(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1"),
		Credentials: credentials.NewStaticCredentials(
			config.AWS_ACCESS_ID,
			config.AWS_SECRET_KEY,
			"",
		),
	})
	if err != nil {
		log.Printf("[UploadFile] error while getting aws session:%+v\n", err)
		return "", err
	}
	size := fileHeader.Size
	buffer := make([]byte, size)
	file.Read(buffer)
	//fmt.Printf("[session AWS] aws sesi: %+v\n", sess)
	getS3 := s3.New(sess)
	fmt.Printf("[s3] get s3: %+v\n", getS3)
	//upload to the s3 bucket
	_, err = getS3.PutObject(&s3.PutObjectInput{
		Bucket: aws.String("diacon-testbucket"),
		Key:    aws.String(fileHeader.Filename),
		//ACL:    aws.String("public-read"),
		Body: bytes.NewReader(buffer),
	})
	if err != nil {
		log.Printf("[UploadFile] error while upload file cause: %+v\n", err)
		return "", err
	}
	filename := fileHeader.Filename
	return filename, nil
}

func (service *AWSService) ListObject() (*s3.ListObjectsOutput, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1"),
		Credentials: credentials.NewStaticCredentials(
			config.AWS_ACCESS_ID,
			config.AWS_SECRET_KEY,
			"",
		),
	})
	if err != nil {
		log.Printf("[ListFile] error while getting aws session:%+v\n", err)
		return nil, err
	}
	getS3 := s3.New(sess)
	//upload to the s3 bucket
	listObjects, err := getS3.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String("diacon-testbucket"),
	})
	if err != nil {
		log.Printf("[ListFile] error while getting list cause: %+v\n", err)
		return nil, err
	}
	return listObjects, nil
}

func (service *AWSService) DeleteObjects() (string, error) {
	var err error
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1"),
		Credentials: credentials.NewStaticCredentials(
			config.AWS_ACCESS_ID,
			config.AWS_SECRET_KEY,
			"",
		),
	})
	if err != nil {
		log.Printf("[DeleteObject] error while getting aws session:%+v\n", err)
		return "", err
	}
	getS3 := s3.New(sess)
	objs, err := getS3.ListObjects(&s3.ListObjectsInput{Bucket: aws.String("diacon-testbucket")})
	if err != nil {
		log.Printf("[DeleteObject] error while getting list object:%+v\n", err)
		return "", err
	}
	var keyString []string
	if len(objs.Contents) == 0 {
		return "", errors.New("data not found")
	}
	for _, obj := range objs.Contents {
		_, err = getS3.DeleteObject(&s3.DeleteObjectInput{
			Bucket: aws.String("diacon-testbucke"),
			Key:    obj.Key,
		})
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
			}
			return "", err
		}
		keyString = append(keyString, *obj.Key)
	}

	var message = fmt.Sprintf("object with keys %+v successfully deleted", keyString)
	return message, nil
}

func (service *AWSService) DeleteObject(key string) (string, error) {
	var err error
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1"),
		Credentials: credentials.NewStaticCredentials(
			config.AWS_ACCESS_ID,
			config.AWS_SECRET_KEY,
			"",
		),
	})
	if err != nil {
		log.Printf("[DeleteObject] error while getting aws session:%+v\n", err)
		return "", err
	}
	getS3 := s3.New(sess)

	_, err = getS3.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String("diacon-testbucket"),
		Key:    aws.String(key), // filename
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return "", err
	}
	var message = "object successfully deleted"
	return message, nil
}
