package main

import (
	"testing"

	"github.com/esamarathon/website/db"

	"github.com/esamarathon/website/config"
	r "gopkg.in/gorethink/gorethink.v3"
)

var tables = []string{
	"users",
	"articles",
}

func TestConnect(t *testing.T) {
	session, err := r.Connect(config.DBConfig())
	checkError(err, t)

	if session == nil {
		t.Fatal("TestConnect session is nil")
	}
}

func TestGetAll(t *testing.T) {
	initDb(t)

	for _, table := range tables {
		_, err := db.GetAll(table)
		checkError(err, t)
	}
}

func TestInsertAndDelete(t *testing.T) {
	initDb(t)

	for _, table := range tables {
		var data map[string]interface{}
		switch table {
		case "users":
			data = map[string]interface{}{
				"username": "test",
			}

		case "article":
			data = map[string]interface{}{
				"title":     "test article",
				"published": false,
			}
		}

		id, err := db.Insert(table, data)
		checkError(err, t)

		if err = db.Delete(table, id); err != nil {
			checkError(err, t)
		}
	}
}

func TestGetAllByOrder(t *testing.T) {
	initDb(t)

	for _, table := range tables {
		var orderBy string
		var data []interface{}

		switch table {
		case "users":
			orderBy = "username"
		case "articles":
			orderBy = "title"
		}

		rows, err := db.GetAllByOrder(table, orderBy, true)
		checkError(err, t)

		checkError(rows.All(&data), t)
	}

}

func TestGetPage(t *testing.T) {
	initDb(t)

	for _, table := range tables {
		var desc string
		switch table {
		case "users":
			desc = "username"
		case "articles":
			desc = "created_at"
		}

		_, err := db.GetPage(table, desc, 0, 10)
		checkError(err, t)
	}
}

func TestGetFilteredPage(t *testing.T) {
	initDb(t)

	for _, table := range tables {
		var desc string
		var filter map[string]interface{}

		switch table {
		case "users":
			desc = "username"
			filter = map[string]interface{}{
				"username": "egreb__",
			}
		case "articles":
			desc = "created_at"
			filter = map[string]interface{}{
				"published": true,
			}
		}

		_, err := db.GetFilteredPage(table, desc, 0, 10, filter)
		checkError(err, t)
	}
}

func TestGetByFilter(t *testing.T) {
	initDb(t)

	for _, table := range tables {
		var filter map[string]interface{}

		switch table {
		case "users":
			filter = map[string]interface{}{
				"username": "egreb__",
			}
		case "articles":
			filter = map[string]interface{}{
				"published": true,
			}
		}

		_, err := db.GetByFilter(table, filter)
		checkError(err, t)
	}
}

func TestCount(t *testing.T) {
	initDb(t)

	for _, table := range tables {
		_, err := db.GetCount(table)
		checkError(err, t)
	}
}
