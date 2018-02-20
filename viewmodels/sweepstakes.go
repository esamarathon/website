package viewmodels

import (
	"time"

	"github.com/esamarathon/website/config"
)

type sweepstakesView struct {
	Meta          meta
	CopyrightYear int
	Livemode      bool
}

// News returns the viewmodel for /news
func Sweepstakes() sweepstakesView {
	view := sweepstakesView{
		Meta: meta{
			Title:       "Sweepstake Rules - ESA Marathon",
			Description: "Rules for sweepstakes",
			Image:       DefaultMeta.Image,
		},
		CopyrightYear: time.Now().Year(),
		Livemode:      config.Config.LiveMode,
	}

	return view
}
