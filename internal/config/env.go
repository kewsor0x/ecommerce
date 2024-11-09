package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
	SSLMode  string
}

// variable to not initilize the function every time we call it
var Envs = InitConfig()

func InitConfig() *Config {
	godotenv.Load() // load the .env file

	return &Config{
		User:     getEnv("DB_USER", "root"),
		Password: getEnv("DB_PASSWORD", "root"),
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		Name:     getEnv("DB_NAME", "ecom"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}
}

// function to load the env file, if not use default values
func getEnv(key, default_value string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return default_value
}
