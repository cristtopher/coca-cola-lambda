package aws

import (
    "context"
    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Adapter struct {
    client   *s3.Client
    bucket   string
}

func NewS3Adapter(bucket string) *S3Adapter {
    cfg, err := config.LoadDefaultConfig(context.TODO())
    if err != nil {
        panic("unable to load SDK config, " + err.Error())
    }
    client := s3.NewFromConfig(cfg)

    return &S3Adapter{
        client: client,
        bucket: bucket,
    }
}

func (s *S3Adapter) ListObjects(ctx context.Context) ([]string, error) {
    input := &s3.ListObjectsV2Input{
        Bucket: aws.String(s.bucket),
    }

    result, err := s.client.ListObjectsV2(ctx, input)
    if err != nil {
        return nil, err
    }

    var objects []string
    for _, item := range result.Contents {
        objects = append(objects, *item.Key)
    }

    return objects, nil
}
