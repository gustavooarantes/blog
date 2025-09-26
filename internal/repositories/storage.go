// Package repositories is responsible for having everything that relates to database
// implementations, transactions, data persistency etc.
package repositories

import (
	"context"

	"github.com/gustavooarantes/blog/internal/repositories/entities"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	Posts interface {
		Create(context.Context, *entities.Post) error
	}
	Users interface {
		Create(context.Context, *entities.User) error
	}
}

func NewStorage(db *pgxpool.Pool) Storage {
	return Storage{
		Posts: &PostsStorage{db: db},
		Users: &UsersStorage{db: db},
	}
}
