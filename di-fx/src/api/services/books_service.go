package services

import (
	"context"
	"github.com/sebasvil20/di-fx/src/api/models"
	"github.com/sebasvil20/di-fx/src/api/repositories"
)

type IBooksService interface {
	GetAll(ctx context.Context) ([]models.Book, error)
}

type BooksService struct {
	Repo repositories.IBooksRepository
}

func (s *BooksService) GetAll(ctx context.Context) ([]models.Book, error) {
	return s.Repo.GetAll(ctx)
}
