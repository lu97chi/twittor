package routers

import (
	"encoding/json"
	"net/http"

	"github.com/lu97chi/twittor/database"
	"github.com/lu97chi/twittor/models"
)

// Register function to register a new user on the database
func Register(rw http.ResponseWriter, r *http.Request) {
	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(rw, "Request error"+err.Error(), 400)
		return
	}

	if len(u.Email) == 0 {
		http.Error(rw, "Email is required", 400)
		return
	}

	if len(u.Password) < 6 {
		http.Error(rw, "Must be at least 6 characters", 400)
		return
	}
	_, found, _ := database.CheckUserExists(u.Email)
	if found == true {
		http.Error(rw, "User is already created", 400)
		return
	}

	_, status, err := database.InsertUser(u)
	if err != nil {
		http.Error(rw, "Something went wrong when creating an user"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(rw, "Couldn't create a new user", 400)
		return
	}
	rw.WriteHeader(http.StatusCreated)
}
