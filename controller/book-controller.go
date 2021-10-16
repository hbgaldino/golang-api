package controller

import (
	"net/http"

	"hbgaldino/golang-api/entity"

	"github.com/gin-gonic/gin"
)

var books = []entity.Book{
	{ID: 1, Title: "Name", Description: "Description"},
}

type BookController interface {
	All(context *gin.Context)
	// FindById(context *gin.Context)
	Create(context *gin.Context)
	// Insert(context *gin.Context)
	// Update(context *gin.Context)
	// Delete(context *gin.Context)
}

type bookController struct{}

func NewBookController() BookController {
	return &bookController{}
}
func (c *bookController) All(context *gin.Context) {
	context.JSON(http.StatusOK, books)
}

func (c *bookController) Create(context *gin.Context) {
	var bookRequest entity.Book
	err := context.ShouldBind(&bookRequest)

	if err != nil {
		context.JSON(http.StatusBadRequest, "bad request")
		return
	}

	books = append(books, bookRequest)
	context.JSON(http.StatusCreated, "created")
}
