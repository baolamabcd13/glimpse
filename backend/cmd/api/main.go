package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/baolamabcd13/glimpse/configs"
	"github.com/baolamabcd13/glimpse/internal/database"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	config, err := configs.LoadConfig("./configs/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Connect to database
	db, err := database.New(&config.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Run migrations
	migrationsPath := filepath.Join(".", "migrations")
	if err := db.RunMigrations(migrationsPath); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Initialize router
	r := gin.Default()

	// Basic health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"name":   "Glimpse API",
		})
	})

	// Start server
	serverAddr := fmt.Sprintf(":%d", config.Server.Port)
	log.Printf("Starting Glimpse API server on %s", serverAddr)
	if err := r.Run(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
