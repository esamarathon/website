package db

import (
	"log"

	"github.com/olenedr/esamarathon/config"
	"github.com/pkg/errors"
	rDB "gopkg.in/gorethink/gorethink.v3"
)

var Session *rDB.Session

func Connect() error {
	session, err := rDB.Connect(config.DBConfig())

	if err != nil {
		return errors.Wrap(err, "db.Connect")
	}

	Session = session

	return nil
}

func Insert(table string, data map[string]interface{}) error {
	result, err := rDB.Table(table).Insert(data).RunWrite(Session)
	if err != nil {
		return errors.Wrap(err, "db.Insert")
	}

	log.Println("ID: " + result.GeneratedKeys[0] + " inserted into table " + table)
	return nil
}

func GetAll(table string) (*rDB.Cursor, error) {
	rows, err := rDB.Table(table).Run(Session)
	if err != nil {
		return nil, errors.Wrap(err, "db.GetAll")
	}

	return rows, nil
}
