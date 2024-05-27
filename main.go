package main

import (
	"dht11_server/api"
	db "dht11_server/db/postgres"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	envFile, err := godotenv.Read(".env.local")

	PORT := "4322"
	if err != nil {
		log.Fatal(err)
	}
	if len(envFile["SERVER_PORT"]) > 0 {
		PORT = envFile["SERVER_PORT"]
	}

	store, err := db.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(fmt.Sprintf(":%s", PORT), store)
	server.Run()
}
