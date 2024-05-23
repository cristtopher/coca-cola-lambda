package services

import (
    "context"
    "log"
    "coca-cola-lambda/pkg/domain/ports"
)

type ClientService struct {
    s3Adapter        ports.S3Port
    secretManagerAdapter ports.SecretManagerPort
    clientRepository ports.ClientRepository
}

func NewClientService(s3Adapter ports.S3Port, secretManagerAdapter ports.SecretManagerPort, clientRepository ports.ClientRepository) *ClientService {
    return &ClientService{
        s3Adapter:        s3Adapter,
        secretManagerAdapter: secretManagerAdapter,
        clientRepository: clientRepository,
    }
}

func (c *ClientService) HandleRequest(ctx context.Context) error {
    // Obtener objetos de S3
    objects, err := c.s3Adapter.ListObjects(ctx)
    if err != nil {
        return err
    }
    log.Println("Objects from S3:", objects)

    // Obtener secretos de Secret Manager
    secrets, err := c.secretManagerAdapter.GetSecret(ctx, "mySecret")
    if err != nil {
        return err
    }
    log.Println("Secrets from Secret Manager:", secrets)

    // Obtener clientes de PostgreSQL
    clients, err := c.clientRepository.ListClients(ctx)
    if err != nil {
        return err
    }
    log.Println("Clients from PostgreSQL:", clients)

    return nil
}
