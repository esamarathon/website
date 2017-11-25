package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/olenedr/esamarathon/str"
	"golang.org/x/oauth2"
	rDB "gopkg.in/gorethink/gorethink.v3"
)

type config struct {
	Port               string
	ArticlesPerPage    int
	LiveMode           bool
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
	ScheduleAPIURL     string
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

	articlesPerPage, err := strconv.Atoi(os.Getenv("ARTICLES_PER_PAGE"))
	if err != nil {
		log.Println("Failed to parse numeric .env value, using default.")
		articlesPerPage = 10
	}
	liveMode, err := strconv.ParseBool(os.Getenv("LIVE_MODE"))
	if err != nil {
		log.Println("Failed to parse bool .env value, using default.")
		liveMode = false
	}

	Config = config{
		Port:               os.Getenv("PORT"),
		ArticlesPerPage:    articlesPerPage,
		LiveMode:           liveMode,
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
		ScheduleAPIURL:     os.Getenv("SCHEDULE_API_URL"),
	}

	if Config.ScheduleAPIURL == "" {
		log.Println("No Schedule API URL defined, utilizing backup")
		Config.ScheduleAPIURL = "https://horaro.org/-/api/v1/schedules/4311u8b52b04si7a1e"
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

// DBConfig returns the database connect options
func DBConfig() rDB.ConnectOpts {
	return rDB.ConnectOpts{
		Address:    Config.DatabaseHost,
		Database:   Config.Database,
		Username:   Config.DatabaseUser,
		Password:   Config.DatabasePassword,
		InitialCap: 10,
		MaxOpen:    10,
	}
}

// ToggleLiveMode toggles the bool in the config
func ToggleLiveMode() {
	Config.LiveMode = !Config.LiveMode
}
