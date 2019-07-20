package viewmodels

import (
	"time"

	"github.com/esamarathon/website/config"
	"github.com/esamarathon/website/models/article"
)

type newsView struct {
	Layout        layout
	Articles      []article.Article
	NextPage      int
	PrevPage      int
	LastPage      int
	CopyrightYear int
	Livemode      bool
}

// News returns the viewmodel for /news
func News() newsView {
	view := newsView{
		Layout:        DefaultLayout(),
		CopyrightYear: time.Now().Year(),
		Livemode:      config.Config.LiveMode,
	}

	return view
}
