package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rrriki/embeddings-search/internal/services"
)


func SearchHandler(c *gin.Context) {

	var req services.SearchRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Error binding request: %e", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	results, err := services.SearchDocuments(req.Query)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search documents"})
		return
	}

	c.JSON(http.StatusOK, results)
 
}