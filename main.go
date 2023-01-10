package main

import (
	"github.com/alirezamor/assesment/flights"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/track", flights.GetStartAndEnd)
	router.Run("localhost:8080")
}
