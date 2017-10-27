package setting

import (
	"log"
	"strconv"

	"github.com/olenedr/esamarathon/db"
	rt "gopkg.in/gorethink/gorethink.v3"
)

// Table is the DB table for this model
var Table = "settings"

// Setting is a key value object that describes how page functionality should act
type Setting struct {
	ID    string `json:"id,omitempty" gorethink:"id,omitempty"`
	Key   string `json:"key,omitempty" gorethink:"key"`
	Value string `json:"value,omitempty" gorethink:"value"`
}

// GetLiveMode returns a pointer to the Livemode setting
func GetLiveMode() *Setting {
	res, err := rt.Table(Table).Filter(map[string]interface{}{
		"key": "livemode",
	}).Run(db.DB)

	var s Setting
	err = res.One(&s)
	if err != nil {
		log.Fatalln(err)
	}

	return &s
}

func (s *Setting) AsBool() (bool, error) {
	b, err := strconv.ParseBool(s.Value)
	if err != nil {
		return false, err
	}
	return b, nil
}

// Toggle attempt to toggle a boolean setting
func (s *Setting) Toggle() error {
	b, err := s.AsBool()
	if err != nil {
		return err
	}
	if b {
		s.Value = "false"
	} else {
		s.Value = "true"
	}
	_, err = rt.Table(Table).Get(s.ID).Update(s).Run(db.DB)
	if err != nil {
		return err
	}
	return nil
}
