package storage

import (
	"context"
	"database/sql"
)

type UserStorage struct {
	db *sql.DB
}

func (s *UserStorage) Create(ctx context.Context) error {
	return nil
}
