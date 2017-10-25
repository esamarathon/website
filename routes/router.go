package routes

import (
	"net/http"

	"github.com/olenedr/esamarathon/routes/admin"
	"github.com/rs/cors"

	"github.com/gorilla/mux"
	"github.com/olenedr/esamarathon/handlers"
)

var router = mux.NewRouter()

func init() {
	router.PathPrefix("/static").Handler(handleStatic("public", "/static"))
	router.HandleFunc("/", handlers.Index).Methods("GET", "OPTIONS")
	router.HandleFunc("/schedule", handlers.Schedule).Methods("GET", "OPTIONS")
	router.HandleFunc("/news", handlers.News).Methods("GET", "OPTIONS")
	router.HandleFunc("/test", handlers.Test).Methods("GET", "OPTIONS")
	router.HandleFunc("/auth", handlers.AuthRedirect).Methods("GET")
	router.HandleFunc("/auth/callback", handlers.AuthCallback).Methods("GET")
	router.HandleFunc("/login", handlers.HandleAuth).Methods("GET")
	router.HandleFunc("/logout", handlers.HandleLogout).Methods("GET")

	admin.Routes("/admin", router)
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
