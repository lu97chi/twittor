package database

import (
	"context"
	"time"

	"github.com/lu97chi/twittor/config"
	"github.com/lu97chi/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

// CheckUserExists function to see if user exits
func CheckUserExists(email string) (models.User, bool, string) {
	conf := config.New()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN().Database(conf.MongoDB.Database)
	col := db.Collection(models.DBCollection)

	condition := bson.M{"email": email}
	var result models.User
	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
