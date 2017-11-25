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

func GetAllOrderBy(table string, index string) (*r.Cursor, error) {
	rows, err := r.Table(table).OrderBy(r.Desc(index)).Run(Session)
	if err != nil {
		return nil, errors.Wrap(err, "db.GetAllOrderBy")
	}

	return rows, nil
}

func GetPage(table string, page int, perPage int) (*r.Cursor, error) {
	skip := page * perPage
	rows, err := r.Table(table).OrderBy(r.Desc("created_at")).Skip(skip).Limit(perPage).Run(Session)
	if err != nil {
		return nil, errors.Wrap(err, "db.GetAll")
	}

	return rows, nil
}

func GetFilteredPage(table string, page, perPage int, published bool) (*r.Cursor, error) {
	skip := page * perPage
	rows, err := r.Table(table).Filter(map[string]interface{}{
		"published": true,
	}).OrderBy(r.Desc("created_at")).Skip(skip).Limit(perPage).Run(Session)

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

func GetCount(table string) (int, error) {
	cursor, err := r.Table(table).Count().Run(Session)
	if err != nil {
		return 0, errors.Wrap(err, "db.GetCount")
	}

	var count int
	cursor.One(&count)
	cursor.Close()

	return count, nil
}
