package main

import (
	"github.com/gin-gonic/gin"
	"maraka/booking"
)

func main() {
	router := gin.Default()
	router.GET("/booking", booking.GetBooks)
	router.GET("/booking/:id", booking.GetBookById)
	router.POST("/booking", booking.NewBook)

	router.Run(":8080")
}
