package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port    int `json:"port"`
	DB_DSN  string `json:"db_dsn"`
	OpenAPI string `json:"openapi_url"`
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
port,_:=strconv.Atoi(os.Getenv("PORT"))
	return &Config{
		Port:    port,
		DB_DSN:  os.Getenv("DB_DSN"),
		OpenAPI: os.Getenv("OPENAPI_URL"),
	}
}
