package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/olenedr/esamarathon/str"
	"golang.org/x/oauth2"
)

type config struct {
	Port               string
	SessionKey         string
	SessionName        string
	Database           string
	DatabaseHost       string
	DatabaseUser       string
	DatabasePassword   string
	TwitchAuthURL      string
	TwitchClientID     string
	TwitchClientSecret string
	TwitchRedirectURL  string
	TwitchTokenURL     string
	TwitchAPIRootURL   string
}

// Config describes the env of the application
var (
	Config            config
	TwitchOauthConfig *oauth2.Config
	OauthStateString  string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Couldn't loading .env file")
	}

	Config = config{
		Port:               os.Getenv("PORT"),
		SessionKey:         os.Getenv("SESSION_KEY"),
		SessionName:        os.Getenv("SESSION_NAME"),
		Database:           os.Getenv("DB_NAME"),
		DatabaseHost:       os.Getenv("DB_HOST"),
		DatabaseUser:       os.Getenv("DB_USER"),
		DatabasePassword:   os.Getenv("DB_PW"),
		TwitchAuthURL:      os.Getenv("TWITCH_AUTH_URL"),
		TwitchClientID:     os.Getenv("TWITCH_CLIENT_ID"),
		TwitchClientSecret: os.Getenv("TWITCH_CLIENT_SECRET"),
		TwitchRedirectURL:  os.Getenv("TWITCH_REDIRECT_URL"),
		TwitchTokenURL:     os.Getenv("TWITCH_TOKEN_URL"),
		TwitchAPIRootURL:   os.Getenv("TWITCH_API_ROOT_URL"),
	}

	buildTwitchAuthConfig()
}

// BuildTwitchAuthConfig sets up the oauth config for authetication with Twitch.tv
func buildTwitchAuthConfig() {
	TwitchOauthConfig = &oauth2.Config{
		ClientID:     Config.TwitchClientID,
		ClientSecret: Config.TwitchClientSecret,
		Scopes:       []string{"user_read"},
		RedirectURL:  Config.TwitchRedirectURL,
		Endpoint: oauth2.Endpoint{
			AuthURL:  Config.TwitchAuthURL,
			TokenURL: Config.TwitchTokenURL,
		},
	}
	OauthStateString = str.RandStringRunes(10)
}
