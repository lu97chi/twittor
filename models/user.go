package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DBCollection collection of the database
const DBCollection = "users"

// User type
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FirstName string             `bson:"firstName" json:"firstName,omitempty"`
	LastName  string             `bson:"lastName" json:"lastName,omitempty"`
	BirthDay  time.Time          `bson:"birthDay" json:"birthDay,omitempty"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password,omitempty"`
	Avatar    string             `bson:"avatar" json:"avatar,omitempty"`
	Banner    string             `bson:"banner" json:"banner,omitempty"`
	Address   string             `bson:"address" json:"address,omitempty"`
	WebSite   string             `bson:"website" json:"website,omitempty"`
}
