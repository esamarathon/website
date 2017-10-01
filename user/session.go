package user

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/olenedr/esamarathon/config"
)

// SessionStore holds the session
var SessionStore = sessions.NewCookieStore([]byte(config.Config.SessionKey))

func UserToSession(w http.ResponseWriter, r *http.Request, u User) {

}
