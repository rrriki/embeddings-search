package services

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/rrriki/embeddings-search/internal/config"
)

type OpenAiEmbeddingsRequest struct {
	Model string `json:"model"`
	Input string `json:"input"`
}

type OpenAiEmbeddingsResponse struct {
	
	Data[]struct {
		Embedding[]float32 `json:"embedding"`
	} `json:"data"`	

	Usage struct {
		PromptTokens int `json:"prompt_tokens"`
		TotalTokens int `json:"total_tokens"`
	} `json:"usage"`	
}

func GenerateEmbeddings(text string) ([]float32, error) {

	cfg := config.LoadConfig()

	client := resty.New()
	defer client.Clone()

	payload := OpenAiEmbeddingsRequest{
		Model: "text-embedding-ada-002",
		Input: text,
	}

	jsonPayload, _ := json.Marshal(payload)

	response, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", cfg.OpenAiApiKey)).
		SetBody(jsonPayload).
		Post(cfg.OpenAiApiUrl)

	if err != nil {
		return nil, fmt.Errorf("error fetching embeddings: %v", err)
	}

	var result OpenAiEmbeddingsResponse 
	
	if err := json.Unmarshal(response.Body(), &result); err != nil {
		return nil, fmt.Errorf("error parsing response: %v", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("no embeddings returned from Open AI")
	}
	fmt.Printf("Fetched embeddings for text %s, tokens used: %d", text, result.Usage.TotalTokens)

	return result.Data[0].Embedding, nil

}
