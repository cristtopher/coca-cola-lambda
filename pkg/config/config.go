package config

import (
    "github.com/kelseyhightower/envconfig"
)

type Config struct {
    S3Bucket      string `envconfig:"S3_BUCKET"`
    SecretManager string `envconfig:"SECRET_MANAGER"`
    DatabaseDSN   string `envconfig:"DATABASE_DSN"`
}

func LoadConfig() (*Config, error) {
    var cfg Config
    err := envconfig.Process("", &cfg)
    if err != nil {
        return nil, err
    }
    return &cfg, nil
}
