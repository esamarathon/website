package viewmodels

import (
	"time"

	"github.com/esamarathon/website/config"
)

type errorView struct {
	Layout        layout
	CopyrightYear int
	Livemode      bool
}

// Error returns the viewmodel for errorpages
func Error() errorView {
	view := errorView{
		Layout:        DefaultLayout(),
		CopyrightYear: time.Now().Year(),
		Livemode:      config.Config.LiveMode,
	}

	return view
}
