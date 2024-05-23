package storage

import (
    "bytes"
    "fmt"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
    "coca-cola-lambda/internal/ports"
)

type AWSS3 struct {
    client *s3.S3
}

func NewAWSS3(region string) ports.Storage {
    sess := session.Must(session.NewSession(&aws.Config{
        Region: aws.String(region),
    }))
    return &AWSS3{
        client: s3.New(sess),
    }
}

func (s *AWSS3) GetObject(bucket, key string) (string, error) {
    s3Result, err := s.client.GetObject(&s3.GetObjectInput{
        Bucket: aws.String(bucket),
        Key:    aws.String(key),
    })
    if err != nil {
        return "", fmt.Errorf("failed to get object from S3: %w", err)
    }

    defer s3Result.Body.Close()
    buf := new(bytes.Buffer)
    buf.ReadFrom(s3Result.Body)
    return buf.String(), nil
}
