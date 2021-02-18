package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/lu97chi/twittor/database"
	"github.com/lu97chi/twittor/handlers"
)

// init is invoked before main()
func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	if database.CheckConnection() == 0 {
		log.Fatal("Error when connecting to database")
		return
	}
	handlers.Handlers()
}
