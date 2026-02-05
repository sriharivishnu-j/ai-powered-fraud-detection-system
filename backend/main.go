package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "UP"})
	})

	r.POST("/fraud-detection", handleFraudDetection)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}

func handleFraudDetection(c *gin.Context) {
	// Logic to call AI service and return result
	c.JSON(http.StatusOK, gin.H{"message": "Fraud detection processed"})
}