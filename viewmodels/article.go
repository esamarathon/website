package viewmodels

import (
	"time"

	"github.com/esamarathon/website/config"
	"github.com/esamarathon/website/models/article"
	"github.com/esamarathon/website/models/menu"
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
		Layout:        layout{DefaultMeta, menu.Get()},
		CopyrightYear: time.Now().Year(),
		Livemode:      config.Config.LiveMode,
	}

	return view
}
