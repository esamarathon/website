package db

import (
	"log"

	"github.com/olenedr/esamarathon/config"
	"github.com/pkg/errors"
	r "gopkg.in/gorethink/gorethink.v3"
)

var Session *r.Session

func Connect() error {
	session, err := r.Connect(config.DBConfig())

	if err != nil {
		return errors.Wrap(err, "db.Connect")
	}

	Session = session

	return nil
}

func Delete(table, id string) error {
	_, err := r.Table(table).Get(id).Delete().Run(Session)
	return err
}

func Update(table string, id string, data map[string]interface{}) error {
	_, err := r.Table(table).Get(id).Update(data).RunWrite(Session)
	if err != nil {
		return errors.Wrap(err, "db.Update")
	}

	return nil
}

func Insert(table string, data map[string]interface{}) error {
	result, err := r.Table(table).Insert(data).RunWrite(Session)
	if err != nil {
		return errors.Wrap(err, "db.Insert")
	}

	log.Println("ID: " + result.GeneratedKeys[0] + " inserted into table " + table)
	return nil
}

func GetAll(table string) (*r.Cursor, error) {
	rows, err := r.Table(table).Run(Session)
	if err != nil {
		return nil, errors.Wrap(err, "db.GetAll")
	}

	return rows, nil
}

func GetOneById(table, id string) (*r.Cursor, error) {
	cursor, err := r.Table(table).Get(id).Run(Session)
	if err != nil {
		return nil, errors.Wrap(err, "db.GetOneByID")
	}

	return cursor, err
}
