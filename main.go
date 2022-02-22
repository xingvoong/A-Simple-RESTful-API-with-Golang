package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// book represents data about a book

type book struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
}

// books slice to seed books data
var books = []book{
	{ID: "1", Name: "How to Be an Antiracist", Author: "Ibram X.Kendi", Genre: "Nonfiction"},
	{ID: "2", Name: "Clean Code: A Handbook of Agile Software Craftmanship", Author: "Robert C.Martin Series", Genre: "Computer Science"},
	{ID: "3", Name: "Geek Love", Author: "Katherine Dunn", Genre: "Fiction"},
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", postBooks)

	router.Run("localhost:8080")
}

// getBooks responds with the list of books as JSON
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// postBooks adds an album from JSOn received in the request body
func postBooks(c *gin.Context) {
	var newBook book

	// Call BindJson to bind the received JSON to newBook.
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	// Add the new book to the slice
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}
