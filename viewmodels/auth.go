package viewmodels

import (
	"time"

	"github.com/esamarathon/website/config"
)

type loginView struct {
	Layout        layout
	Livemode      bool
	CopyrightYear int
}

// Login returns the viewmodel for the loginview
func Login() loginView {
	view := loginView{
		Layout:        DefaultLayout(),
		Livemode:      config.Config.LiveMode,
		CopyrightYear: time.Now().Year(),
	}

	return view
}
