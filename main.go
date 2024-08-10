package main

import (
	"log"
	"os"

  "github.com/joho/godotenv"
)

func seedAccount(store Storage, fname, lname, pw string) *Account {
	acc, err := NewAccount(fname, lname, pw)
	if err != nil {
		log.Fatal(err)
	}

	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}

	return acc
}

func seedAccounts(store Storage) {
  seedAccount(store, "Alice", "Smith", "password")
  seedAccount(store, "Bob", "Smith", "password")
  seedAccount(store, "Charlie", "Smith", "password")
}

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

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

  // seed accounts
  seedAccounts(store)

	server := NewAPIServer(":"+port, store)
	server.Run()
}
