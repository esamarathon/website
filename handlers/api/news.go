package api

import (
	"net/http"

	"github.com/esamarathon/website/handlers/helpers"
	"github.com/esamarathon/website/models/article"
)

func News(w http.ResponseWriter, r *http.Request) {
	articles := article.All()

	amount, ok := parseInt(r.URL.Query()["amount"])

	if limit > 0 {
		articles = articles[0:limit]
	}

	// Reduce body to a teaser
	for i, a := range articles {
		a.ParseHTML()
		a.FormatTimestamp()
		articles[i] = a
	}

	helpers.Renderer.JSON(w, http.StatusOK, articles)
}

func parseInt(s string, ok bool) (int, bool) {
	if !ok {
		return -1, false
	}

	var i, err := strconv.Atoi()
	if err != nil {
		return -1, false
	}

	return i, true
}
