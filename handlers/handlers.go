package handlers

import (
	"net/http"

	"github.com/dannyvankooten/grender"
)

var renderer = grender.New(grender.Options{
	TemplatesGlob: "templates/*.html",
})

var m = meta{
	"ESA Marathon",
	"Welcome to European Speedrunner Assembly!",
	"http://www.esamarathon.com/images/esa/europeanspeedrunnerassembly.png",
}
var c = content{
	"Welcome to European Speedrunner Assembly!",
	"",
}
var p = page{
	m,
	c,
}

// Index returns index view
func Index(w http.ResponseWriter, r *http.Request) {
	renderer.HTML(w, http.StatusOK, "index.html", p)
}

// Test is a test handler for Ole to debug stuff
func Test(w http.ResponseWriter, r *http.Request) {

}
