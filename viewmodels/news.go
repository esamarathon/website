package viewmodels

import (
	"time"

	"github.com/olenedr/esamarathon/config"
	"github.com/olenedr/esamarathon/models/article"
)

type newsView struct {
	Meta          meta
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
		Meta:          DefaultMeta,
		CopyrightYear: time.Now().Year(),
		Livemode:      config.Config.LiveMode,
	}
	view.Meta.Title = "News - ESA Marathon"
	view.Meta.Description = "We constantly update with news about our events."

	return view
}
