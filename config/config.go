package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	Port               string
	Database           string
	DatabaseHost       string
	DatabaseUser       string
	DatabasePassword   string
	TwitchAuthURL      string
	TwitchClientID     string
	TwitchClientSecret string
	TwitchRedirectURL  string
	TwitchTokenURL     string
}

// Config describes the env of the application
var Config config

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Couldn't loading .env file")
	}

	Config = config{
		Port:               os.Getenv("PORT"),
		Database:           os.Getenv("DB_NAME"),
		DatabaseHost:       os.Getenv("DB_HOST"),
		DatabaseUser:       os.Getenv("DB_USER"),
		DatabasePassword:   os.Getenv("DB_PW"),
		TwitchAuthURL:      os.Getenv("TWITCH_AUTH_URL"),
		TwitchClientID:     os.Getenv("TWITCH_CLIENT_ID"),
		TwitchClientSecret: os.Getenv("TWITCH_CLIENT_SECRET"),
		TwitchRedirectURL:  os.Getenv("TWITCH_REDIRECT_URL"),
		TwitchTokenURL:     os.Getenv("TWITCH_TOKEN_URL"),
	}

	// user.BuildTwitchAuthConfig()
}
