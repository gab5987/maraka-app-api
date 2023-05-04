package booking

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetBookById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range books {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "booking not found"})
}
