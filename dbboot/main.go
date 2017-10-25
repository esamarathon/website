package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/olenedr/esamarathon/db"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
		return
	}

	if err := db.Connect(); err != nil {
		log.Println("Could not connect to the database:", err)
		return
	}
	log.Println("Successfully connected to the database")

	if err := db.Migrate(); err != nil {
		log.Println("Could not ensure database structure:", err)
		return
	}
	log.Println("Successfully migrated database")

	// @TODO: Add flag to determine whether to seed or not
	if err := db.Seed(); err != nil {
		log.Println("Error while seeding:", err)
		return
	}
	log.Println("Successfully seeded the database")
}
