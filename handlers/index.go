package handlers

import (
	"net/http"

	"github.com/dannyvankooten/grender"
	"github.com/olenedr/esamarathon/viewmodels"
)

var renderer = grender.New(grender.Options{
	TemplatesGlob: "templates/*.html",
	PartialsGlob:  "templates/partials/*.html",
})

// Index returns index view
func Index(w http.ResponseWriter, r *http.Request) {
	data := viewmodels.Index()

	renderer.HTML(w, http.StatusOK, "index.html", data)
}
