package handlers

import (
	"fmt"
	"net/http"

	"github.com/olenedr/esamarathon/config"
	"github.com/olenedr/esamarathon/models/user"
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
		fmt.Printf("Invalid oauth state, expected '%s', got '%s'\n", config.OauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	code := r.FormValue("code")
	token, err := config.TwitchOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Printf("Code exchange failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	u, err := user.RequestTwitchUser(token)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// Store the session
	if err := user.UserToSession(w, r, u); err != nil {
		http.Redirect(w, r, "500.html", http.StatusTemporaryRedirect)
		return
	}

	http.Redirect(w, r, "/admin", http.StatusTemporaryRedirect)
}

// Index returns index view
func HandleAuth(w http.ResponseWriter, r *http.Request) {
	data := getPagedata()
	renderer.HTML(w, http.StatusOK, "login.html", data)
}

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	user.SessionStore.MaxAge(-1)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
