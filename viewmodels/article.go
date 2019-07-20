package viewmodels

import (
	"time"

	"github.com/esamarathon/website/config"
	"github.com/esamarathon/website/models/article"
)

type articleView struct {
	Layout        layout
	Article       article.Article
	CopyrightYear int
	Livemode      bool
}

// Article returns the viewmodel for /news/{id}
func Article() articleView {
	view := articleView{
		Layout:        DefaultLayout(),
		CopyrightYear: time.Now().Year(),
		Livemode:      config.Config.LiveMode,
	}

	return view
}
