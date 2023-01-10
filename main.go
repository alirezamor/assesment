package main

import (
	"github.com/alirezamor/assesment/flights"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/calculate/", flights.GetStartAndEnd)
	router.Run("localhost:8080")
}
