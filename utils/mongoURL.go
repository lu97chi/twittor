package utils

import (
	"fmt"

	"github.com/lu97chi/twittor/config"
)

// GetMongoURL function to get the url connection to mongo
func GetMongoURL() string {
	conf := config.New()

	username := conf.MongoDB.User
	password := conf.MongoDB.Password
	url := conf.MongoDB.URL
	database := conf.MongoDB.Database
	options := conf.MongoDB.Options
	databaseConnection := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?%s", username, password, url, database, options)
	return databaseConnection
}
