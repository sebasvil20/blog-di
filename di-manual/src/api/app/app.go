package app

import (
	"github.com/sebasvil20/di-manual/src/api/app/providers"
)

func Run() {
	mockedDB := providers.ProvideDatabase()
	bookRepo := providers.ProvideBooksRepository(mockedDB)
	bookService := providers.ProvideBooksService(bookRepo)
	bookController := providers.ProvideBooksController(bookService)
	router := providers.ProvideRouter()
	providers.RegisterRoutes(router, bookController)
	router.Run(":8080")
}
