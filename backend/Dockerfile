# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /glimpse-api ./cmd/api

# Final stage
FROM alpine:latest

WORKDIR /

# Copy binary from builder stage
COPY --from=builder /glimpse-api /glimpse-api
COPY --from=builder /app/configs /configs
COPY --from=builder /app/migrations /migrations

# Expose port
EXPOSE 8080

# Run the application
CMD ["/glimpse-api"]
