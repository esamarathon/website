package main

import (
	"testing"

	"github.com/olenedr/esamarathon/config"
	r "gopkg.in/gorethink/gorethink.v3"
)

func TestConnect(t *testing.T) {
	session, err := r.Connect(config.DBConfig())
	checkError(err, t)

	if session == nil {
		t.Fatal("TestConnect session is nil")
	}
}

func TestGetAll(t *testing.T) {
	session, err := r.Connect(config.DBConfig())
	checkError(err, t)

	_, err = r.Table("articles").Run(session)
	checkError(err, t)
}

func TestInsert(t *testing.T) {
	session, err := r.Connect(config.DBConfig())
	checkError(err, t)

	result, err := r.Table("users").Insert(map[string]interface{}{
		"username": "test",
	}).RunWrite(session)
	checkError(err, t)

	if len(result.GeneratedKeys) == 0 {
		t.Fatal("TestInsert: No keys generated")
	}

	r.Table("users").Get(result.GeneratedKeys[0]).Delete().Run(session)
}
