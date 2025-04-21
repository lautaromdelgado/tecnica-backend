package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	JWTSecret string
	Port      string
	DBHost    string
	DBPort    string
	DBUser    string
	DBPass    string
	DBName    string
}

func LoadConfig() *Config {
	// Carga el archivo .env
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("No se pudo cargar .env, usando variables del sistema")
	}

	return &Config{
		JWTSecret: getEnv("JWT_SECRET", "hetmo_app_secret_key"),
		Port:      getEnv("PORT", "8080"),
		DBHost:    getEnv("DB_HOST", "localhost"),
		DBPort:    getEnv("DB_PORT", "3306"),
		DBUser:    getEnv("DB_USER", "root"),
		DBPass:    getEnv("DB_PASS", ""),
		DBName:    getEnv("DB_NAME", "hetmo_app"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
