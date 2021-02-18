package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// GetEnv function to get env variable
func GetEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading the .env file")
		return ""
	}
	return os.Getenv(key)
}
