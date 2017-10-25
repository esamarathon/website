package user

import (
	"github.com/olenedr/esamarathon/db"
	"github.com/pkg/errors"
)

type User struct {
	ID   string `gorethink:"id" json:"id,omitempty"`
	Name string `gorethink:"name" json:"name,omitempty"`
}

const Table = "users"

func Insert(name string) error {
	var data = map[string]interface{}{
		"name": name,
	}

	return db.Insert(Table, data)
}

func Get() ([]User, error) {
	rows, err := db.GetAll(Table)
	var users []User
	if err != nil {
		return users, errors.Wrap(err, "user.Get")
	}

	if err = rows.All(&users); err != nil {
		return users, errors.Wrap(err, "user.Get")
	}

	return users, nil
}
