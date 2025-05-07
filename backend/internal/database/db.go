package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/baolamabcd13/glimpse/configs"
	"github.com/jackc/pgx/v5/pgxpool"
)

// DB represents the database connection pool
type DB struct {
    Pool *pgxpool.Pool
}

// New creates a new database connection
func New(config *configs.DatabaseConfig) (*DB, error) {
    connString := fmt.Sprintf(
        "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
        config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
    )

    // Create a connection pool
    poolConfig, err := pgxpool.ParseConfig(connString)
    if err != nil {
        return nil, fmt.Errorf("unable to parse connection string: %w", err)
    }

    // Set max connections (ensure it's at least 1)
    maxConns := int32(config.MaxConnections)
    if maxConns < 1 {
        maxConns = 1
    }
    poolConfig.MaxConns = maxConns

    // Create the connection pool
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
    if err != nil {
        return nil, fmt.Errorf("unable to create connection pool: %w", err)
    }

    // Verify connection
    if err := pool.Ping(ctx); err != nil {
        return nil, fmt.Errorf("unable to ping database: %w", err)
    }

    log.Println("Successfully connected to database")
    return &DB{Pool: pool}, nil
}

// Close closes the database connection
func (db *DB) Close() {
    if db.Pool != nil {
        db.Pool.Close()
    }
}
