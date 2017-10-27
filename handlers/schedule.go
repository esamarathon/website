package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
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
	Meta        meta            `json:"data,omitempty"`
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
	Length    int              `json:"length_t,omitempty"`
}

type schedulePlayer struct {
	Name       string `json:"name,omitempty"`
	ProfileURL string `json:"profile_url,omitempty"`
}

func Schedule(w http.ResponseWriter, r *http.Request) {
	var s scheduleResponse

	resp, err := http.Get("https://horaro.org/-/api/v1/schedules/4311u8b52b04si7a1e")
	if err != nil {
		log.Println(errors.Wrap(err, "handlers.Schedule"))
		renderer.HTML(w, http.StatusOK, "500.html", page{Meta, content{}})
	}

	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&s); err != nil {
		log.Println(errors.Wrap(err, "handlers.Schedule"))
		renderer.HTML(w, http.StatusOK, "500.html", page{Meta, content{}})
	}

	for i, e := range s.Schedule.Entries {
		log.Println("Original: " + e.Data[1])
		e.Game = getAnchorText(e.Data[0])
		e.Players = getPlayers(e.Data[1])
		e.Platform = e.Data[2]
		e.Note = e.Data[3]
		// e.Estimate = getEstimate(e.Length)

		s.Schedule.Entries[i] = e
	}
	s.Schedule.Meta = Meta
	renderer.HTML(w, http.StatusOK, "schedule.html", s.Schedule)
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

func getEstimate(length int) string {
	// TODO:implement this function
	return ""
}
