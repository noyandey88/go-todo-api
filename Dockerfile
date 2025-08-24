# Build stage
FROM golang:1.24-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Generate Swagger documentation
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g cmd/main.go

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-todo-app ./cmd/main.go

# Final stage
FROM alpine:latest

# Add necessary packages
RUN apk --no-cache add ca-certificates tzdata

# Set working directory
WORKDIR /app

# Copy the binary from builder
COPY --from=builder /go-todo-app /app/

# Copy configuration files
COPY --from=builder /app/configs /app/configs

# Copy Docker-specific config and rename it to be used by default
COPY --from=builder /app/configs/config.docker.yaml /app/configs/config.yaml

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["/app/go-todo-app"]