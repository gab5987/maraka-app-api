package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/booking", GetBooks)
	router.GET("/booking/:id", GetBookById)
	router.POST("/booking", NewBook)

	router.Run(":8080")
}
