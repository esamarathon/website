package user

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/olenedr/esamarathon/config"
	"golang.org/x/oauth2"
)

func RequestTwitchUser(token *oauth2.Token) {
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
		// handle
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%v", string(body))

}
