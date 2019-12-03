package api

import (
	"github.com/gorilla/mux"
)

/*
Registers the API routes to
*/
func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/news", News).Methods("GET")
}
