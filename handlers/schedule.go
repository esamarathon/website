package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	blackfriday "gopkg.in/russross/blackfriday.v2"
)

type scheduleResponse struct {
	Schedule schedule `json:"data,omitempty"`
}
type schedule struct {
	Name        string          `json:"name,omitempty"`
	Description string          `json:"description,omitempty"`
	Updated     string          `json:"updated,omitempty"`
	Link        string          `json:"link,omitempty"`
	Entries     []scheduleEntry `json:"items,omitempty"`
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

	resp, _ := http.Get("https://horaro.org/-/api/v1/schedules/4311u8b52b04si7a1e")

	defer resp.Body.Close()

	_ = json.NewDecoder(resp.Body).Decode(&s)

	for _, e := range s.Schedule.Entries {
		o := blackfriday.Run([]byte(e.Data[0]))
		e.Game = string(o)
		o = blackfriday.Run([]byte(e.Data[1]))
		e.Players = string(o)

		// fmt.Printf("Game: %v\n", e.Data[0])

		// fmt.Printf("Runner: %v\n", e.Data[1])
	}
	fmt.Println(s.Schedule.Name)
	renderer.HTML(w, http.StatusOK, "schedule.html", s)
}
