package providers

import (
	"github.com/sebasvil20/di-manual/src/api/controllers"
	"github.com/sebasvil20/di-manual/src/api/services"
)

func ProvideBooksController(srv *services.BooksService) *controllers.BooksController {
	return &controllers.BooksController{
		Service: srv,
	}
}
