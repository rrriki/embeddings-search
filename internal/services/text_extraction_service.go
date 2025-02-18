package services

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
	"bytes"
	"os"

	// "github.com/otiai10/gosseract/v2"
)


func ExtractTextFromFile(path string) (string, error) {
	ext := strings.ToLower(filepath.Ext(path)) 

	switch ext {
		case ".txt":
			return extractTextFromTxt(path)
		case ".pdf": 
			return extractTextFromPDF(path)	
		default:
			return "", fmt.Errorf("invalid file format: %s", ext)
	}
}

func extractTextFromTxt(path string) (string, error) {
	content, err := os.ReadFile(path)

	if err != nil {	
		return "", fmt.Errorf("failed to read text file: %s", err)
	}	

	return string(content), nil	
}

func extractTextFromPDF(path string) (string, error) {
	filename := filepath.Base(path[:len(path)- len(filepath.Ext(path))])
	uploadsDir := filepath.Join("/", "app", "uploads")
	
	inputFilePath := filepath.Join(uploadsDir, filename + ".pdf")
	outputFilePath := filepath.Join(uploadsDir, filename + ".txt")

	cmd := exec.Command("docker", "exec", "pdfbox", "java", "-jar", "pdfbox-app-1.8.11.jar", "ExtractText", inputFilePath, outputFilePath)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()

	if err != nil {
		fmt.Printf("failed pdf output %v", out.String())
		return "", fmt.Errorf("failed to extract text from pdf: %s", err)
	}

	readCmd := exec.Command("cat", outputFilePath)

	readCmdOutput, err := readCmd.Output()
	
	if err != nil {
		return "", fmt.Errorf("failed to read extracted text: %s", err)
	}

	return string(readCmdOutput), nil

}	



