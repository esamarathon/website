package db

import (
	"log"

	"github.com/olenedr/esamarathon/config"
	"github.com/pkg/errors"
	rDB "gopkg.in/gorethink/gorethink.v3"
)

var DB *rDB.Session

func Connect() error {
	session, err := rDB.Connect(config.DBConfig())

	if err != nil {
		return errors.Wrap(err, "db.Connect")
	}

	DB = session

	return nil
}

func Insert(table string, data map[string]interface{}) error {
	result, err := rDB.Table(table).Insert(data).RunWrite(DB)
	if err != nil {
		return errors.Wrap(err, "db.Insert")
	}

	log.Println("ID: " + result.GeneratedKeys[0] + " inserted into table " + table)
	return nil
}

func GetAll(table string) (*rDB.Cursor, error) {
	rows, err := rDB.Table(table).Run(DB)
	if err != nil {
		return nil, errors.Wrap(err, "db.GetAll")
	}

	return rows, nil
}
