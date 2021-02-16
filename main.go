package main

import (
	"log"

	"github.com/lu97chi/twittor/database"
	"github.com/lu97chi/twittor/handlers"
)

func main() {
	if database.CheckConnection() == 0 {
		log.Fatal("Error when connecting to database")
		return
	}
	handlers.Handlers()
}
