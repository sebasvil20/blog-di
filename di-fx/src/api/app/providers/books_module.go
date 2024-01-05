package providers

import (
	"github.com/sebasvil20/di-fx/src/api/controllers"
	"github.com/sebasvil20/di-fx/src/api/models"
	"github.com/sebasvil20/di-fx/src/api/repositories"
	"github.com/sebasvil20/di-fx/src/api/services"
	"go.uber.org/fx"
)

var BooksModule = fx.Module("Books",
	fx.Provide(func(db []models.Book) repositories.IBooksRepository {
		return &repositories.BooksRepository{
			DB: db,
		}
	}),

	fx.Provide(func(repo repositories.IBooksRepository) services.IBooksService {
		return &services.BooksService{
			Repo: repo,
		}
	}),

	fx.Provide(func(srv services.IBooksService) *controllers.BooksController {
		return &controllers.BooksController{
			Service: srv,
		}
	}),
)
