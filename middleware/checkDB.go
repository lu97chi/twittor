package middleware

import (
	"net/http"

	"github.com/lu97chi/twittor/database"
)

// CheckDB function to check the connection to database from the api
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if database.CheckConnection() == 0 {
			http.Error(rw, "Connection lost", 500)
			return
		}
		next.ServeHTTP(rw, r)
	}
}
