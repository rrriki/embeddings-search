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