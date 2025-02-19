package main

import (
	"fmt"
	"log"
	"github.com/rrriki/embeddings-search/internal/storage"
	"github.com/rrriki/embeddings-search/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	storage.InitializePinecone()


	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.GET("/health", handlers.HealthHandler)

	router.POST("upload", handlers.UploadFileHandler)

	router.POST("search", handlers.SearchHandler)
	
	port := "8080"

	fmt.Println("Embeddings Search API is running on port: ", port)

	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
