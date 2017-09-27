package handlers

import (
	"fmt"
	"net/http"

	"github.com/olenedr/esamarathon/config"
	"github.com/olenedr/esamarathon/user"
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
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", config.OauthStateString, state)
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

	fmt.Println("TOKEN GET:", token.AccessToken)

	user.RequestTwitchUser(token)

	// user, _ := getGithubUser(token)
	// user.Organizations, _ = getGithubOrgs(token)

	// // Store the session
	// userToSession(w, r, user)

	fmt.Println("User authenticated")
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
