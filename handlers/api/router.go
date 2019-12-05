package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

/*
Registers the API routes to
*/
func RegisterRoutes(router *mux.Router) {
	router.Use(setCORSHeader)
	router.HandleFunc("/news", News).Methods("GET")
}

func setCORSHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}
