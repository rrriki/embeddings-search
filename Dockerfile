FROM golang:1.24-alpine

WORKDIR /app

RUN apk add --no-cache docker-cli

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd/api

RUN go build -o /app/embeddings-search

EXPOSE 8080

CMD ["/app/embeddings-search"]
