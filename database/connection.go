package database

import (
	"context"
	"log"

	"github.com/lu97chi/twittor/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCN is the variable used across the proyect
var MongoCN = ConnectDB()
var clientOptions = options.Client().ApplyURI(utils.GetMongoURL())

// ConnectDB is a function to connect to mongoatlas
func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Connected to database")
	return client
}

// CheckConnection to detect the database connection
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
