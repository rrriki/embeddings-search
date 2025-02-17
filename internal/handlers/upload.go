package handlers

import (
	"net/http"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rrriki/embeddings-search/internal/services"
)

func UploadFileHandler(c *gin.Context) {
	file, err := c.FormFile("file")

	if err != nil {
		log.Printf("Error reading file: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file upload"})
		return
	}

	savePath, err := services.SaveUploadedFile(c, file)

	if err != nil {
		log.Printf("Error saving file: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	text, err := services.ExtractTextFromFile(savePath)

	if err != nil {
		log.Printf("Error extracting text: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to extract text from file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "File uploaded successfully",
		"file":    file.Filename,
		"path":    savePath,
		"content": text,
	})
}