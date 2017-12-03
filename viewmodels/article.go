package viewmodels

import (
	"time"

	"github.com/esamarathon/website/config"
	"github.com/esamarathon/website/models/article"
)

type articleView struct {
	Meta          meta
	Article       article.Article
	CopyrightYear int
	Livemode      bool
}

// Article returns the viewmodel for /news/{id}
func Article() articleView {
	view := articleView{
		Meta:          DefaultMeta,
		CopyrightYear: time.Now().Year(),
		Livemode:      config.Config.LiveMode,
	}

	return view
}
