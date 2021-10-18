package controller

import (
	"fmt"
	"net/http"

	"hbgaldino/golang-api/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookController interface {
	All(context *gin.Context)
	// FindById(context *gin.Context)
	Create(context *gin.Context)
	// Insert(context *gin.Context)
	// Update(context *gin.Context)
	// Delete(context *gin.Context)
}

type bookController struct {
	connection *gorm.DB
}

func NewBookController(conn *gorm.DB) BookController {
	return &bookController{
		connection: conn,
	}
}
func (c *bookController) All(context *gin.Context) {
	var books []entity.Book
	c.connection.Find(&books)
	context.JSON(http.StatusOK, books)
}

func (c *bookController) Create(context *gin.Context) {
	var bookRequest entity.Book
	err := context.ShouldBind(&bookRequest)

	if err != nil {
		context.JSON(http.StatusBadRequest, "bad request")
		return
	}

	result := c.connection.Create(&bookRequest)

	fmt.Println("Rows Affected", result.RowsAffected)
	context.JSON(http.StatusCreated, "created")
}
