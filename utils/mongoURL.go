package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// GetMongoURL function to get the url connection to mongo
func GetMongoURL() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading the .env file")
	}
	username := os.Getenv("MONGO_USER")
	password := os.Getenv("MONGO_PASSWORD")
	url := os.Getenv("MONGO_URL")
	database := os.Getenv("MONGO_DATABASE")
	options := os.Getenv("MONGO_CONNECTION_OPTIONS")
	databaseConnection := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?%s", username, password, url, database, options)
	return databaseConnection
}
