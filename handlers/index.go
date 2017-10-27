package handlers

import (
	"net/http"

	"github.com/dannyvankooten/grender"
)

var renderer = grender.New(grender.Options{
	TemplatesGlob: "templates/*.html",
})

var Meta = meta{
	"ESA Marathon",
	"Welcome to European Speedrunner Assembly!",
	"http://www.esamarathon.com/images/esa/europeanspeedrunnerassembly.png",
}
var Content = content{
	"Welcome to European Speedrunner Assembly!",
	"",
}
var Page = page{
	Meta,
	Content,
}

// Index returns index view
func Index(w http.ResponseWriter, r *http.Request) {
	renderer.HTML(w, http.StatusOK, "index.html", Page)
}
