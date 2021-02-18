package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/lu97chi/twittor/database"
	"github.com/lu97chi/twittor/jwt"
	"github.com/lu97chi/twittor/models"
)

// Login route to login
func Login(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	var u models.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(rw, "User or Password are incorrect"+err.Error(), 400)
		return
	}
	if len(u.Email) == 0 {
		http.Error(rw, "Email is required", 400)
		return
	}

	usr, exists := database.Login(u.Email, u.Password)
	if exists == false {
		http.Error(rw, "User or Password are incorrect", 400)
		return
	}

	token, tokenError := jwt.GenerateToken(usr)
	if tokenError != nil {
		http.Error(rw, "Something was wrong generating token"+tokenError.Error(), 400)
		return
	}

	response := models.LoginResponse{
		Token: token,
	}

	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(response)

	expiration := time.Now().Add(24 * time.Hour)
	http.SetCookie(rw, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expiration,
	})
}
