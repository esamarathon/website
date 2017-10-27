package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/olenedr/esamarathon/config"
	"github.com/olenedr/esamarathon/db"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

const Table = "users"

type User struct {
	ID       string `gorethink:"id" json:"id,omitempty"`
	Username string `gorethink:"username" json:"user_name,omitempty"`
}

type TwitchResponse struct {
	User       User `json:"token,omitempty"`
	Identified bool `json:"identified,omitempty"`
}

type List struct {
	Users []User
}

func Insert(username string) error {
	var data = map[string]interface{}{
		"username": username,
	}

	return db.Insert(Table, data)
}

func All() (List, error) {
	rows, err := db.GetAll(Table)
	var userList List
	var users []User
	if err != nil {
		return userList, errors.Wrap(err, "user.All")
	}

	if err = rows.All(&users); err != nil {
		return userList, errors.Wrap(err, "user.All")
	}

	userList.Users = users
	return userList, nil
}

func RequestTwitchUser(token *oauth2.Token) (User, error) {
	c := &http.Client{}
	var res TwitchResponse

	req, err := http.NewRequest("GET", config.Config.TwitchAPIRootURL, nil)
	if err != nil {
		return User{}, err
	}

	req.Header.Add("Accept", "application/vnd.twitchtv.v5+json")
	req.Header.Add("Authorization", "OAuth "+token.AccessToken)
	req.Header.Add("Client-ID", config.Config.TwitchClientID)
	resp, err := c.Do(req)
	if err != nil {
		return User{}, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&res)

	// Check for err or empty struct
	if err != nil || res.User == (User{}) {
		fmt.Printf("%v", err)
		return User{}, err
	}

	return res.User, nil

}
