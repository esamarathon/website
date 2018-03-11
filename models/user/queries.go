package user

import (
	"github.com/esamarathon/website/db"
	"github.com/pkg/errors"
	r "gopkg.in/gorethink/gorethink.v3"
)

// Create inserts a new user entry to the DB
func Create(username string) (User, error) {
	var data = map[string]interface{}{
		"username": username,
	}

	_, err := db.Insert(Table, data)
	if err != nil {
		return User{}, errors.Wrap(err, "user.Create")
	}

	return GetUserByUsername(username)
}

// All returns all the users in the DB
func All() ([]User, error) {
	rows, err := db.GetAllByOrder(Table, "username", true)
	var users []User
	if err != nil {
		return users, errors.Wrap(err, "user.All")
	}

	if err = rows.All(&users); err != nil {
		return users, errors.Wrap(err, "user.All")
	}

	return users, nil
}

// Exists checks the db for the user by Username
func (u *User) Exists() (bool, error) {
	res, err := r.Table(Table).Filter(map[string]interface{}{
		"username": u.Username,
	}).Run(db.Session)

	if err != nil {
		return false, err
	}

	defer res.Close()

	var rows []interface{}
	res.All(&rows)
	if len(rows) == 0 {
		return false, nil
	}

	return true, nil
}

// GetUserByUsername takes a username and attempts to return the corresponding user in the DB
func GetUserByUsername(username string) (User, error) {
	res, err := r.Table(Table).Filter(map[string]interface{}{
		"username": username,
	}).Run(db.Session)

	var u User
	if err != nil {
		return u, errors.Wrap(err, "user.GetUserByUsername")
	}

	defer res.Close()

	if err = res.One(&u); err != nil {
		return u, err
	}

	return u, nil
}

// Delete deletes a user in the DB by ID
func Delete(id string) error {
	return db.Delete(Table, id)
}
