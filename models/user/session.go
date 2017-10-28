package user

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"

	"github.com/gorilla/sessions"
	"github.com/olenedr/esamarathon/config"
)

// SessionStore holds the session
var SessionStore = sessions.NewCookieStore([]byte(config.Config.SessionKey))

func UserToSession(w http.ResponseWriter, r *http.Request, u User) error {
	session, err := SessionStore.Get(r, config.Config.SessionName)
	if err != nil {
		return err
	}

	usrJSON, err := json.Marshal(u)
	if err != nil {
		return err
	}

	session.Values["user"] = string(usrJSON)

	return SessionStore.Save(r, w, session)
}

func UserFromSession(r *http.Request) (User, error) {
	var u User
	session, err := SessionStore.Get(r, config.Config.SessionName)
	if err != nil {
		return u, err
	}

	if session.Values["user"] == nil {
		err := errors.New("Session values empty")
		return u, errors.Wrap(err, "UserFromSession")
	}

	usrStr := session.Values["user"].(string)
	if err := json.Unmarshal([]byte(usrStr), &u); err != nil {
		return u, errors.Wrap(err, "UserFromSession")
	}

	return GetUserByUsername(u.Username)
}

func SetSessionAge(age int) {
	SessionStore.MaxAge(age)
}
