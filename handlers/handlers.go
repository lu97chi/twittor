package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lu97chi/twittor/config"
	"github.com/lu97chi/twittor/middleware"
	"github.com/lu97chi/twittor/routers"
	"github.com/rs/cors"
)

// Handlers set Port, handler and set the server ready
func Handlers() {
	conf := config.New()
	router := mux.NewRouter()
	PORT := conf.Network.Port
	router.HandleFunc("/register", middleware.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middleware.CheckDB(routers.Login)).Methods("POST")

	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
