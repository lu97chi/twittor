package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/lu97chi/twittor/config"
	"github.com/lu97chi/twittor/models"
)

// GenerateToken get the token
func GenerateToken(u models.User) (string, error) {
	conf := config.New()
	seed := []byte(conf.JWT.Seed)
	claims := jwt.MapClaims{
		"email": u.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	payload, err := token.SignedString(seed)
	if err != nil {
		return payload, err
	}
	return payload, nil
}
