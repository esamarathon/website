package user

import (
	"github.com/olenedr/esamarathon/db"
	"github.com/pkg/errors"
	r "gopkg.in/gorethink/gorethink.v3"
)

func Create(username string) error {
	var data = map[string]interface{}{
		"username": username,
	}

	return db.Insert(Table, data)
}

func All() ([]User, error) {
	rows, err := db.AllOrderBy(Table, "username")
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

func Delete(id string) error {
	return db.Delete(Table, id)
}
