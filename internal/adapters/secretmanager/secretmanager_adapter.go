package secretmanager

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

type SecretManagerAdapter struct {
	client *secretsmanager.SecretsManager
}

func NewSecretManagerAdapter(region string) (*SecretManagerAdapter, error) {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	if err != nil {
		return nil, err
	}
	return &SecretManagerAdapter{
		client: secretsmanager.New(sess),
	}, nil
}

func (s *SecretManagerAdapter) GetSecret(secretName string) (map[string]string, error) {
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	}

	result, err := s.client.GetSecretValue(input)
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
