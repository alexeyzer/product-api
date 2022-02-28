package client

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
	"os"
)

type S3Client interface {
	UploadFileD(ctx context.Context, key string, file io.Reader, contentType string) (*s3manager.UploadOutput, error)
}

type s3Client struct {
	bucketName string
	client     *s3.S3
	uploader   *s3manager.Uploader
}

func NewS3Client(bucketName, secretID, secretKey, region, endpoint string) (S3Client, error) {
	session, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(os.Getenv(secretID), os.Getenv(secretKey), ""),
		Region:      aws.String(region),
		Endpoint:    aws.String(endpoint),
	},
	)
	if err != nil {
		return nil, err
	}
	client := s3.New(session)
	uploader := s3manager.NewUploaderWithClient(client, func(u *s3manager.Uploader) {})

	return &s3Client{client: client, uploader: uploader, bucketName: bucketName}, nil
}

func (s *s3Client) UploadFileD(ctx context.Context, key string, file io.Reader, contentType string) (*s3manager.UploadOutput, error) {
	uploadResponse, err := s.uploader.UploadWithContext(ctx, &s3manager.UploadInput{
		ACL:         aws.String("public-read"),
		Body:        file,
		Bucket:      &s.bucketName,
		ContentType: &contentType,
		Key:         &key,
	})
	if err != nil {
		return nil, err
	}

	return uploadResponse, nil
}
