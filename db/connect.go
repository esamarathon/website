package db

import "labix.org/v2/mgo"
import "fmt"

const (
	host       = "localhost:27017"
	database   = "esamarathon"
	collection = "articles"
)

// Connection
var Connection *mgo.Database

// Connect initializes the database connection
func Connect() {

	i := mgo.DialInfo{
		Addrs:    []string{host},
		Database: database,
	}

	s, err := mgo.DialWithInfo(&i)
	if err != nil {
		fmt.Println("Failed to connect to DB")
		fmt.Printf("Error: %v \n", err)
	}

	Connection = s.DB(database)
}
