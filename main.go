package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"maraka/auth"
	"maraka/booking"
	"maraka/db"
	"os"
	"time"
)

func main() {
	if os.Getenv("GIN_MODE") != "release" {
		if godotenv.Load(".env") != nil {
			log.Fatalf("Error loading .env file")
		}
	}
	router := gin.Default()

	headers := []string{
		"Content-Type",
		"Content-Length",
		"Accept-Encoding",
		"X-CSRF-Token",
		"Authorization",
		"accept",
		"origin",
		"Cache-Control",
		"X-Requested-With",
	}

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     headers,
		ExposeHeaders:    headers,
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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
