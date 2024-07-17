package main

import (
	"product-catalog-solid/product-catalog-solid/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/books", handler.GetBooks)
	router.GET("/books/:id", handler.GetBookByID)
	router.POST("/books", handler.PostBook)
	router.Run("0.0.0.0:8081")
}
