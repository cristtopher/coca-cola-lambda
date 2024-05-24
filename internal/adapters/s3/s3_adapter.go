package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Adapter struct {
	client     *s3.S3
	bucketName string
}

func NewS3Adapter(region, bucketName string) (*S3Adapter, error) {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	if err != nil {
		return nil, err
	}
	return &S3Adapter{
		client:     s3.New(sess),
		bucketName: bucketName,
	}, nil
}

// Implement your methods to interact with S3
func (s *S3Adapter) ListObjects() ([]string, error) {
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(s.bucketName),
	}

	result, err := s.client.ListObjectsV2(input)
	if err != nil {
		return nil, err
	}

	var objects []string
	for _, item := range result.Contents {
		objects = append(objects, *item.Key)
	}

	return objects, nil
}
