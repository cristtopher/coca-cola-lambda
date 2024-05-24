package ports

import (
	"coca-cola-lambda/pkg/domain/models"
	"context"
)

type ClientRepository interface {
	ListClients(ctx context.Context) ([]models.Client, error)
	CreateClient(ctx context.Context, client models.Client) error
}
