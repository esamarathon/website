package viewmodels

import (
	"time"

	"github.com/esamarathon/website/cache"
	"github.com/esamarathon/website/config"
	"github.com/esamarathon/website/models/menu"
)

type scheduleView struct {
	Layout        layout
	Schedule      interface{}
	Cached        bool
	CopyrightYear int
	Livemode      bool
}

// Schedule returns the viewmodel for /schedule
func Schedule() scheduleView {
	view := scheduleView{
		Layout:        layout{DefaultMeta, menu.Get()},
		Livemode:      config.Config.LiveMode,
		CopyrightYear: time.Now().Year(),
	}

	// Attempt to find a cached schedule
	schedule, found := cache.Get("schedule")
	view.Cached = found
	if view.Cached {
		view.Schedule = schedule
	}

	return view
}
