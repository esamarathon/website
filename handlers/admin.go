package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dannyvankooten/grender"
	"github.com/pkg/errors"

	"github.com/olenedr/esamarathon/models/user"

	"github.com/gorilla/mux"
	"github.com/olenedr/esamarathon/middleware"
)

func AdminRoutes(base string, router *mux.Router) {
	requireAuth := middleware.AuthMiddleware
	router.HandleFunc(base+"/user/create", requireAuth(create)).Methods("POST")
	router.HandleFunc(base, requireAuth(index)).Methods("GET", "POST")
}

var adminRenderer = grender.New(grender.Options{
	TemplatesGlob: "templates_admin/*.html",
})

func index(w http.ResponseWriter, r *http.Request) {
	userList, err := user.All()
	if err != nil {
		log.Println(errors.Wrap(err, "admin.index"))
		fmt.Fprint(w, err)
		return
	}

	adminRenderer.HTML(w, http.StatusOK, "index.html", userList)
}

func create(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userName := r.Form.Get("username")

	if err := user.Insert(userName); err != nil || userName == "" {
		// TODO:Handle error better
		fmt.Fprint(w, err)
		return
	}

	log.Println("redirect")
	http.Redirect(w, r, "/admin", http.StatusTemporaryRedirect)
}
