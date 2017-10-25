package db

import (
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
