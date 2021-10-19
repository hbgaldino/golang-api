package main

import (
	"hbgaldino/golang-api/conf"
	"hbgaldino/golang-api/controller"
	"hbgaldino/golang-api/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// dependency injection and object initialization
var (
	db             *gorm.DB                  = conf.SetupDatabase()
	bookRepository repository.BookRepository = repository.NewBookRepository(db)
	bookController controller.BookController = controller.NewBookController(bookRepository)
)

func main() {
	router := gin.Default()
	router.GET("/book", bookController.All)
	router.GET("/book/:id", bookController.FindById)
	router.POST("/book", bookController.Create)
	router.DELETE("/book/:id", bookController.Delete)
	router.PUT("/book/:id", bookController.Update)
	router.Run()
}
