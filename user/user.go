package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/olenedr/esamarathon/config"
	"golang.org/x/oauth2"
)

type User struct {
	ID       string `json:"id,omitempty"`
	Username string `json:"user_name,omitempty"`
}

type TwitchResponse struct {
	User       User `json:"token,omitempty"`
	Identified bool `json:"identified,omitempty"`
}

func RequestTwitchUser(token *oauth2.Token) (User, error) {
	c := &http.Client{}
	req, err := http.NewRequest("GET", config.Config.TwitchAPIRootURL, nil)
	if err != nil {
		// handle
	}
	fmt.Println("HELLO TOKEN HERE YES", token.AccessToken)

	req.Header.Add("Accept", "application/vnd.twitchtv.v5+json")
	req.Header.Add("Authorization", "OAuth "+token.AccessToken)
	req.Header.Add("Client-ID", config.Config.TwitchClientID)
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("%v", err)
	}
	defer resp.Body.Close()

	var response TwitchResponse
	var user User
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Printf("%v", err)
		return user, err
	}

	fmt.Printf("%v", response.User)

	return user, nil

}
