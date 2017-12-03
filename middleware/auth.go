package middleware

import (
	"net/http"

	"github.com/esamarathon/website/models/user"
)

// AuthMiddleware verifies that the current user is logged in before
// being allowed to access the given http.HandlerFunc
func AuthMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Retrieve the current user from the session
		u, err := user.FromSession(r)

		if err != nil {
			// Redirect to login
			redirect(w, r)
			return
		}
		// Got a user, check if user exists
		exists, err := u.Exists()

		if err != nil || !exists {
			// Got an error or user doesn't exist
			// redirect to login
			redirect(w, r)
			return
		}

		// User is authenticated, proceed
		h.ServeHTTP(w, r)
		return
	})
}

// redirect redirects the user to the login-page
func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
}
