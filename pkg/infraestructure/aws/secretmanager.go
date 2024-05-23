package aws

import (
    "context"
    "encoding/json"
    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type SecretManagerAdapter struct {
    client *secretsmanager.Client
}

func NewSecretManagerAdapter() *SecretManagerAdapter {
    cfg, err := config.LoadDefaultConfig(context.TODO())
    if err != nil {
        panic("unable to load SDK config, " + err.Error())
    }
    client := secretsmanager.NewFromConfig(cfg)

    return &SecretManagerAdapter{
        client: client,
    }
}

func (s *SecretManagerAdapter) GetSecret(ctx context.Context, secretName string) (map[string]string, error) {
    input := &secretsmanager.GetSecretValueInput{
        SecretId: aws.String(secretName),
    }

    result, err := s.client.GetSecretValue(ctx, input)
    if err != nil {
        return nil, err
    }

    var secret map[string]string
    err = json.Unmarshal([]byte(*result.SecretString), &secret)
    if err != nil {
        return nil, err
    }

    return secret, nil
}
