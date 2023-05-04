package auth

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"maraka/db"
	"net/http"
)

type RegisterInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

func saveUser(newUser RegisterInput) (r *mongo.InsertOneResult, err error) {
	doc, err := db.ToDoc(newUser)
	return db.UserCollection.InsertOne(db.Ctx, doc)
}

func Register(c *gin.Context) {
	var newUser RegisterInput
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := saveUser(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, result)
}
