package handlers

import (
	"net/http"

	. "github.com/esamarathon/website/handlers/helpers"
	"github.com/esamarathon/website/models/article"
	"github.com/esamarathon/website/models/user"
	"github.com/esamarathon/website/viewmodels"

	"github.com/gorilla/mux"
)

// News renders the news page
func News(w http.ResponseWriter, r *http.Request) {
	data := viewmodels.News()
	p := GetPagination(r)

	articles, err := article.Page(p, true)
	// If we failed to get the articles
	// we return the 500 error page
	if err != nil {
		HandleInternalError(w)
		return
	}

	// Reduce body to a teaser
	for i, a := range articles {
		a.ShortenBody()
		a.ParseHTML()
		a.FormatTimestamp()
		articles[i] = a
	}

	// Attach needed values
	data.Articles = articles
	data.NextPage = p + 1
	data.PrevPage = p - 1
	data.LastPage, err = article.PageCount(true)

	if err != nil {
		HandleInternalError(w)
		return
	}

	Renderer.HTML(w, http.StatusOK, "news.html", data)
}

// Article renders the page of a specific article
func Article(w http.ResponseWriter, r *http.Request) {
	// Get the ID
	id := mux.Vars(r)["id"]

	// Check auth filter
	_, err := user.FromSession(r)
	// Define an article variable
	var a article.Article
	if err == nil {
		// Attempt to find the article (no filter for published status due to authed user)
		a, err = article.Get(id, nil)
	} else {
		// Request a the published article
		published := true
		a, err = article.Get(id, &published)
	}

	if err != nil {
		// Not found, return 404
		HandleNotFound(w, r)
		return
	}

	// Build the markup
	a.ParseHTML()
	a.FormatTimestamp()

	// Prepare the view
	data := viewmodels.Article()
	data.Article = a
	data.Layout.Meta.Title = a.Title + " - ESA Marathon"

	Renderer.HTML(w, http.StatusOK, "article.html", data)
}
