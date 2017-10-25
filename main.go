package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/olenedr/esamarathon/db"
	"github.com/olenedr/esamarathon/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	if err := db.Connect(); err != nil {
		log.Println("Could not connect to the database:", err)
	} else {
		log.Println("Successfully connected to the database")
	}

	router := routes.GetRouter()
	fmt.Println("Listening to localhost on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
