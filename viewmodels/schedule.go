package viewmodels

import (
	"time"

	"github.com/esamarathon/website/cache"
	"github.com/esamarathon/website/config"
)

type scheduleView struct {
	Layout        layout
	Schedule      interface{}
	Cached        bool
	CopyrightYear int
	Livemode      bool
}

type noscheduleView struct {
	Layout        layout
	CopyrightYear int
	Livemode      bool
}

// Schedule returns the viewmodel for /schedule
func Schedule() scheduleView {
	view := scheduleView{
		Layout:        DefaultLayout(),
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

func NoSchedule() noscheduleView {
	return noscheduleView{
		Layout:        DefaultLayout(),
		Livemode:      config.Config.LiveMode,
		CopyrightYear: time.Now().Year(),
	}
}
