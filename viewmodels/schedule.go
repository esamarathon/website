package viewmodels

import (
	"time"

	"github.com/esamarathon/website/cache"
	"github.com/esamarathon/website/config"
)

type scheduleView struct {
	Meta          meta
	Schedule      interface{}
	Cached        bool
	CopyrightYear int
	Livemode      bool
}

// Schedule returns the viewmodel for /schedule
func Schedule() scheduleView {
	view := scheduleView{
		Meta: meta{
			Title:       "Schedule - ESA Marathon",
			Description: "Check out the schedule for the next great ESA event!",
			Image:       DefaultMeta.Image,
		},
		CopyrightYear: time.Now().Year(),
		Livemode:      config.Config.LiveMode,
	}

	// Attempt to find a cached schedule
	schedule, found := cache.Get("schedule")
	view.Cached = found
	if view.Cached {
		view.Schedule = schedule
	}

	return view
}
