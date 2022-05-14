package main

import (
	"log"
	"os"

	"github.com/johnmcoro/quetzalapi/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	server := server.New()
	server.Logger.Info("Server started")
	if os.Getenv("LOCAL_MIGRATIONS") == "true" {
		migrationErr := server.Migrate()
		if migrationErr != nil {
			log.Println("migration error ", migrationErr)
		}
	}
	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}

}
