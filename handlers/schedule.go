package handlers

import (
	"log"
	"net/http"
	"sync"

	"github.com/esamarathon/website/cache"
	. "github.com/esamarathon/website/handlers/helpers"
	"github.com/esamarathon/website/horaro"
	"github.com/esamarathon/website/models/schedule"
	"github.com/esamarathon/website/viewmodels"
	"github.com/pkg/errors"
)

// Schedule displays the marathon schedule
func Schedule(w http.ResponseWriter, r *http.Request) {
	view := viewmodels.Schedule()

	scheduleRefs, err := schedule.All()

	if err != nil || len(scheduleRefs) == 0 {
		NoSchedule(w, r)
		return
	}

	if view.Cached {
		// Render cached view
		Renderer.HTML(w, http.StatusOK, "schedule.html", view)
		return
	}

	ss := getSchedules(scheduleRefs)

	// Attach the schedule
	view.Schedules = ss

	// Render
	Renderer.HTML(w, http.StatusOK, "schedule.html", view)
	// Write to the cache
	cache.Set("schedules", ss, cache.Duration())
}

func NoSchedule(w http.ResponseWriter, r *http.Request) {
	err := Renderer.HTML(w, http.StatusOK, "noschedule.html", viewmodels.NoSchedule())
	if err != nil {
		log.Println(errors.Wrap(err, "handlers.NoSchedule"))
	}
}

func getSchedules(refs []schedule.ScheduleRef) (ss []schedule.Schedule) {
	var wg sync.WaitGroup
	wg.Add(len(refs))
	ss = make([]schedule.Schedule, len(refs), len(refs))

	for i, ref := range refs {
		go func(ref schedule.ScheduleRef, place int) {
			defer wg.Done()

			s, err := horaro.GetSchedule(ref.Url)
			if err != nil {
				return
			}
			s.ID = ref.ID
			s.Name = ref.Title
			ss[place] = *s
		}(ref, i)
	}

	wg.Wait()

	return
}
