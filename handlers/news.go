package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/olenedr/esamarathon/models/article"
)

// News renders the news page
func News(w http.ResponseWriter, r *http.Request) {
	data := getPagedata()
	// Default page number
	page := "0"

	// If a page is specified we use that instead
	if query := r.URL.Query()["page"]; len(query) != 0 {
		page = query[0]
	}
	p, err := strconv.Atoi(page)

	// If we failed to get the page number
	// we just set it to 0 (first page)
	if err != nil {
		p = 0
	}
	articles, err := article.Page(p)
	// If we failed to get the articles
	// we return the 500 error page
	if err != nil {
		renderer.HTML(w, http.StatusOK, "500.html", data)
		return
	}

	// Reduce body to a teaser
	for i, a := range articles {
		a.ShortenBody()
		a.ParseHTML()
		articles[i] = a
	}

	data["Articles"] = articles
	data["NextPage"] = p + 1
	data["PrevPage"] = p - 1
	data["LastPage"], err = article.PageCount()

	if err != nil {
		renderer.HTML(w, http.StatusOK, "500.html", data)
		return
	}

	renderer.HTML(w, http.StatusOK, "news.html", data)
}

// Article renders the page of a specific article
func Article(w http.ResponseWriter, r *http.Request) {
	data := getPagedata()
	id := mux.Vars(r)["id"]
	a, err := article.Get(id)

	if err != nil {
		renderer.HTML(w, http.StatusOK, "404.html", data)
		return
	}

	a.ParseHTML()
	a.FormatTimestamp()
	data["Article"] = a

	renderer.HTML(w, http.StatusOK, "article.html", data)
}
