package handlers

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type scheduleResponse struct {
	Schedule *schedule `json:"data,omitempty"`
}
type schedule struct {
	Name        string          `json:"name,omitempty"`
	Description string          `json:"description,omitempty"`
	Updated     string          `json:"updated,omitempty"`
	Link        string          `json:"link,omitempty"`
	Entries     []scheduleEntry `json:"items,omitempty"`
	Columns     []string        `json:"columns,omitempty"`
}

type scheduleEntry struct {
	Scheduled string           `json:"scheduled,omitempty"`
	Game      string           `json:"game,omitempty"`
	GameLink  string           `json:"game_link,omitempty"`
	Estimate  string           `json:"estimate,omitempty"`
	Players   []schedulePlayer `json:"players,omitempty"`
	Platform  string           `json:"platform,omitempty"`
	Category  string           `json:"category,omitempty"`
	Note      string           `json:"note,omitempty"`
	Data      []string         `json:"data,omitempty"`
	Length    float64          `json:"length_t,omitempty"`
}

type schedulePlayer struct {
	Name       string `json:"name,omitempty"`
	ProfileURL string `json:"profile_url,omitempty"`
}

// Schedule displays the marathon schedule
// TODO: The result of this method should be cached, so we don't have to parse the JSON every time.
func Schedule(w http.ResponseWriter, r *http.Request) {
	var s scheduleResponse
	data := getPagedata()

	// TODO: Should not be hard coded
	resp, err := http.Get("https://horaro.org/-/api/v1/schedules/4311u8b52b04si7a1e")
	if err != nil {
		log.Println(errors.Wrap(err, "handlers.Schedule"))
		renderer.HTML(w, http.StatusOK, "500.html", data)
	}

	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&s); err != nil {
		log.Println(errors.Wrap(err, "handlers.Schedule"))
		renderer.HTML(w, http.StatusOK, "500.html", data)
	}

	// Get all the indexes for the columns in order to identify them
	// on scheduleEntry.Data later
	columnIndexes := make(map[string]int)
	for i, c := range s.Schedule.Columns {
		columnIndexes[c] = i
	}

	// Go through each entry and attempt to set the correct values on the struct
	for i, e := range s.Schedule.Entries {
		if index, ok := columnIndexes["Game"]; ok {
			e.Game = getAnchorText(e.Data[index])
		}
		if index, ok := columnIndexes["Player(s)"]; ok {
			e.Players = getPlayers(e.Data[index])
		}
		if index, ok := columnIndexes["Platform"]; ok {
			e.Platform = e.Data[index]
		}
		if index, ok := columnIndexes["Category"]; ok {
			e.Category = e.Data[index]
		}
		if index, ok := columnIndexes["Note"]; ok {
			e.Note = e.Data[index]
		}
		e.Estimate = getEstimate(e.Length)

		s.Schedule.Entries[i] = e
	}
	data["Schedule"] = s.Schedule

	renderer.HTML(w, http.StatusOK, "schedule.html", data)
}

func getPlayers(str string) []schedulePlayer {
	nameRE := regexp.MustCompile("\\[([^]]+)\\]")
	names := nameRE.FindAllString(str, -1)

	linkRE := regexp.MustCompile("\\(([^]]+)\\)")
	links := linkRE.FindAllString(str, -1)

	players := make([]schedulePlayer, len(names))
	for i, name := range names {
		name = strings.Replace(name, "[", "", -1)
		name = strings.Replace(name, "]", "", -1)
		players[i].Name = name
	}

	for i, link := range links {
		link = strings.Replace(link, "(", "", -1)
		link = strings.Replace(link, ")", "", -1)
		players[i].ProfileURL = link
	}
	return players
}

func getAnchorText(str string) string {
	re := regexp.MustCompile("\\[([^]]+)\\]")
	match := re.FindStringSubmatch(str)
	if len(match) >= 1 {
		return match[1]
	}
	return str
}

func getAnchorLink(str string) string {
	re := regexp.MustCompile("\\(([^]]+)\\)")
	match := re.FindStringSubmatch(str)
	if len(match) >= 1 {
		return match[1]
	}
	return str
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
