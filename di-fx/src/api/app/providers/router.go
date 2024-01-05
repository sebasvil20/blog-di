package providers

import (
	"github.com/gin-gonic/gin"
	"github.com/sebasvil20/di-fx/src/api/controllers"
)

func ProvideRouter() *gin.Engine {
	return gin.Default()
}

func RegisterRoutes(r *gin.Engine, booksController *controllers.BooksController) {
	r.GET("/books", booksController.GetAll)
}
