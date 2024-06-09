package main

import (
	"log"
	"os"
)

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT not set in .env")
	}

	server := NewAPIServer(":" + port, store)
	server.Run()
}

