package user

import (
	"net/http"
)

// HandleLogout deletes the session and redirects to root
func HandleLogout(w http.ResponseWriter, r *http.Request) {
	SessionStore.MaxAge(-1)
	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
}
