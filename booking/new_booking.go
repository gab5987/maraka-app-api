package booking

import (
	"github.com/gin-gonic/gin"
	"maraka/auth"
	"maraka/db"
	"net/http"
)

func NewBook(c *gin.Context) {
	if auth.TokenValid(c) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var newBook Booking
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	bk, err := db.BookingCollection.InsertOne(db.Ctx, newBook)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, bk)
}
