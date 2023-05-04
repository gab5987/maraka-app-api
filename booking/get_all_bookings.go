package booking

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}
