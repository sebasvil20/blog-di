package providers

import (
	"github.com/sebasvil20/di-manual/src/api/models"
	"github.com/sebasvil20/di-manual/src/api/repositories"
)

func ProvideDatabase() []models.Book {
	return []models.Book{
		{ID: 1, Title: "El señor de los anillos", Author: "J.R.R. Tolkien"},
		{ID: 2, Title: "Cien años de soledad", Author: "Gabriel García Márquez"},
	}
}

func ProvideBooksRepository(db []models.Book) *repositories.BooksRepository {
	return &repositories.BooksRepository{
		DB: db,
	}
}
