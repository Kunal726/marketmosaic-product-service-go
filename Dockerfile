# Build stage
FROM golang:1.21-alpine AS builder

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/build/marketmosaic-product-service cmd/marketmosaic-product-service/main.go

# Final stage
FROM alpine:latest

# Install CA certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/build/marketmosaic-product-service .
COPY --from=builder /app/.env .

# Create necessary directories
RUN mkdir -p /app/target

# Expose port
EXPOSE 8080

# Set environment variables
ENV SERVICE_NAME=marketmosaic-product-service
ENV GIN_MODE=release

# Run the application
CMD ["./marketmosaic-product-service"] 