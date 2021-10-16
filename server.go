package main

import (
	"hbgaldino/golang-api/controller"

	"github.com/gin-gonic/gin"
)

var bookController = controller.NewBookController()

func main() {
	router := gin.Default()
	router.GET("/book", bookController.All)
	router.POST("book", bookController.Create)
	router.Run()
}
