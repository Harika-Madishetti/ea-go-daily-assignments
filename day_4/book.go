package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id    int
	Title string
	Price float64
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/books", getBooks)
	return router
}

func getBooks(c *gin.Context) {
	book := Book{1, "Learn Go", 100.6}
	c.JSON(http.StatusOK, book)
}

func main() {
	router := setupRouter()
	router.Run(":8080")
}
