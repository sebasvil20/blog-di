package repositories

import (
	"context"
	"github.com/sebasvil20/di-manual/src/api/models"
)

type IBooksRepository interface {
	GetAll(ctx context.Context) ([]models.Book, error)
}

type BooksRepository struct {
	DB []models.Book
}

func (r *BooksRepository) GetAll(ctx context.Context) ([]models.Book, error) {
	return r.DB, nil
}
