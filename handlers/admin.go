package handlers

import (
	"net/http"

	"github.com/dannyvankooten/grender"
)

var adminRenderer = grender.New(grender.Options{
	TemplatesGlob: "templates_admin/*.html",
	PartialsGlob:  "templates_admin/partials/*.html",
})

func AdminIndex(w http.ResponseWriter, r *http.Request) {
	adminRenderer.HTML(w, http.StatusOK, "index.html", p)
}
