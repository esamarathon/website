package viewmodels

import (
	"time"

	"github.com/esamarathon/website/config"
	"github.com/esamarathon/website/models/article"
	"github.com/esamarathon/website/models/menu"
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
		Layout:        layout{DefaultMeta, menu.Get()},
		CopyrightYear: time.Now().Year(),
		Livemode:      config.Config.LiveMode,
	}

	return view
}
