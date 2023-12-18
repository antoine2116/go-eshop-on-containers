FROM golang:1.20-alpine AS builder

# Build Delve
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# Build Air
RUN go install github.com/cosmtrek/air@latest

WORKDIR /app/internal/services/catalog_service

COPY go.mod ./
RUN go mod download

CMD ["air", "-c", "air.toml", "-d"]
