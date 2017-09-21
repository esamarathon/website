package db

import "labix.org/v2/mgo"
import "fmt"

const (
	host       = "localhost:27017"
	database   = "esamarathon"
	username   = "root"
	password   = "root"
	collection = "articles"
)

var Db *mgo.Database

// Connect initializes the database connection
func Connect() {
	s, err := mgo.Dial(host)
	if err != nil {
		fmt.Println("Failed to connect to DB")
	}

	Db = s.DB(database)
}
