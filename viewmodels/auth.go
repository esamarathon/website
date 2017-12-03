package viewmodels

import (
	"time"

	"github.com/esamarathon/website/config"
)

type loginView struct {
	Meta          meta
	Livemode      bool
	CopyrightYear int
}

// Login returns the viewmodel for the loginview
func Login() loginView {
	view := loginView{
		Meta:          DefaultMeta,
		Livemode:      config.Config.LiveMode,
		CopyrightYear: time.Now().Year(),
	}

	return view
}
