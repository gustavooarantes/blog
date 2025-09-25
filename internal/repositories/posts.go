package repositories

import (
	"context"

	"github.com/gustavooarantes/blog/internal/repositories/entities"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostsStorage struct {
	db *pgxpool.Pool
}

func (s *PostsStorage) Create(ctx context.Context, post *entities.Post) error {
	query := `
		INSERT INTO posts (content, title, user_id, tags)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`

	err := s.db.QueryRow(
		ctx,
		query,
		post.Content,
		post.Title,
		post.UserID,
		post.Tags,
	).Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}
