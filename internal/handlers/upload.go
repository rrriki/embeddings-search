package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rrriki/embeddings-search/internal/services"
	"github.com/rrriki/embeddings-search/internal/storage"
)

func UploadFileHandler(c *gin.Context) {
	file, err := c.FormFile("file")

	if err != nil {
		log.Printf("Error reading file: %e", err)
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
		log.Printf("Error extracting text: %e", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to extract text from file"})
		return
	}

	embeddings, err := services.GenerateEmbeddings(text)

	if err != nil {
		log.Printf("Error generating embeddings: %e", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate embeddings"})	
	}

	err = storage.InsertVector(file.Filename, embeddings, text)

	if err != nil {
		log.Printf("Error inserting vector: %e", err)
	}


	c.JSON(http.StatusOK, gin.H{
		"message": "File uploaded successfully",
		"file":    file.Filename,
		"path":    savePath,
		"content": text,
		"embeddings": embeddings,	
	})
}