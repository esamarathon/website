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
	a, err := article.Page(p)
	// If we failed to get the articles
	// we return the 500 error page
	if err != nil {
		renderer.HTML(w, http.StatusOK, "500.html", data)
		return
	}

	data["Articles"] = a
	data["CurrPage"] = p
	data["LastPage"], err = article.PageCount()
	fmt.Printf("LAST PAGE: %v\n", data["LastPage"])
	fmt.Printf("CURR PAGE: %v\n", data["CurrPage"])
	if err != nil {
		renderer.HTML(w, http.StatusOK, "500.html", data)
		return
	}

	renderer.HTML(w, http.StatusOK, "news.html", data)
}

func Article(w http.ResponseWriter, r *http.Request) {
	data := getPagedata()
	renderer.HTML(w, http.StatusOK, "article.html", data)
}
