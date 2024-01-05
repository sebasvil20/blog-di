package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sebasvil20/di-fx/src/api/services"
)

type IBooksController interface {
	GetAll(c *gin.Context)
}

type BooksController struct {
	Service services.IBooksService
}

func (c *BooksController) GetAll(ctx *gin.Context) {
	books, _ := c.Service.GetAll(ctx)
	ctx.JSON(200, gin.H{
		"books": books,
	})
}
