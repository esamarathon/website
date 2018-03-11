package viewmodels

import (
	"time"

	"github.com/esamarathon/website/config"
	"github.com/esamarathon/website/models/menu"
)

type errorView struct {
	Layout        layout
	CopyrightYear int
	Livemode      bool
}

// Error returns the viewmodel for errorpages
func Error() errorView {
	view := errorView{
		Layout:        layout{DefaultMeta, menu.Get()},
		CopyrightYear: time.Now().Year(),
		Livemode:      config.Config.LiveMode,
	}

	return view
}
