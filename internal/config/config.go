package config

import (
	"log"
	"os"
	"sync"
)

type Config struct {
	OpenAiApiUrl string
	OpenAiApiKey string	

	PineconeAPIKey string
	PineconeIndexHost string
}

var instance *Config
var once sync.Once

func LoadConfig() *Config {

	once.Do(func(){	
		openApiUrl := os.Getenv("OPENAI_API_URL")
		openApiKey := os.Getenv("OPENAI_API_KEY")

		if openApiUrl == "" || openApiKey == "" {
			log.Fatalf("OPENAI env vars missing")
		}


		pineconeApiKey := os.Getenv("PINECONE_API_KEY")
		pineconeIndexHost := os.Getenv("PINECONE_INDEX_HOST")

		if pineconeApiKey == "" || pineconeIndexHost == "" {
			log.Fatalf("Pinecone env vars missing")
		}

		instance = &Config{
			OpenAiApiUrl: openApiUrl,
			OpenAiApiKey: openApiKey,

			PineconeAPIKey: pineconeApiKey,
			PineconeIndexHost: pineconeIndexHost,
		}
	})

	return instance	
}
