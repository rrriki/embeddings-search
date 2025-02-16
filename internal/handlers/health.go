package handlers

import (

	"github.com/gin-gonic/gin"
)

func HealthHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
		"message": "Embeddings Search API is running",
	})
}