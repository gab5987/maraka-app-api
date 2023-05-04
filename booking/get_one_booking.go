package booking

import (
	"github.com/gin-gonic/gin"
	"maraka/auth"
	"net/http"
)

func GetBookById(c *gin.Context) {
	if auth.TokenValid(c) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	id := c.Param("id")

	for _, a := range books {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "booking not found"})
}
