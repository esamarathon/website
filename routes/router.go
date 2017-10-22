package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/olenedr/esamarathon/handlers"
)

// GetRouter returns an instance of the Mux router
func GetRouter() *mux.Router {
	router := mux.NewRouter()

	router.PathPrefix("/static").Handler(handleStatic("public", "/static"))
	router.HandleFunc("/", handlers.Index).Methods("GET", "OPTIONS")
	router.HandleFunc("/schedule", handlers.Schedule).Methods("GET", "OPTIONS")
	router.HandleFunc("/news", handlers.News).Methods("GET", "OPTIONS")
	router.HandleFunc("/test", handlers.Test).Methods("GET", "OPTIONS")
	router.HandleFunc("/auth", handlers.AuthRedirect).Methods("GET")
	router.HandleFunc("/auth/callback", handlers.AuthCallback).Methods("GET")
	// http://localhost:3001/auth/callback

	return router
}

func handleStatic(dir, prefix string) http.HandlerFunc {
	fs := http.FileServer(http.Dir(dir))
	realHandler := http.StripPrefix(prefix, fs).ServeHTTP
	return func(w http.ResponseWriter, req *http.Request) {
		realHandler(w, req)
	}
}
