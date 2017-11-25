package seed

/**
*	This file could probably be improved by the use of reflection(?)
* 	by building a slice of functions to iterate through and call in sequence
 */

import (
	"github.com/olenedr/esamarathon/db"
	"github.com/olenedr/esamarathon/models/user"

	"github.com/pkg/errors"
	rt "gopkg.in/gorethink/gorethink.v3"
)

var row interface{}

// Seed initiates the seeding of default data into the DB
func Seed() error {
	if err := users(); err != nil {
		return err
	}

	return nil
}

func users() error {
	t := "users"
	var users = []user.User{
		{
			ID:       "1",
			Username: "korkn",
		},
		{
			ID:       "2",
			Username: "egreb__",
		},
		{
			ID:       "3",
			Username: "morganbro",
		},
	}

	res, err := rt.Table(t).Insert(users).Run(db.Session)
	if err != nil {
		return err
	}
	defer res.Close()

	if err = validateResult(res, len(users), t); err != nil {
		return err
	}

	return nil
}

// Validates that all the entries have been inserted successfully
func validateResult(res *rt.Cursor, count int, name string) error {
	for res.Next(&row) {
		m := row.(map[string]interface{})
		if m["inserted"] != float64(count) {
			err := errors.New("Couldn't or didn't seed all " + name)
			return errors.Wrap(err, "seed."+name)
		}
	}
	return nil
}
