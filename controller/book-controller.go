package controller

import (
	"net/http"
	"strconv"

	"hbgaldino/golang-api/entity"
	"hbgaldino/golang-api/repository"

	"github.com/gin-gonic/gin"
)

type BookController interface {
	All(context *gin.Context)
	FindById(context *gin.Context)
	Create(context *gin.Context)
	// Insert(context *gin.Context)
	// Update(context *gin.Context)
	// Delete(context *gin.Context)
}

type bookController struct {
	bookRepository repository.BookRepository
}

func NewBookController(repository repository.BookRepository) BookController {
	return &bookController{
		bookRepository: repository,
	}
}
func (c *bookController) All(context *gin.Context) {
	books := c.bookRepository.AllBook()
	context.JSON(http.StatusOK, books)
}

func (c *bookController) Create(context *gin.Context) {
	var bookRequest entity.Book
	err := context.ShouldBind(&bookRequest)

	if err != nil {
		context.JSON(http.StatusBadRequest, "bad request")
		return
	}

	result := c.bookRepository.InsertBook(bookRequest)
	context.JSON(http.StatusCreated, result)
}

func (c *bookController) FindById(context *gin.Context) {
	id, _ := strconv.ParseUint(context.Param("id"), 10, 64)
	book := c.bookRepository.FindBookById(id)
	context.JSON(http.StatusOK, book)
}
