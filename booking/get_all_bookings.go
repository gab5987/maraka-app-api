package booking

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"maraka/auth"
	"maraka/db"
	"net/http"
)

func GetBooks(c *gin.Context) {
	if auth.TokenValid(c) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	paginateOptions := Paginate(c)

	var books []Booking

	cursor, err := db.BookingCollection.Find(db.Ctx, bson.D{}, &paginateOptions)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for cursor.Next(db.Ctx) {
		var elem Booking
		err := cursor.Decode(&elem)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		books = append(books, elem)
	}

	count, _ := db.BookingCollection.CountDocuments(db.Ctx, bson.D{})

	c.IndentedJSON(http.StatusOK, gin.H{
		"data":      books,
		"totalDocs": count,
		"limit":     paginateOptions.Limit,
	})
}
