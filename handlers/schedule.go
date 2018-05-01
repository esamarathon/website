package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"math"
	"net/http"
	"strconv"

	. "github.com/esamarathon/website/handlers/helpers"
	"github.com/esamarathon/website/cache"
	"github.com/esamarathon/website/config"
	"github.com/esamarathon/website/models/schedule"
	"github.com/esamarathon/website/viewmodels"
	"github.com/pkg/errors"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

type scheduleResponse struct {
	Schedule *schedule.Schedule `json:"data,omitempty"`
}

// Schedule displays the marathon schedule
func Schedule(w http.ResponseWriter, r *http.Request) {
	view := viewmodels.Schedule()

	if view.Cached {
		// Render cached view
		Renderer.HTML(w, http.StatusOK, "schedule.html", view)
		return
	}
	var s scheduleResponse

	// Request the schedule JSON-resource
	resp, err := http.Get(config.Config.ScheduleAPIURL)

	// If something goes wrong, we return the 500-view
	if err != nil {
		log.Println(errors.Wrap(err, "handlers.Schedule"))
		HandleInternalError(w)
		return
	}

	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&s); err != nil {
		log.Println(errors.Wrap(err, "handlers.Schedule"))
		HandleInternalError(w)
		return
	}

	// Get all the indexes for the columns in order to identify them
	// on scheduleEntry.Data later
	columnIndexes := make(map[string]int)
	for i, c := range s.Schedule.Columns {
		columnIndexes[c] = i
	}

	// Go through each entry and attempt to set the correct values on the struct
	// Added some old formats indexes for backwards compatibility for good measure
	for i, e := range s.Schedule.Entries {
		if index, ok := columnIndexes["Game"]; ok {
			e.Game = getHTML(e.Data[index])
		}
		if index, ok := columnIndexes["Runner"]; ok {
			e.Players = getHTML(e.Data[index])
		}
		if index, ok := columnIndexes["Runner/Runners"]; ok {
			e.Players = getHTML(e.Data[index])
		}
		if index, ok := columnIndexes["Player(s)"]; ok {
			e.Players = getHTML(e.Data[index])
		}
		if index, ok := columnIndexes["Platform"]; ok {
			e.Platform = e.Data[index]
		}
		if index, ok := columnIndexes["Console"]; ok {
			e.Platform = e.Data[index]
		}
		if index, ok := columnIndexes["Category"]; ok {
			e.Category = e.Data[index]
		}
		if index, ok := columnIndexes["Region"]; ok {
			e.Note = e.Data[index]
		}
		if index, ok := columnIndexes["Note"]; ok {
			e.Note = e.Data[index]
		}
		e.Estimate = getEstimate(e.Length)

		s.Schedule.Entries[i] = e
	}
	// Attach the schedule
	view.Schedule = s.Schedule

	// Render
	Renderer.HTML(w, http.StatusOK, "schedule.html", view)
	// Write to the cache
	cache.Cache.Set("schedule", s.Schedule, cache.Duration())
}

// returns HTML based a on markdown string
func getHTML(str string) template.HTML {
	markdown := string(blackfriday.Run([]byte(str)))
	return template.HTML(markdown)
}

// getEstimate returns a formated string representing
// the estimated time of a speedrun in hours:minutes
func getEstimate(length float64) string {
	// Convert length to hours
	hours := math.Floor(length / 3600)
	// Convert length to minutes
	minutes := (int(length) % 3600) / 60

	// Convert to strings and add leading zeros
	var strMinutes, strHours string
	if hours < 10 {
		strHours = "0" + strconv.FormatFloat(hours, 'f', -1, 64)
	} else {
		strHours = strconv.FormatFloat(hours, 'f', -1, 64)
	}

	if minutes < 10 {
		strMinutes = "0" + strconv.Itoa(minutes)
	} else {
		strMinutes = strconv.Itoa(minutes)
	}

	// Return string formated estimate
	return strHours + ":" + strMinutes
}
