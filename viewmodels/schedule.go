package viewmodels

import (
	"time"

	"github.com/olenedr/esamarathon/cache"
	"github.com/olenedr/esamarathon/config"
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
	metadata := meta{
		Title:       "ESA Schedule",
		Description: "See the schedule for the next great ESA event!",
		Image:       DefaultMeta.Image,
	}
	view := scheduleView{
		Meta:          metadata,
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
