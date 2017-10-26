package user

import (
	"github.com/olenedr/esamarathon/db"
	rDB "gopkg.in/gorethink/gorethink.v3"
)

// Exists checks the db for the user by Username
func (u *User) Exists() (bool, error) {
	res, err := rDB.Table(table).Filter(map[string]interface{}{
		"Username": u.Username,
	}).Run(db.DB)

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
