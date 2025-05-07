package configs

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Auth     AuthConfig
	Storage  StorageConfig
}

type ServerConfig struct {
	Port    int
	Timeout time.Duration
}

type DatabaseConfig struct {
	Host           string
	Port           int
	User           string
	Password       string
	DBName         string
	SSLMode        string
	MaxConnections int
}

type AuthConfig struct {
	JWTSecret   string
	TokenExpiry time.Duration
}

type StorageConfig struct {
	Type      string
	LocalPath string
	S3Bucket  string
	S3Region  string
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// Override with environment variables if they exist
	if host := os.Getenv("POSTGRES_HOST"); host != "" {
		config.Database.Host = host
	}
	
	if portStr := os.Getenv("POSTGRES_PORT"); portStr != "" {
		if port, err := strconv.Atoi(portStr); err == nil {
			config.Database.Port = port
		}
	}
	
	if user := os.Getenv("POSTGRES_USER"); user != "" {
		config.Database.User = user
	}
	
	if password := os.Getenv("POSTGRES_PASSWORD"); password != "" {
		config.Database.Password = password
	}
	
	if dbName := os.Getenv("POSTGRES_DB"); dbName != "" {
		config.Database.DBName = dbName
	}

	return &config, nil
}
