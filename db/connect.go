package db

import (
	"log"

	"github.com/esamarathon/website/config"
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
	if err != nil {
		log.Println("ID: " + id + " deleted from table " + table)
	}
	return err
}

func Update(table string, id string, data map[string]interface{}) error {
	_, err := r.Table(table).Get(id).Update(data).RunWrite(Session)
	if err != nil {
		return errors.Wrap(err, "db.Update")
	}

	return nil
}

func Insert(table string, data map[string]interface{}) (string, error) {
	result, err := r.Table(table).Insert(data).RunWrite(Session)
	if err != nil {
		return "", errors.Wrap(err, "db.Insert")
	}

	id := result.GeneratedKeys[0]
	log.Println("ID: " + id + " inserted into table " + table)
	return id, nil
}

func GetAll(table string) (*r.Cursor, error) {
	rows, err := r.Table(table).Run(Session)
	if err != nil {
		return nil, errors.Wrap(err, "db.GetAll")
	}

	return rows, nil
}

func GetAllByOrder(table string, index string) (*r.Cursor, error) {
	rows, err := r.Table(table).OrderBy(r.Desc(index)).Run(Session)
	if err != nil {
		return nil, errors.Wrap(err, "db.GetAllOrderBy")
	}

	return rows, nil
}

// GetPage with pagination
func GetPage(table, desc string, page int, perPage int) (*r.Cursor, error) {
	skip := page * perPage
	rows, err := r.Table(table).OrderBy(r.Desc(desc)).Skip(skip).Limit(perPage).Run(Session)
	if err != nil {
		return nil, errors.Wrap(err, "db.GetAll")
	}

	return rows, nil
}

// GetFilteredPage with pagination and an additional filter parameter
func GetFilteredPage(table, desc string, page, perPage int, filter map[string]interface{}) (*r.Cursor, error) {
	skip := page * perPage
	rows, err := r.Table(table).Filter(filter).OrderBy(r.Desc(desc)).Skip(skip).Limit(perPage).Run(Session)

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

// GetByFilter: apply filter by an map[string]interface{}
// example: map[string]interface{}{"id": "something", "published": true}
func GetByFilter(table string, filter map[string]interface{}) (*r.Cursor, error) {
	cursor, err := r.Table(table).Filter(filter).Run(Session)
	if err != nil {
		return nil, errors.Wrap(err, "db.GetOneByIDFiltered")
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

func GetFilteredCount(table string, filter map[string]interface{}) (int, error) {
	cursor, err := r.Table(table).Filter(filter).Count().Run(Session)
	if err != nil {
		return 0, errors.Wrap(err, "db.GetCount")
	}

	var count int
	cursor.One(&count)
	cursor.Close()

	return count, nil
}
