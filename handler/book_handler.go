package handler

import (
	"net/http"

	"product-catalog-solid/data"
	"product-catalog-solid/service"

	"github.com/gin-gonic/gin"
)

// GetBooks handles GET /books endpoint.
func GetBooks(c *gin.Context) {
	books := service.GetBooks()
	c.IndentedJSON(http.StatusOK, books)
}

// GetBookByID handles GET /books/:id endpoint.
func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	book, err := service.GetBookByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

// PostBook handles POST /books endpoint.
func PostBook(c *gin.Context) {
	var newBook data.Book
	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	service.PostBook(newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}
