package handlers

import (
	"fmt"
	"net/http"

	"github.com/esamarathon/website/config"
	. "github.com/esamarathon/website/handlers/helpers"
	"github.com/esamarathon/website/models/user"
	"github.com/esamarathon/website/viewmodels"
	"golang.org/x/oauth2"
)

func AuthRedirect(w http.ResponseWriter, r *http.Request) {
	user.SessionStore.MaxAge(86400 * 7)
	redirectURL := config.TwitchOauthConfig.AuthCodeURL(config.OauthStateString)
	http.Redirect(w, r, redirectURL, http.StatusTemporaryRedirect)
}

func AuthCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != config.OauthStateString {
		// Someone might be trying to tinker with the request, good thing we're using  our trusty StateString
		fmt.Printf("Invalid oauth state, expected '%s', got '%s'\n", config.OauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get code
	code := r.FormValue("code")
	// Get the token
	token, err := config.TwitchOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		// Failed to get the toke
		fmt.Printf("Code exchange failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// Requesting user via token
	u, err := user.RequestTwitchUser(token)
	if err != nil {
		// Failed to get the user
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// Store the session
	if err := user.ToSession(w, r, u); err != nil {
		// Failed to store user
		HandleInternalError(w)
		return
	}

	// User was authorized, redirect to admin panel
	http.Redirect(w, r, "/admin", http.StatusTemporaryRedirect)
}

// HandleAuth returns login view
func HandleAuth(w http.ResponseWriter, r *http.Request) {
	Renderer.HTML(w, http.StatusOK, "login.html", viewmodels.Login())
}

// HandleLogout deletes the session and redirects to index
func HandleLogout(w http.ResponseWriter, r *http.Request) {
	user.SessionStore.MaxAge(-1)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
