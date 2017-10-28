package handlers

import "net/http"

func News(w http.ResponseWriter, r *http.Request) {
	data := getPagedata()
	renderer.HTML(w, http.StatusOK, "news.html", data)
}
