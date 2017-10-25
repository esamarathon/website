package middleware

import (
	"net/http"

	"github.com/olenedr/esamarathon/user"
)

func AuthMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := user.UserFromSession(r)

		// @TODO: check database for users

		if err != nil {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}

		h.ServeHTTP(w, r)
		return
	})
}
