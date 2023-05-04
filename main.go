package main

import (
	"github.com/gin-gonic/gin"
	"maraka/auth"
	"maraka/booking"
	"maraka/db"
	"os"
)

func main() {
	//if godotenv.Load(".env") != nil {
	//	log.Fatalf("Error loading .env file")
	//}
	router := gin.Default()

	db.Init(os.Getenv("DB_URL"))

	public := router.Group("auth")
	public.POST("/register", auth.Register)
	public.POST("/login", auth.Login)

	router.GET("/booking", booking.GetBooks)
	router.GET("/booking/:id", booking.GetBookById)
	router.POST("/booking", booking.NewBook)

	router.Run(":8080")
}
