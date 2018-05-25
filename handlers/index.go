package handlers

import (
	"net/http"

	. "github.com/esamarathon/website/handlers/helpers"
	"github.com/esamarathon/website/viewmodels"
)

// Index returns index view
func Index(w http.ResponseWriter, r *http.Request) {
	data := viewmodels.Index()

	Renderer.HTML(w, http.StatusOK, "index.html", data)
}
