package helpers

import (
	"net/http"

	"github.com/esamarathon/website/viewmodels"
)

// HandleNotFound handles the requests that doesn't have route associated with it
func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	Renderer.HTML(w, http.StatusNotFound, "404.html", viewmodels.Error())
}

// HandleInternalError handles requests that
// result in an internal server error
func HandleInternalError(w http.ResponseWriter) {
	Renderer.HTML(w, http.StatusInternalServerError, "500.html", viewmodels.Error())
}
