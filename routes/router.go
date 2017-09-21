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
	// router.Handle().Methods("GET", "OPTIONS")

	return router
}

func handleStatic(dir, prefix string) http.HandlerFunc {
	fs := http.FileServer(http.Dir(dir))
	realHandler := http.StripPrefix(prefix, fs).ServeHTTP
	return func(w http.ResponseWriter, req *http.Request) {
		realHandler(w, req)
	}
}
