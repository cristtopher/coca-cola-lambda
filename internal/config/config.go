package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AWSRegion    string `envconfig:"AWS_REGION"`
	S3BucketName string `envconfig:"S3_BUCKET_NAME"`
	PostgresDSN  string `envconfig:"POSTGRES_DSN"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
