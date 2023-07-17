package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type example struct {
	Name        string `json:"name"`
	Description string `json:"Description"`
}

var examples = []example{
	{Name: "teste1", Description: "Description"},
}

func GetExamples(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, examples)
}

func main() {
	router := gin.Default()
	router.GET("/books", GetExamples)
	router.Run("localhost:8000")
}
