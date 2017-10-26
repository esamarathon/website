package user

import (
	"github.com/olenedr/esamarathon/db"
	"github.com/pkg/errors"
)

type User struct {
	ID   string `gorethink:"id" json:"id,omitempty"`
	Name string `gorethink:"name" json:"name,omitempty"`
}

type List struct {
	Users []User
}

const Table = "users"

func Insert(name string) error {
	var data = map[string]interface{}{
		"name": name,
	}

	return db.Insert(Table, data)
}

func Get() (List, error) {
	rows, err := db.GetAll(Table)
	var userList List
	var users []User
	if err != nil {
		return userList, errors.Wrap(err, "user.Get")
	}

	if err = rows.All(&users); err != nil {
		return userList, errors.Wrap(err, "user.Get")
	}

	userList.Users = users
	return userList, nil
}
