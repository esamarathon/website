package handlers

import (
	"net/http"

	"github.com/dannyvankooten/grender"
)

type page struct {
	Meta    *meta
	Content *content
}

type content struct {
	Title string
	Body  string
}

type meta struct {
	Title       string
	Description string
	Image       string
}

var renderer = grender.New(grender.Options{
	TemplatesGlob: "templates/*.html",
})

// Returns index view
func Index(w http.ResponseWriter, r *http.Request) {
	m := meta{
		"ESA Marathon",
		"Welcome to European Speedrunner Assembly!",
		"http://www.esamarathon.com/images/esa/europeanspeedrunnerassembly.png",
	}
	c := content{
		"ESA Marathon",
		"Lorem ipsum",
	}

	p := page{
		&m,
		&c,
	}
	renderer.HTML(w, http.StatusOK, "index.html", p)
}

func Test(w http.ResponseWriter, r *http.Request) {

}
