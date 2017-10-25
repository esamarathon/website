package db

import (
	"github.com/olenedr/esamarathon/config"
	"github.com/pkg/errors"
	rDB "gopkg.in/gorethink/gorethink.v3"
)

func Connect() error {
	dbConf := config.Config
	_, err := rDB.Connect(rDB.ConnectOpts{
		Address:    dbConf.DatabaseHost,
		Database:   dbConf.Database,
		Username:   dbConf.DatabaseUser,
		Password:   dbConf.DatabasePassword,
		InitialCap: 10,
		MaxOpen:    10,
	})

	if err != nil {
		return errors.Wrap(err, "db.Connect")
	}

	return nil
}
