package database

import "golang.org/x/crypto/bcrypt"

// EncriptPassword function to encrypt password
func EncriptPassword(pass string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(bytes), err
}
