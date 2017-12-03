package user

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"

	"github.com/esamarathon/website/config"
	"github.com/gorilla/sessions"
)

// SessionStore holds the session
var SessionStore = sessions.NewCookieStore([]byte(config.Config.SessionKey))

// ToSession attaches the user to the session
func ToSession(w http.ResponseWriter, r *http.Request, u User) error {
	// Get the session
	session, err := SessionStore.Get(r, config.Config.SessionName)
	if err != nil {
		return err
	}

	// Json marshal user object
	userJSON, err := json.Marshal(u)
	if err != nil {
		return err
	}

	// Insert the marshaled user in the session
	session.Values["user"] = string(userJSON)

	// Store the new session
	return SessionStore.Save(r, w, session)
}

// FromSession retrieves the user object from session
func FromSession(r *http.Request) (User, error) {
	// Define our user object
	var u User
	// Get the session
	session, err := SessionStore.Get(r, config.Config.SessionName)

	if err != nil {
		// Error retrieving session
		return u, err
	}

	if session.Values["user"] == nil {
		// No user in session
		err := errors.New("Session values empty")
		return u, errors.Wrap(err, "FromSession")
	}

	// Get the user json string
	userStr := session.Values["user"].(string)
	if err := json.Unmarshal([]byte(userStr), &u); err != nil {
		// Failed to unmarshal, returning
		return u, errors.Wrap(err, "FromSession")
	}

	// Return user-object based on username
	return GetUserByUsername(u.Username)
}

// SetSessionAge sets the lifetime of the session
func SetSessionAge(age int) {
	SessionStore.MaxAge(age)
}

// SetFlashMessage attaches a flash message to the session
func SetFlashMessage(w http.ResponseWriter, r *http.Request, name string, message string) error {
	// Get the sessiont
	session, _ := SessionStore.Get(r, config.Config.SessionName)
	// Attach message
	session.AddFlash(message, name)
	// Store the new session
	return SessionStore.Save(r, w, session)
}

// GetFlashMessage retrieves a flash message from the session
func GetFlashMessage(w http.ResponseWriter, r *http.Request, name string) string {
	// Retrieve the session
	session, err := SessionStore.Get(r, config.Config.SessionName)
	// Return empty string if it fails
	if err != nil {
		return ""
	}
	// Get all the messages by name
	flashes := session.Flashes(name)
	// Save the session after emptying messages
	session.Save(r, w)
	// If there are message(s) we return the first we find
	if len(flashes) > 0 {
		return flashes[0].(string)
	}
	// No message was found, return empty string
	return ""
}
