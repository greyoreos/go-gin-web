package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func postBooks(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func getBookByid(c *gin.Context) {
	ID := c.Param("ID")

	for _, b := range books {
		if b.ID == ID {
			c.IndentedJSON(http.StatusOK, b)
			return

		}

	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})

}

type book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

var books = []book{
	{ID: "1", Title: "The Comfort Book", Author: "Matt Haig", Price: 10.90},
	{ID: "2", Title: "The Alchemist", Author: "Paulo Coelho", Price: 7.59},
	{ID: "3", Title: "A Brief History of Time", Author: "Stephen Hawking", Price: 39.99},
}

func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:ID", getBookByid)
	router.POST("/books", postBooks)

	router.Run("localhost:8080")

}
