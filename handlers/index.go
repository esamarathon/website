package handlers

import (
	"net/http"

	"github.com/dannyvankooten/grender"
	"github.com/esamarathon/website/viewmodels"
)

var renderer = grender.New(grender.Options{
	TemplatesGlob: "templates_minified/*.html",
	PartialsGlob:  "templates_minified/partials/*.html",
})

// Index returns index view
func Index(w http.ResponseWriter, r *http.Request) {
	data := viewmodels.Index()

	renderer.HTML(w, http.StatusOK, "index.html", data)
}
