package middleware

import (
	"net/http"

	"github.com/dannyvankooten/grender"
	"github.com/olenedr/esamarathon/user"
)

var renderer = grender.New(grender.Options{
	TemplatesGlob: "templates/*.html",
})

func AuthMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := user.UserFromSession(r)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}

		h.ServeHTTP(w, r)
		return
	})
}
