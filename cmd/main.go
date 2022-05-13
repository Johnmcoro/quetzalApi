package main

import (
	"log"

	"github.com/johnmcoro/quetzalapi/internal/server"
)

func main() {
	server := server.New()
	migrationErr := server.Migrate()
	if migrationErr != nil {
		log.Fatal(migrationErr)
	}
	err := server.Run()
	if err != nil {
		log.Fatal(err)
	}

}
