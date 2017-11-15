package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/olenedr/esamarathon/models/article"
)

func News(w http.ResponseWriter, r *http.Request) {
	data := getPagedata()
	page := "0"

	if query := r.URL.Query()["page"]; len(query) != 0 {
		page = query[0]
	}
	p, err := strconv.Atoi(page)

	// If we failed to get the page number
	// we just set it to 0 (first page)
	if err != nil {
		p = 0
	}
	// fmt.Printf("%v\n", p)
	a, err := article.Page(p)

	if err != nil {
		renderer.HTML(w, http.StatusOK, "404.html", data)
		return
	}

	fmt.Printf("%v\n", a)
	renderer.HTML(w, http.StatusOK, "news.html", data)
}

func Article(w http.ResponseWriter, r *http.Request) {
	data := getPagedata()
	renderer.HTML(w, http.StatusOK, "article.html", data)
}
