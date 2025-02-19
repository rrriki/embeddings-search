package services

import (
	"fmt"

	"github.com/rrriki/embeddings-search/internal/storage"
)

type SearchRequest struct {
	Query string `json:"query"`
}

func SearchDocuments(query string) ([]storage.SearchResult, error) {
	queryEmbeddings, err := GenerateEmbeddings(query)

	if err != nil {
		fmt.Printf("Error generating embeddings from query: %e", err)
		return nil, err
	}

	results, err := storage.SearchVectors(queryEmbeddings)

	if err != nil {
		fmt.Printf("Error searching vectors: %e", err)
		return nil, err
	}

	return results, nil

}
