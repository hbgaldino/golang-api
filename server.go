package main

import (
	"hbgaldino/golang-api/conf"
	"hbgaldino/golang-api/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB = conf.SetupDatabase()
var bookController = controller.NewBookController(db)

func main() {

	router := gin.Default()
	router.GET("/book", bookController.All)
	router.POST("book", bookController.Create)
	router.Run()
}
