package main

import (
    "context"
    "fmt"
    "log"
    "coca-cola-lambda/internal/adapters/secretmanager"
    "coca-cola-lambda/internal/adapters/storage"
    "coca-cola-lambda/internal/config"
    "github.com/aws/aws-lambda-go/lambda"
)

func handleRequest(ctx context.Context) (string, error) {
    secretName := config.GetEnv("SECRET_NAME", "secreto/prueba")
    region := config.GetEnv("AWS_REGION", "us-east-1")
    bucketName := config.GetEnv("BUCKET_NAME", "proyecto-prueba")
    key := "aws.key"

    sm := secretmanager.NewAWSSecretManager(region)
    secret, err := sm.GetSecret(secretName)
    if err != nil {
        return "", fmt.Errorf("failed to get secret: %w", err)
    }

    s3 := storage.NewAWSS3(region)
    content, err := s3.GetObject(bucketName, key)
    if err != nil {
        return "", fmt.Errorf("failed to get object from S3: %w", err)
    }

    log.Printf("Content from S3: %s", content)
    log.Printf("Secret Username: %s", secret.Username)

    return "Operation successful", nil
}

func main() {
    lambda.Start(handleRequest)
}
