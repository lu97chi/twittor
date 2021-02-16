package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/lu97chi/twittor/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertUser function to create user
func InsertUser(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading the .env file")
	}
	db := MongoCN.Database(os.Getenv("MONGO_DATABASE"))
	col := db.Collection(models.DBCollection)
	u.Password, _ = EncriptPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
