package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UsersStorage struct {
	db *pgxpool.Pool
}

func (s *UsersStorage) Create(ctx context.Context) error {
	return nil
}
