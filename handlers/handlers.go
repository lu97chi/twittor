package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/lu97chi/twittor/middleware"
	"github.com/lu97chi/twittor/routers"
	"github.com/rs/cors"
)

// Handlers set Port, handler and set the server ready
func Handlers() {
	router := mux.NewRouter()
	PORT := os.Getenv("PORT")
	router.HandleFunc("/register", middleware.CheckDB(routers.Register)).Methods("POST")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
