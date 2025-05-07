package main

import (
	"fmt"
	"log"

	"github.com/baolamabcd13/glimpse/configs"
	"github.com/gin-gonic/gin"
)

func main() {
    // Load configuration
    config, err := configs.LoadConfig("./configs/config.yaml")
    if err != nil {
        log.Fatalf("Failed to load configuration: %v", err)
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