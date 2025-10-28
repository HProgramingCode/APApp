package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port          string
	DBDriver      string
	DBName        string
	JWTSecretKey  string
	TokenLifeTime string
	Env           string
}

var AppConfig *Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	AppConfig = &Config{
		Port:          getEnv("PORT", "8080"),
		DBDriver:      getEnv("DB_DRIVER", "sqlite3"),
		DBName:        getEnv("DB_NAME", "app.db"),
		JWTSecretKey:  getEnv("JWT_SECRET_KEY", "default_secret"),
		TokenLifeTime: getEnv("TOKEN_LIFE_TIME", "1"),
		Env:           getEnv("ENV", "development"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
