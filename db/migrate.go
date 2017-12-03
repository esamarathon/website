package db

import (
	"github.com/esamarathon/website/config"
	rt "gopkg.in/gorethink/gorethink.v3"
)

var tables = []string{
	"users",
	"articles",
}

var row interface{}

func Migrate() error {
	if err := ensureDBExists(); err != nil {
		return err
	}

	for _, t := range tables {
		if err := ensureTableExists(t); err != nil {
			return err
		}
	}
	return nil
}

func ensureDBExists() error {
	res, err := rt.DBList().Contains(config.Config.Database).Run(Session)
	if err != nil {
		return err
	}
	defer res.Close()

	for res.Next(&row) {
		if row.(bool) {
			return nil
		}
	}
	rt.DBCreate(config.Config.Database).Run(Session)

	return nil
}

func ensureTableExists(t string) error {
	res, err := rt.TableList().Contains(t).Run(Session)
	if err != nil {
		return err
	}
	defer res.Close()

	for res.Next(&row) {
		if row.(bool) {
			return nil
		}
	}

	if _, err = rt.TableCreate(t).Run(Session); err != nil {
		return err
	}
	return nil
}
