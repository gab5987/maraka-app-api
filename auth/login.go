package auth

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"maraka/db"
	"net/http"
)

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	ID    primitive.ObjectID `bson:"_id"`
	Name  string
	Email string
	Token string
}

func Login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var result User
	err := db.UserCollection.FindOne(db.Ctx, bson.D{
		{"email", input.Email},
		{"password", input.Password},
	}).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusUnauthorized, "User does not exist")
			return
		}
		return
	}

	tk, err := GenerateToken(result.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error generating token")
		return
	}
	result.Token = tk

	c.JSON(http.StatusOK, result)
}
