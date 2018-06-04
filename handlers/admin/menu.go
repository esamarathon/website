package admin

import (
	"net/http"

	"github.com/esamarathon/website/models/menu"
	"github.com/esamarathon/website/models/user"
	"github.com/esamarathon/website/viewmodels"

	"github.com/gorilla/mux"
)

func menuIndex(w http.ResponseWriter, r *http.Request) {
	if !menu.IsStored() {
		m := menu.Get()
		err := m.Insert()
		if len(err) > 0 {
			user.SetFlashMessage(w, r, "alert", "Couldn't get data from DB. There might be connection issues or the table might not exist!")
		}
	}
	adminRenderer.HTML(w, http.StatusOK, "menu.html", viewmodels.AdminMenuIndex(w, r))
}

func menuUpdate(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	m, err := menu.Find(id)
	if err != nil {
		user.SetFlashMessage(w, r, "alert", "Couldn't find the menu item you wanted to update")
		http.Redirect(w, r, "/admin/menu", http.StatusSeeOther)
		return
	}

	r.ParseForm()
	m.Title = r.Form.Get("title")
	m.Link = r.Form.Get("link")
	if r.Form.Get("new_tab") == "true" {
		m.NewTab = true
	} else {
		m.NewTab = false
	}
	err = m.Update()
	if err != nil {
		user.SetFlashMessage(w, r, "alert", "Something went wrong while trying to update")
		http.Redirect(w, r, "/admin/menu", http.StatusSeeOther)
		return
	}

	user.SetFlashMessage(w, r, "success", "The menu was updated")
	http.Redirect(w, r, "/admin/menu", http.StatusSeeOther)
}
