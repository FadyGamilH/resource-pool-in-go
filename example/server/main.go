package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Response struct {
	Message string `json:"message"`
}

func main() {
	router := gin.Default()

	router.POST("/greet", func(c *gin.Context) {
		var request Request
		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		response := Response{
			Message: "Hello " + request.Name + ", you are " + strconv.Itoa(request.Age) + " years old!",
		}

		c.JSON(http.StatusOK, response)
	})

	router.Run(":9080")
}
