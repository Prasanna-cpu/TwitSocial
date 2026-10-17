package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port           string
	DBUrlMigration string
	SecretJwt      string
	DBHost         string
	DBUser         string
	DBPassword     string
	DBPort         string
	DBName         string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}
	return &Config{
		Port:           os.Getenv("PORT"),
		DBUrlMigration: os.Getenv("DATABASE_URL"),
		SecretJwt:      os.Getenv("SECRET_KEY"),
		DBHost:         os.Getenv("DATABASE_HOST"),
		DBUser:         os.Getenv("DATABASE_USER"),
		DBPassword:     os.Getenv("DATABASE_PASSWORD"),
		DBPort:         os.Getenv("DATABASE_PORT"),
		DBName:         os.Getenv("DATABASE_NAME"),
	}, nil
}
