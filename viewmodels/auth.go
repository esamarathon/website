package viewmodels

import (
	"time"

	"github.com/esamarathon/website/config"
	"github.com/esamarathon/website/models/menu"
)

type loginView struct {
	Layout        layout
	Livemode      bool
	CopyrightYear int
}

// Login returns the viewmodel for the loginview
func Login() loginView {
	view := loginView{
		Layout:        layout{DefaultMeta, menu.Get()},
		Livemode:      config.Config.LiveMode,
		CopyrightYear: time.Now().Year(),
	}

	return view
}
