package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// book represents data about a book

type book struct {
	ID     string `json:"id"`
	Name   string `json:"title"`
	Author string `json:"author"`
	Genre  string `json:"category"`
}

// books slice to seed books data
var books = []book{
	{ID: "1", Name: "How to Be an Antiracist", Author: "Ibram X.Kendi", Genre: "Nonfiction"},
	{ID: "2", Name: "Clean Code: A Handbook of Agile Software Craftmanship", Author: "Robert C.Martin Series", Genre: "Computer Science"},
	{ID: "3", Name: "Geek Love", Author: "Katherine Dunn", Genre: "Fiction"},
}

// getBooks responds with the list of books as JSON
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)

	router.Run("localhost:8080")
}
