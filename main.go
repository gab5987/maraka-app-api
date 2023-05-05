package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"maraka/auth"
	"maraka/booking"
	"maraka/db"
	"os"
)

func main() {
	if os.Getenv("GIN_MODE") != "release" {
		if godotenv.Load(".env") != nil {
			log.Fatalf("Error loading .env file")
		}
	}
	router := gin.Default()

	db.Init(os.Getenv("DB_URL"))

	public := router.Group("auth")
	public.POST("/register", auth.Register)
	public.POST("/login", auth.Login)

	private := router.Group("booking")

	private.POST("", booking.NewBook)
	private.GET("", booking.GetBooks)
	private.GET("/:id", booking.GetBookById)

	router.Run(":8080")
}
