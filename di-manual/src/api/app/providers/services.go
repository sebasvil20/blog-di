package providers

import (
	"github.com/sebasvil20/di-manual/src/api/repositories"
	"github.com/sebasvil20/di-manual/src/api/services"
)

func ProvideBooksService(repo *repositories.BooksRepository) *services.BooksService {
	return &services.BooksService{
		Repo: repo,
	}
}
