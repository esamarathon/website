package admin

import (
	"log"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	"github.com/esamarathon/website/models/user"
	"github.com/esamarathon/website/viewmodels"

	"github.com/gorilla/mux"
)

/*
*	User routes
 */
// List all the users
func userIndex(w http.ResponseWriter, r *http.Request) {
	adminRenderer.HTML(w, http.StatusOK, "user.html", viewmodels.AdminUserIndex(w, r))
}

// Store the user in the database
func userStore(w http.ResponseWriter, r *http.Request) {
	// Parse form and get the username submitted
	r.ParseForm()
	// Username is lowercase since that's what Twitch
	// returns through their Oauth response
	username := strings.ToLower(r.Form.Get("username"))
	// Create a user object
	u := user.User{
		Username: username,
	}

	exists, err := u.Exists()
	if err != nil {
		user.SetFlashMessage(w, r, "alert", "Something went wrong.")
		log.Println(errors.Wrap(err, "handlers.userStore"))
		http.Redirect(w, r, "/admin/user", http.StatusBadRequest)
		return
	}

	if exists {
		user.SetFlashMessage(w, r, "alert", "User already exists.")
		http.Redirect(w, r, "/admin/user", http.StatusFound)
		return
	}

	_, err = user.Create(username)
	if err != nil || username == "" {
		user.SetFlashMessage(w, r, "alert", "Failed to add user to database.")
		log.Println(errors.Wrap(err, "handlers.userStore"))
		http.Redirect(w, r, "/admin/user", http.StatusBadRequest)
		return
	}

	user.SetFlashMessage(w, r, "success", "User was added.")
	http.Redirect(w, r, "/admin/user", http.StatusSeeOther)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if err := user.Delete(id); err != nil {
		user.SetFlashMessage(w, r, "alert", "Something went wrong while trying to delete the user.")
		log.Println(errors.Wrap(err, "handlers.deleteUser"))
	} else {
		user.SetFlashMessage(w, r, "success", "The user was deleted")
	}

	http.Redirect(w, r, "/admin/user", http.StatusSeeOther)
}