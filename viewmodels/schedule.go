package viewmodels

import (
	"time"

	"github.com/esamarathon/website/cache"
	"github.com/esamarathon/website/config"
	"github.com/esamarathon/website/models/schedule"
)

type scheduleView struct {
	Layout        layout
	Schedules     []schedule.Schedule
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
	schedules, found := cache.Get("schedules")
	view.Cached = found
	if view.Cached {
		view.Schedules = schedules.([]schedule.Schedule)
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
