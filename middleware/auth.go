package middleware

import (
	"net/http"

	"github.com/olenedr/esamarathon/models/user"
)

func AuthMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u, err := user.UserFromSession(r)

		if err != nil {
			redirect(w, r)
			return
		}
		exists, err := u.Exists()

		if err != nil || !exists {
			redirect(w, r)
			return
		}

		h.ServeHTTP(w, r)
		return
	})
}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
}
