package booking

import (
	"github.com/gin-gonic/gin"
	"maraka/auth"
	"net/http"
)

func GetBooks(c *gin.Context) {
	if auth.TokenValid(c) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	c.IndentedJSON(http.StatusOK, books)
}
