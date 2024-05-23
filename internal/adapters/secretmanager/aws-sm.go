package secretmanager

import (
    "encoding/json"
    "fmt"
    "coca-cola-lambda/internal/core"
    "coca-cola-lambda/internal/ports"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/secretsmanager"
)

type AWSSecretManager struct {
    client *secretsmanager.SecretsManager
}

func NewAWSSecretManager(region string) ports.SecretManager {
    sess := session.Must(session.NewSession(&aws.Config{
        Region: aws.String(region),
    }))
    return &AWSSecretManager{
        client: secretsmanager.New(sess),
    }
}

func (sm *AWSSecretManager) GetSecret(name string) (*core.Secret, error) {
    result, err := sm.client.GetSecretValue(&secretsmanager.GetSecretValueInput{
        SecretId: aws.String(name),
    })
    if err != nil {
        return nil, fmt.Errorf("failed to retrieve secret: %w", err)
    }

    var secret core.Secret
    err = json.Unmarshal([]byte(*result.SecretString), &secret)
    if err != nil {
        return nil, fmt.Errorf("failed to unmarshal secret: %w", err)
    }

    return &secret, nil
}
