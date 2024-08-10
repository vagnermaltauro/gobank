package main

import (
	"flag"
	"fmt"
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

  fmt.Println("Created account:", acc)

	return acc
}

func seedAccounts(store Storage) {
	seedAccount(store, "Alice", "Smith", "password")
	seedAccount(store, "Bob", "Smith", "password")
	seedAccount(store, "Charlie", "Smith", "password")
}

func main() {
	seed := flag.Bool("seed", false, "seed database")
	flag.Parse()

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
	if *seed {
		fmt.Println("Seeding database")
		seedAccounts(store)
	}

	server := NewAPIServer(":"+port, store)
	server.Run()
}
