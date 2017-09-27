package user

import (
	"github.com/gorilla/sessions"
	"github.com/olenedr/esamarathon/config"
)

// SessionStore holds the session
var SessionStore = sessions.NewCookieStore([]byte(config.Config.SessionKey))
