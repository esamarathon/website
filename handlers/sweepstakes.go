package handlers

import (
	"net/http"

	"github.com/esamarathon/website/viewmodels"
)

// Index returns index view
func Sweepstakes(w http.ResponseWriter, r *http.Request) {
	var data = viewmodels.Sweepstakes()
	renderer.HTML(w, http.StatusOK, "sweepstakes.html", data)
}
