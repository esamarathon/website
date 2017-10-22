package db

import (
	"fmt"

	"github.com/olenedr/esamarathon/config"
	"labix.org/v2/mgo"
)

// Connection is the live database connection
var Connection *mgo.Database

// Connect initializes the database connection
func Connect() error {

	i := mgo.DialInfo{
		Addrs:    []string{config.Config.DatabaseHost},
		Database: config.Config.Database,
		Username: config.Config.DatabaseUser,
		Password: config.Config.DatabasePassword,
	}
	fmt.Printf("Credentials: %v", i)

	s, err := mgo.DialWithInfo(&i)
	if err != nil {
		fmt.Println("Failed to connect to DB")
		fmt.Printf("Error: %v \n", err)
		return err
	}

	fmt.Println("Connected to the DB!")

	Connection = s.DB(config.Config.Database)
	return nil
}
