package main

import (
	"flag"
	"log"

	"github.com/esamarathon/website/db"
	"github.com/esamarathon/website/db/seed"
	"github.com/joho/godotenv"
)

var (
	seedPtr = flag.Bool("seed", false, "Will run seeds")
)

func main() {
	flag.Parse()
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file. Continuing without one.")
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

	// If `--seed` is present
	if !*seedPtr {
		return
	}

	if err := seed.Seed(); err != nil {
		log.Println("Error while seeding:", err)
		return
	}
	log.Println("Successfully seeded the database")
}
