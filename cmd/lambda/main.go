package main

import (
	"log"

	"coca-cola-lambda/internal/adapters/postgres"
	"coca-cola-lambda/internal/adapters/s3"
	"coca-cola-lambda/internal/adapters/secretmanager"
	"coca-cola-lambda/internal/config"
	"coca-cola-lambda/internal/domain"
	"coca-cola-lambda/internal/handlers"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := sqlx.Connect("postgres", cfg.PostgresDSN)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	defer db.Close()

	s3Adapter, err := s3.NewS3Adapter(cfg.AWSRegion, cfg.S3BucketName)
	if err != nil {
		log.Fatalf("Failed to initialize S3 adapter: %v", err)
	}

	secretManagerAdapter, err := secretmanager.NewSecretManagerAdapter(cfg.AWSRegion)
	if err != nil {
		log.Fatalf("Failed to initialize Secret Manager adapter: %v", err)
	}

	clientRepository := postgres.NewPostgresClientRepository(db)

	clientService := domain.NewClientService(clientRepository)

	httpHandler := handlers.NewHTTPHandler(clientService, s3Adapter, secretManagerAdapter)

	lambda.Start(httpHandler.HandleRequest)
}
