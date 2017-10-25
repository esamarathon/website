package handlers

import (
	"fmt"
	"html/template"
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

	fmt.Println("TOKEN GET:", token.AccessToken)

	u, err := user.RequestTwitchUser(token)
	if err != nil {
		fmt.Printf("Failed to get the user '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// Store the session
	user.UserToSession(w, r, u)

	fmt.Println("User authenticated", u.Username)
	http.Redirect(w, r, "/admin", http.StatusTemporaryRedirect)
}

func HandleAuth(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/login.html")
	if err != nil {
		//@TODO: Some better error handeling needed
		fmt.Fprint(w, err)
		return
	}

	t.Execute(w, nil)
}

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	middleware.SessionStore.MaxAge(-1)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
