package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sebasvil20/di-fx/src/api/app/providers"
	"go.uber.org/fx"
)

func Run() {
	fx.New(
		fx.Provide(providers.ProvideDatabase),
		providers.BooksModule,
		fx.Provide(providers.ProvideRouter),
		fx.Invoke(providers.RegisterRoutes),
		fx.Invoke(serve),
	).Run()
}

func serve(lifecycle fx.Lifecycle, router *gin.Engine) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go router.Run()
			return nil
		},
	})
}
