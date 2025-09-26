package repositories

import (
	"context"

	"github.com/gustavooarantes/blog/internal/repositories/entities"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UsersStorage struct {
	db *pgxpool.Pool
}

func (s *UsersStorage) Create(ctx context.Context, user *entities.User) error {
	query := `
		INSERT INTO users (username, email, password)
		VALUES ($1, $2, $3)
		RETURNING id, created_at
	`

	err := s.db.QueryRow(
		ctx,
		query,
		user.Username,
		user.Email,
		user.Password,
	).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}
