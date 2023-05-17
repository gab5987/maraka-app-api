package booking

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"maraka/auth"
	"maraka/db"
	"net/http"
	"strconv"
)

func GetBooks(c *gin.Context) {
	if auth.TokenValid(c) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	paginateOptions := Paginate(c)

	var books []Booking

	var query = bson.D{}
	_room, hasKey := c.GetQuery("room")

	if hasKey {
		room, _ := strconv.ParseInt(_room, 10, 64)
		query = bson.D{{"room", room}}
	}

	cursor, err := db.BookingCollection.Find(db.Ctx, query, &paginateOptions)

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

	count, _ := db.BookingCollection.CountDocuments(db.Ctx, query)

	c.IndentedJSON(http.StatusOK, gin.H{
		"data":      books,
		"totalDocs": count,
		"limit":     paginateOptions.Limit,
	})
}
