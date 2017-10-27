package handlers

import (
	"net/http"

	"github.com/rs/cors"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func init() {
	router.PathPrefix("/static").Handler(handleStatic("public", "/static"))
	router.HandleFunc("/", Index).Methods("GET", "OPTIONS")
	router.HandleFunc("/schedule", Schedule).Methods("GET", "OPTIONS")
	router.HandleFunc("/news", News).Methods("GET", "OPTIONS")
	router.HandleFunc("/auth", AuthRedirect).Methods("GET")
	router.HandleFunc("/auth/callback", AuthCallback).Methods("GET")
	router.HandleFunc("/login", HandleAuth).Methods("GET")
	router.HandleFunc("/logout", HandleLogout).Methods("GET")

	Routes("/admin", router)
}

func handleStatic(dir, prefix string) http.HandlerFunc {
	fs := http.FileServer(http.Dir(dir))
	realHandler := http.StripPrefix(prefix, fs).ServeHTTP
	return func(w http.ResponseWriter, req *http.Request) {
		realHandler(w, req)
	}
}

func Router(version string) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	return c.Handler(router)
}
