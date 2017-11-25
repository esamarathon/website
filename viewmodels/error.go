package viewmodels

import (
	"time"

	"github.com/olenedr/esamarathon/config"
)

type errorView struct {
	Meta          meta
	CopyrightYear int
	Livemode      bool
}

// Error returns the viewmodel for errorpages
func Error() errorView {
	view := errorView{
		Meta:          DefaultMeta,
		CopyrightYear: time.Now().Year(),
		Livemode:      config.Config.LiveMode,
	}

	return view
}
