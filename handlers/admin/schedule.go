package admin

import (
	"github.com/esamarathon/website/cache"
	"github.com/esamarathon/website/config"
	"github.com/esamarathon/website/models/schedule"
	"github.com/esamarathon/website/models/user"
	"github.com/esamarathon/website/viewmodels"

	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

const scheduleBaseRoute = "/admin/schedule"

/*
*	Admin Schedule Index route
 */
func scheduleIndex(w http.ResponseWriter, r *http.Request) {
	viewmodel := viewmodels.AdminSchedules(w, r)

	scheds, err := schedule.All()
	if err != nil {
		user.SetFlashMessage(w, r, "alert", "Unable to load schedules.")
		http.Redirect(w, r, scheduleBaseRoute, http.StatusTemporaryRedirect)
		return
	}
	viewmodel.Schedules = scheds

	adminRenderer.HTML(w, http.StatusOK, "schedule.html", viewmodel)
}

func toggleShowSchedule(w http.ResponseWriter, r *http.Request) {
	config.ToggleShowSchedule()
	http.Redirect(w, r, scheduleBaseRoute, http.StatusTemporaryRedirect)
}

func scheduleCreate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	url := r.Form.Get("url")
	title := r.Form.Get("title")
	order, err := strconv.Atoi(r.Form.Get("order"))
	if err != nil {
		order = 99
	}

	validateUrl(w, r, url)

	model := schedule.ScheduleRef{
		Url:   url,
		Title: title,
		Order: order,
	}

	err = model.Create()
	if err != nil {
		user.SetFlashMessage(w, r, "alert", "Unable to create new schedule.")
		log.Println("Error creating new schedule.", err)
		http.Redirect(w, r, scheduleBaseRoute, http.StatusSeeOther)
		return
	}

	cache.Clear("schedules")
	user.SetFlashMessage(w, r, "success", "New schedule has been created!")
	http.Redirect(w, r, scheduleBaseRoute, http.StatusSeeOther)
}

// updateSchedule parses a form and updates the ScheduleAPIURL
// if the new URL seems valid
func scheduleUpdate(w http.ResponseWriter, r *http.Request) {
	// Parse form and get the submitted URL
	r.ParseForm()
	id := mux.Vars(r)["id"]
	url := r.Form.Get("url")
	title := r.Form.Get("title")
	order, err := strconv.Atoi(r.Form.Get("order"))
	if err != nil {
		order = 99
	}

	validateUrl(w, r, url)

	// URL seems fine, updating
	model, err := schedule.Get(id)
	if err != nil {
		user.SetFlashMessage(w, r, "alert", "Schedule was not found in database.")
		http.Redirect(w, r, scheduleBaseRoute, http.StatusSeeOther)
		return
	}
	model.Url = url
	model.Title = title
	model.Order = order
	model.Update()

	cache.Clear("schedules")
	user.SetFlashMessage(w, r, "success", "Schedule URL has been updated!")
	http.Redirect(w, r, scheduleBaseRoute, http.StatusSeeOther)
}

func scheduleDelete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	schedule.Delete(id)
	cache.Clear("schedules")
	http.Redirect(w, r, scheduleBaseRoute, http.StatusSeeOther)
}

func validateUrl(w http.ResponseWriter, r *http.Request, url string) {
	// Validate URL
	if !strings.Contains(url, "https://horaro.org/-/api/v1/schedules/") {
		user.SetFlashMessage(w, r, "alert", "Not a valid Horaro API URL. Not updating. Correct format is \"https://horaro.org/-/api/v1/schedules/\"")
		http.Redirect(w, r, scheduleBaseRoute, http.StatusSeeOther)
		return
	}

	// Attempt to get the resource
	resp, err := http.Get(url)
	if err != nil {
		user.SetFlashMessage(w, r, "alert", "Request to resource failed, not updating.")
		http.Redirect(w, r, scheduleBaseRoute, http.StatusSeeOther)
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		user.SetFlashMessage(w, r, "alert", "Request to resource failed, not updating.")
		http.Redirect(w, r, scheduleBaseRoute, http.StatusSeeOther)
		return
	}
}
