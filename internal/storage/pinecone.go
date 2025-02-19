package storage

import (
	"context"
	"fmt"
	"log"

	"github.com/pinecone-io/go-pinecone/v3/pinecone"
	"github.com/rrriki/embeddings-search/internal/config"
	"google.golang.org/protobuf/types/known/structpb"
)
var pineconeClient *pinecone.Client

type SearchResult struct {
	Id string `json:"id"`	
	Score float32 `json:"score"`
	Text string `json:"text"`
}

func InitializePinecone() {
	cfg := config.LoadConfig()

	clientParams := pinecone.NewClientParams{
		ApiKey: cfg.PineconeAPIKey,
	}
	client, err := pinecone.NewClient(clientParams)

	if err != nil {
		log.Fatalf("Failed to create Pinecone client: %v", err)		
	}

	pineconeClient = client

	fmt.Println("Pinecone initialized successfully")
}

func InsertVector(id string, embeddings []float32, sourceText string) error {
	if pineconeClient == nil {
		log.Fatal("Pinecone client is not initialized")
	}

	cfg := config.LoadConfig()
	context := context.Background()

	metadata, _ := structpb.NewStruct(map[string]interface{}{
		"source_text": sourceText,
	})

	vectors := []*pinecone.Vector{
		{
			Id: id,
			Values: &embeddings,
			Metadata: metadata,
		},
	}
	
	indexConnParams := pinecone.NewIndexConnParams{ Host: cfg.PineconeIndexHost }
	index, err := pineconeClient.Index(indexConnParams)

	if err != nil { 
		return fmt.Errorf("failed to describe pinecone index: %v", err)
	}

	index.UpsertVectors(context, vectors)

	return nil
}

func SearchVectors(vectors []float32) ([]SearchResult, error){

	if pineconeClient == nil {
		log.Fatal("Pinecone client is not initialized")
	}

	cfg := config.LoadConfig()
	context := context.Background()

	indexConnParams := pinecone.NewIndexConnParams{ Host: cfg.PineconeIndexHost }
	index, err := pineconeClient.Index(indexConnParams)

	if err != nil {
		return nil, fmt.Errorf("failed to describe pinecone index: %v", err)
	}

	search := pinecone.QueryByVectorValuesRequest{
		Vector: vectors,
		TopK: 5,
		IncludeMetadata: true,
	}

	result, err := index.QueryByVectorValues(context, &search)

	if err != nil {
		return nil, fmt.Errorf("failed to query pinecone index: %e", err)
	}

	results := make([]SearchResult, 0)

	for _, result := range result.Matches {
		results = append(results, SearchResult{
			Id: result.Vector.Id,
			Score: result.Score,
			Text: result.Vector.Metadata.Fields["source_text"].GetStringValue(),
		})
	}

	return results, nil

}