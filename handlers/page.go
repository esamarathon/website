package handlers


import (
	"net/http"

	. "github.com/esamarathon/website/handlers/helpers"
	"github.com/esamarathon/website/models/user"
	"github.com/esamarathon/website/models/page"
	"github.com/esamarathon/website/viewmodels"
	"github.com/gorilla/mux"
)

func Page(w http.ResponseWriter, r *http.Request) {
	fname := mux.Vars(r)["name"]

	// Check auth filter
	_, err := user.FromSession(r)
	// Define an article variable
	var p page.Page
	if err == nil {
		// Attempt to find the article (no filter for published status due to authed user)
		p, err = page.GetFromName(fname, nil)
	} else {
		// Request a the published article
		published := true
		p, err = page.GetFromName(fname, &published)
	}

	if err != nil {
		// Not found, return 404
		HandleNotFound(w, r)
		return
	}

	// Build the markup
	p.ParseHTML()
	p.FormatTimestamp()

	// Prepare the view
	data := viewmodels.Article()
	data.Article = p.Article
	data.Layout.Meta.Title = p.Title + " - ESA Marathon"

	Renderer.HTML(w, http.StatusOK, "standardpage.html", data)
}