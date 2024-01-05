package providers

import "github.com/sebasvil20/di-fx/src/api/models"

func ProvideDatabase() []models.Book {
	return []models.Book{
		{ID: 1, Title: "El señor de los anillos", Author: "J.R.R. Tolkien"},
		{ID: 2, Title: "Cien años de soledad", Author: "Gabriel García Márquez"},
	}
}
