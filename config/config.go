package config

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/esamarathon/website/str"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	rDB "gopkg.in/gorethink/gorethink.v3"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

type config struct {
	Port               string
	ArticlesPerPage    int
	LiveMode           bool
	SiteURL            string
	FrontpageDataPath  string
	MarkdownExtensions blackfriday.Extensions
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
	frontpageDataPath, err := filepath.Abs("")
	if err != nil {
		log.Println("Failed to create path, defaulting to tmp folder.")
		frontpageDataPath = "/tmp/frontpage.json"
	} else {
		frontpageDataPath = frontpageDataPath + "/templates/frontpage.json"
	}

	Config = config{
		Port:               os.Getenv("PORT"),
		ArticlesPerPage:    articlesPerPage,
		LiveMode:           liveMode,
		SiteURL:            os.Getenv("SITE_URL"),
		FrontpageDataPath:  frontpageDataPath,
		MarkdownExtensions: blackfriday.CommonExtensions | blackfriday.HardLineBreak,
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
