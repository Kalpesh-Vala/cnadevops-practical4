package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// setupRouter configures and returns a Gin router
func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	return r
}

func main() {
	fmt.Println("Starting microservice...")
	r := setupRouter()
	fmt.Println("Server running on :8081")
	r.Run(":8081")
}
