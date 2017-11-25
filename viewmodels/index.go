package viewmodels

import (
	"html/template"
	"time"

	"github.com/olenedr/esamarathon/config"
)

type meta struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Image       string `json:"image,omitempty"`
}

// DefaultMata is a set of default metadata values
var DefaultMeta = meta{
	"ESA Marathon",
	"Welcome to European Speedrunner Assembly!",
	"http://www.esamarathon.com/images/esa/europeanspeedrunnerassembly.png",
}

type indexView struct {
	Meta          meta
	Title         string
	Body          template.HTML
	Livemode      bool
	CopyrightYear int
}

// Index returns the viewmodel for the indexview
func Index() indexView {
	// TODO: Should return frontpage data from DB or file
	view := indexView{
		Meta:          DefaultMeta,
		Title:         "Welcome to European Speedrunner Assembly!",
		Body:          "",
		Livemode:      config.Config.LiveMode,
		CopyrightYear: time.Now().Year(),
	}

	return view
}

// GetPagedata returns the basic page data
func GetPagedata() map[string]interface{} {
	s := config.Config.LiveMode
	t := time.Now()

	p := map[string]interface{}{
		"Meta":          DefaultMeta,
		"Livemode":      s,
		"CopyrightYear": t.Year(),
	}
	return p
}