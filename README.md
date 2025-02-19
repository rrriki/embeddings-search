# embeddings-search
A Go API for semantic search on documents using OpenAI embeddings and PineconeDB. 

Upload PDFs & text files, extract content, generate embeddings, and perform vector-based similarity searches.

## Getting Started

**1.- Set up the environment**

Create an `.env` file in the root directory with

```sh
OPENAI_API_URL=https://api.openai.com/v1/embeddings
OPENAI_API_KEY=your_openai_api_key
PINECONE_API_KEY=your_pinecone_api_key
PINECONE_INDEX_HOST=your_pinecone_index_host_url
```

**2.- Start the service**

Run the API and dependencies ([PDFBox](https://pdfbox.apache.org/))

```sh
docker-compose up --build
```



## API Endpoints

- `POST /upload`

    Uploads a `.pdf` or `.txt` file, extracts text, generates embeddings using OpenAI and stores it in PineconeDB for fast seach

    ```sh
    curl -X POST "http://localhost:8080/upload" -F "file=@test.pdf"
    ```

- `POST /search`

    Pass a query string to search stored documents using embedding cosine similarity.

    ```sh
    curl -X POST "http://localhost:8080/search" -H "Content-Type: application/json" -d '{"query": "breach of contract"}'

    ```



## License

[GPL 3.0](https://choosealicense.com/licenses/gpl-3.0/)

