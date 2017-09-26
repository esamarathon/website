package user

import (
	"github.com/olenedr/esamarathon/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var (
	twitchOauthConfig *oauth2.Config
)

// BuildTwitchAuthConfig sets up the oauth config for authetication with Twitch.tv
func BuildTwitchAuthConfig() {
	twitchOauthConfig = &oauth2.Config{
		ClientID:     config.Config.GithubClientID,
		ClientSecret: config.Config.GithubClientSecret,
		Scopes:       []string{"user"},
		RedirectURL:  config.Config.GithubRedirectURL,
		Endpoint:     github.Endpoint,
	}
}
