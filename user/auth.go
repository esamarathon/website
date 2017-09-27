package user

import (
	"net/http"
)

// HandleAuthRedirect redirects to the Twitch auth url
func HandleAuthRedirect(w http.ResponseWriter, r *http.Request) {

}

// HandleLogout deletes the session and redirects to root
func HandleLogout(w http.ResponseWriter, r *http.Request) {
	SessionStore.MaxAge(-1)
	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
}

// HandleAuthCallback handles the callback from the signin on Twitch
func HandleAuthCallback(w http.ResponseWriter, r *http.Request) {

}
