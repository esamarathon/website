package admin

import (
	"fmt"
	"net/http"
	"olenedr/esamarathon/handlers"

	"github.com/olenedr/esamarathon/models/user"

	"github.com/gorilla/mux"
	"github.com/olenedr/esamarathon/middleware"
)

func Routes(base string, router *mux.Router) {
	requireAuth := middleware.AuthMiddleware
	router.HandleFunc(base+"/user", requireAuth(get)).Methods("GET")
	router.HandleFunc(base+"/user/create", requireAuth(create)).Methods("POST")
	router.HandleFunc(base, requireAuth(handlers.AdminIndex)).Methods("GET")
}

func create(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userName := r.Form.Get("username")

	if err := user.Insert(userName); err != nil || userName == "" {
		// @TODO: Handle error better
		fmt.Fprint(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func get(w http.ResponseWriter, r *http.Request) {
	users, err := user.Get()
	if err != nil {
		// @TODO: Handle error better
		fmt.Fprint(w, err)
		return
	}

	fmt.Printf("%#v\n", users)

	w.WriteHeader(http.StatusOK)
}
