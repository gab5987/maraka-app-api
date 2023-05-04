package booking

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewBook(c *gin.Context) {
	var newBook Booking
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}
