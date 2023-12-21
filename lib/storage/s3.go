package storage

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3 struct {
	client *s3.S3
	bucket string
}

func NewS3(
	client *s3.S3,
	bucket string,
) *S3 {
	return &S3{
		client: client,
		bucket: bucket,
	}
}

func (s *S3) Put(ctx context.Context, key string, contentType string, blob io.ReadSeeker) error {
	if _, err := s.client.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(key),
		ContentType: aws.String(contentType),
		Body:        blob,
	}); err != nil {
		return err
	}

	return nil
}

func (s *S3) Get(ctx context.Context, key string) (io.ReadCloser, error) {
	out, err := s.client.GetObjectWithContext(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}

	return out.Body, nil
}

func (s *S3) Delete(ctx context.Context, key string) error {
	if _, err := s.client.DeleteObjectWithContext(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	}); err != nil {
		return err
	}

	return nil
}
