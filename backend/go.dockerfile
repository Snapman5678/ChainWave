# Use Go 1.23 which matches the toolchain directive in go.mod
FROM golang:1.23-alpine

WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download and install the dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app, ensuring to specify the entry point
RUN go build -o api ./cmd/main.go  # Assuming your main file is here

# Expose port 8000 to the outside world
EXPOSE 8000

# Run the Go app
CMD ["./api"]
