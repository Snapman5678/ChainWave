# backend/go.dockerfile
FROM golang:1.21-alpine

WORKDIR /app

# Install Air for hot reloading
RUN go install github.com/cosmtrek/air@v1.44.0

COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Create tmp directory for air
RUN mkdir -p tmp

# Use air for hot reload
CMD ["air"]