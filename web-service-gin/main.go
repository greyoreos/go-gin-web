package main 

import ( 
	"net/http"

	"github.com/gin-gonic/gin"
)

type books { 
	ID string `json:"id"`
	Title string `json:"title"` 
	Author string `json:"author"`
	Price float64 `json:"price"` 
} 

var books = []book{ 
	{ID: "1", Title: "The Comfort Book", Author: "Matt Haig", Price: 10.90},
    {ID: "2", Title: "The Alchemist", Author: "Paulo Coelho", Price: 7.59},
    {ID: "3", Title: "A Brief History of Time", Author: "Stephen Hawking", Price: 39.99},

} 

func main(){ 
	router := gin.Default()
	router.GET("/books", getBooks)
	router.Run("localhost:8080")

}


func getBooks(c *gin.Context){ 
	c.JSON(http.StatusOK, books)
}