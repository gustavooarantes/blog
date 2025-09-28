package storage

import (
	"context"
	"database/sql"
)

type PostStorage struct {
	db *sql.DB
}

func (s *PostStorage) Create(ctx context.Context) error {
	return nil
}
