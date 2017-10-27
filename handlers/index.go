package handlers

import (
	"net/http"

	"github.com/dannyvankooten/grender"
	"github.com/olenedr/esamarathon/models/setting"
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

var Page = map[string]interface{}{
	"Meta":    Meta,
	"Content": Content,
}

// Index returns index view
func Index(w http.ResponseWriter, r *http.Request) {
	s, err := setting.GetLiveMode().AsBool()
	if err == nil {
		Page["Livemode"] = s
	}
	renderer.HTML(w, http.StatusOK, "index.html", Page)
}
