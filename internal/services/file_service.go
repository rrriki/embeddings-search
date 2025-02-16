package services

import (
	"fmt"
	"log"
	"mime/multipart"
	"path/filepath"
	"github.com/gin-gonic/gin"
)

var AllowedFileTypes = map[string]bool{
	".pdf": true,
	".txt": true,
}


func SaveUploadedFile(c *gin.Context, file *multipart.FileHeader) (string, error) {
	ext := filepath.Ext(file.Filename)

	if !AllowedFileTypes[ext] {
		return "", fmt.Errorf("invalid file format: %s", ext)
	}

	savePath := fmt.Sprintf("/app/uploads/%s", file.Filename)

	if err := c.SaveUploadedFile(file, savePath); err != nil {
		log.Printf("Failed to save file %s: %v", savePath, err)
		return "", err
	}

	return savePath, nil
}
