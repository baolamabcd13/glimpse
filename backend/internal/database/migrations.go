package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// RunMigrations runs all SQL migrations in the migrations directory
func (db *DB) RunMigrations(migrationsPath string) error {
    log.Println("Running database migrations...")

    // Create migrations table if it doesn't exist
    _, err := db.Pool.Exec(context.Background(), `
        CREATE TABLE IF NOT EXISTS migrations (
            id SERIAL PRIMARY KEY,
            name VARCHAR(255) NOT NULL UNIQUE,
            applied_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
        )
    `)
    if err != nil {
        return fmt.Errorf("failed to create migrations table: %w", err)
    }

    // Get list of applied migrations
    appliedMigrations := make(map[string]bool)
    rows, err := db.Pool.Query(context.Background(), "SELECT name FROM migrations")
    if err != nil {
        return fmt.Errorf("failed to query migrations: %w", err)
    }
    defer rows.Close()

    for rows.Next() {
        var name string
        if err := rows.Scan(&name); err != nil {
            return fmt.Errorf("failed to scan migration name: %w", err)
        }
        appliedMigrations[name] = true
    }

    // Read migration files
    files, err := os.ReadDir(migrationsPath)
    if err != nil {
        return fmt.Errorf("failed to read migrations directory: %w", err)
    }

    var migrationFiles []string
    for _, file := range files {
        if !file.IsDir() && strings.HasSuffix(file.Name(), ".sql") {
            migrationFiles = append(migrationFiles, file.Name())
        }
    }

    // Sort migrations by name
    sort.Strings(migrationFiles)

    // Apply migrations
    for _, fileName := range migrationFiles {
        if appliedMigrations[fileName] {
            log.Printf("Migration %s already applied, skipping", fileName)
            continue
        }

        log.Printf("Applying migration: %s", fileName)
        filePath := filepath.Join(migrationsPath, fileName)
        
        content, err := os.ReadFile(filePath)
        if err != nil {
            return fmt.Errorf("failed to read migration file %s: %w", fileName, err)
        }

        // Begin transaction
        tx, err := db.Pool.Begin(context.Background())
        if err != nil {
            return fmt.Errorf("failed to begin transaction: %w", err)
        }

        // Execute migration
        _, err = tx.Exec(context.Background(), string(content))
        if err != nil {
            tx.Rollback(context.Background())
            return fmt.Errorf("failed to execute migration %s: %w", fileName, err)
        }

        // Record migration
        _, err = tx.Exec(context.Background(), "INSERT INTO migrations (name) VALUES ($1)", fileName)
        if err != nil {
            tx.Rollback(context.Background())
            return fmt.Errorf("failed to record migration %s: %w", fileName, err)
        }

        // Commit transaction
        if err := tx.Commit(context.Background()); err != nil {
            return fmt.Errorf("failed to commit transaction: %w", err)
        }

        log.Printf("Successfully applied migration: %s", fileName)
    }

    log.Println("All migrations applied successfully")
    return nil
}
