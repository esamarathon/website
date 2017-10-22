package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	blackfriday "gopkg.in/russross/blackfriday.v2"
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
	Scheduled string   `json:"scheduled,omitempty"`
	Game      string   `json:"game,omitempty"`
	Estimage  string   `json:"estimage,omitempty"`
	Players   string   `json:"players,omitempty"`
	Platform  string   `json:"platform,omitempty"`
	Category  string   `json:"category,omitempty"`
	Note      string   `json:"note,omitempty"`
	Data      []string `json:"data,omitempty"`
}

func Schedule(w http.ResponseWriter, r *http.Request) {
	var s scheduleResponse

	resp, err := http.Get("https://horaro.org/-/api/v1/schedules/4311u8b52b04si7a1e")
	if err != nil {
		fmt.Printf("%v", err)
		renderer.HTML(w, http.StatusOK, "500.html", page{m, content{}})
	}

	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&s); err != nil {
		fmt.Printf("%v", err)
		renderer.HTML(w, http.StatusOK, "500.html", page{m, content{}})
	}

	for i, e := range s.Schedule.Entries {
		o := blackfriday.Run([]byte(e.Data[0]))
		e.Game = string(o)
		o = blackfriday.Run([]byte(e.Data[1]))
		e.Players = string(o)

		s.Schedule.Entries[i] = e
	}
	s.Schedule.Meta = m
	renderer.HTML(w, http.StatusOK, "schedule.html", s.Schedule)
}
