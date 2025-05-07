# Glimpse Backend

Backend service for Glimpse social media platform built with Go, GraphQL, PostgreSQL, and WebSockets.

## Technologies

- Go (Golang)
- GraphQL (gqlgen)
- PostgreSQL
- JWT Authentication
- WebSockets for real-time features

## Getting Started

### Prerequisites

- Go 1.20+
- PostgreSQL 14+
- Docker (optional)

### Setup

1. Clone the repository
2. Install dependencies: `go mod download`
3. Configure database in `configs/config.yaml`
4. Run the application: `go run cmd/api/main.go`

## Project Structure

- `cmd/api`: Application entry point
- `internal`: Private application code
  - `auth`: Authentication logic
  - `graph`: GraphQL schema and resolvers
  - `models`: Database models
  - `repository`: Database access layer
  - `services`: Business logic
  - `websocket`: WebSocket manager
- `pkg`: Public libraries
- `migrations`: Database migrations
- `configs`: Configuration files
