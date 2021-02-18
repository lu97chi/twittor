package database

import (
	"github.com/lu97chi/twittor/models"
	"golang.org/x/crypto/bcrypt"
)

// Login function to login
func Login(email string, password string) (models.User, bool) {
	usr, exists, _ := CheckUserExists(email)
	if exists == false {
		return usr, false
	}

	hashedPassword := []byte(usr.Password)

	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		return usr, false
	}
	return usr, true
}
