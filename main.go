package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", handlerUser)
	router.GET("/books/:id", booksHandler)
	router.Run()
}

func handlerUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Eko Kurniadi",
		"bio":  "A Software Engineer",
	})
}

func booksHandler(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
