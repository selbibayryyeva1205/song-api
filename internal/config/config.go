package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port    int    `json:"port"`
	DB_DSN  string `json:"db_dsn"`
	OpenAPI string `json:"openapi_url"`
	Host    string `json:"host"`
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	return &Config{
		//	OpenAPI: os.Getenv("OPEN_API"),
		Host:    os.Getenv("HOST"),
		Port:    port,
		DB_DSN:  os.Getenv("DB_DSN"),
		OpenAPI: os.Getenv("OPENAPI_URL"),
	}
}
