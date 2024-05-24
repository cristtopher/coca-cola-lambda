package postgres

import (
	"coca-cola-lambda/internal/domain"

	"github.com/jmoiron/sqlx"
)

type PostgresClientRepository struct {
	db *sqlx.DB
}

func NewPostgresClientRepository(db *sqlx.DB) *PostgresClientRepository {
	return &PostgresClientRepository{db: db}
}

func (r *PostgresClientRepository) GetAllClients() ([]domain.Client, error) {
	var clients []domain.Client
	err := r.db.Select(&clients, "SELECT * FROM clients")
	if err != nil {
		return nil, err
	}
	return clients, nil
}

func (r *PostgresClientRepository) CreateClient(client domain.Client) error {
	_, err := r.db.NamedExec(`INSERT INTO clients (name, email) VALUES (:name, :email)`, client)
	return err
}
