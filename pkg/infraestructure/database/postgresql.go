package database

import (
	"coca-cola-lambda/pkg/domain/models"
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

type PostgresAdapter struct {
	db *sql.DB
}

func NewPostgresAdapter(dsn string) *PostgresAdapter {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	return &PostgresAdapter{
		db: db,
	}
}

func (p *PostgresAdapter) ListClients(ctx context.Context) ([]models.Client, error) {
	rows, err := p.db.QueryContext(ctx, "SELECT id, name FROM clients")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clients []models.Client
	for rows.Next() {
		var client models.Client
		if err := rows.Scan(&client.ID, &client.Name); err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}

	return clients, nil
}
