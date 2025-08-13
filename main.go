package main

import (
	"fmt"
	"net/http"

	"testing"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World")
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

	r.Run(":8081")
}

// Intentionally failing test for CI
func TestFail(t *testing.T) {
	t.Fatal("Intentional failure for CI workflow")
}
