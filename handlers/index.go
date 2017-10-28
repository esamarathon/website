package handlers

import (
	"net/http"

	"github.com/dannyvankooten/grender"
	"github.com/olenedr/esamarathon/models/setting"
)

var renderer = grender.New(grender.Options{
	TemplatesGlob: "templates/*.html",
	PartialsGlob:  "templates/partials/*.html",
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

// Index returns index view
func Index(w http.ResponseWriter, r *http.Request) {
	data := getPagedata()
	renderer.HTML(w, http.StatusOK, "index.html", data)
}

// getPagedata returns the basic page data
func getPagedata() map[string]interface{} {
	s, err := setting.GetLiveMode().AsBool()
	if err != nil {
		s = false
	}

	p := map[string]interface{}{
		"Meta":     Meta,
		"Content":  Content,
		"Livemode": s,
	}
	return p
}
