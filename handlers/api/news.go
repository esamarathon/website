package api

import (
	"net/http"
	"strconv"

	"github.com/esamarathon/website/handlers/helpers"
	"github.com/esamarathon/website/models/article"
)

// News serves a JSON document of all published news articles.
// The amount returned can be controlled with the "limit" HTTP GET parameter.
func News(w http.ResponseWriter, r *http.Request) {
	articles, err := article.AllPublished()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	limitStrings, ok := r.URL.Query()["limit"]

	if ok {
		limit, err := strconv.Atoi(limitStrings[0])
		if err != nil {
			http.Error(w, "Invalid limit parameter. Must be not set or positive integer.", http.StatusBadRequest)
			return
		}

		if limit > 0 {
			articles = articles[0:limit]
		}
	}

	// Reduce body to a teaser
	for i, a := range articles {
		a.ParseHTML()
		a.FormatTimestamp()
		articles[i] = a
	}

	helpers.Renderer.JSON(w, http.StatusOK, articles)
}
