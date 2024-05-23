package main

import (
    "context"
    "log"
    "github.com/aws/aws-lambda-go/lambda"
    "coca-cola-lambda/pkg/infraestructure/aws"
    "coca-cola-lambda/pkg/infraestructure/database"
    "coca-cola-lambda/pkg/domain/services"
    "coca-cola-lambda/pkg/config"
)

func main() {
    // Cargar configuración
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Error loading config: %v", err)
    }

    // Inicializar adaptadores
    s3Adapter := aws.NewS3Adapter(cfg.S3Bucket)
    secretManagerAdapter := aws.NewSecretManagerAdapter()
    dbAdapter := database.NewPostgresAdapter(cfg.DatabaseDSN)

    // Inicializar servicio
    clientService := services.NewClientService(s3Adapter, secretManagerAdapter, dbAdapter)

    // Inicializar Lambda handler
    handler := func(ctx context.Context) error {
        return clientService.HandleRequest(ctx)
    }

    lambda.Start(handler)
}
