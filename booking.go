package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Booking struct {
	ID           string `json:"id"`
	CustomerName string `json:"customerName"`
	Room         uint8  `json:"room"`
	StartDate    string `json:"startDate"`
	DueDate      string `json:"dueDate"`
	BookingType  string `json:"bookingType"`
	Contact      string `json:"contact"`
}

var books []Booking

func NewBook(c *gin.Context) {
	var newBook Booking
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

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
