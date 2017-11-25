package handlers

import (
	"log"
	"net/http"
	"strings"

	"github.com/dannyvankooten/grender"
	"github.com/pkg/errors"

	"github.com/olenedr/esamarathon/cache"
	"github.com/olenedr/esamarathon/config"
	"github.com/olenedr/esamarathon/models/article"
	"github.com/olenedr/esamarathon/models/user"

	"github.com/gorilla/mux"
	"github.com/olenedr/esamarathon/middleware"
)

// AdminRoutes adds the admin routes to the router
func AdminRoutes(base string, router *mux.Router) {
	requireAuth := middleware.AuthMiddleware
	router.HandleFunc(base, requireAuth(indexAdmin)).Methods("GET", "POST")
	router.HandleFunc(base+"/toggle", requireAuth(toggleLivemode)).Methods("GET")
	router.HandleFunc(base+"/schedule", requireAuth(updateSchedule)).Methods("POST")
	router.HandleFunc(base+"/user", requireAuth(userIndex)).Methods("GET")
	router.HandleFunc(base+"/user", requireAuth(userStore)).Methods("POST")
	router.HandleFunc(base+"/user/{id}/delete", requireAuth(deleteUser)).Methods("GET")
	router.HandleFunc(base+"/article", requireAuth(articleIndex)).Methods("GET")
	router.HandleFunc(base+"/article/create", requireAuth(articleCreate)).Methods("GET")
	router.HandleFunc(base+"/article/create", requireAuth(articleStore)).Methods("POST")
	router.HandleFunc(base+"/article/{id}", requireAuth(articleEdit)).Methods("GET")
	router.HandleFunc(base+"/article/{id}", requireAuth(articleUpdate)).Methods("POST")
	router.HandleFunc(base+"/article/{id}/delete", requireAuth(articleDelete)).Methods("GET")
}

// Boots a renderer for the admin views
var adminRenderer = grender.New(grender.Options{
	TemplatesGlob: "templates_admin/*.html",
	PartialsGlob:  "templates_admin/partials/*.html",
})

/*
*	Admin Index routes
 */
func indexAdmin(w http.ResponseWriter, r *http.Request) {
	// Change with actual status from DB
	u, userErr := user.FromSession(r)
	if userErr != nil {
		log.Println(errors.Wrap(userErr, "admin.index"))
	}
	data := map[string]interface{}{
		"User":           u,
		"Status":         config.Config.LiveMode,
		"ScheduleAPIURL": config.Config.ScheduleAPIURL,
		"Alert":          user.GetFlashMessage(w, r, "alert"),
		"Success":        user.GetFlashMessage(w, r, "success"),
	}

	adminRenderer.HTML(w, http.StatusOK, "index.html", data)
}

// Toggles the stream on the frontpage
func toggleLivemode(w http.ResponseWriter, r *http.Request) {
	config.ToggleLiveMode()
	http.Redirect(w, r, "/admin", http.StatusTemporaryRedirect)
}

// updateSchedule parses a form and updates the ScheduleAPIURL
// if the new URL seems valid
func updateSchedule(w http.ResponseWriter, r *http.Request) {
	// Parse form and get the submitted URL
	r.ParseForm()
	URL := r.Form.Get("url")

	// Validate URL
	if !strings.Contains(URL, "https://horaro.org/-/api/v1/schedules/") {
		user.SetFlashMessage(w, r, "alert", "Not a valid Horaro API URL. Not updating. Correct format is \"https://horaro.org/-/api/v1/schedules/\"")
		http.Redirect(w, r, "/admin", http.StatusTemporaryRedirect)
		return
	}

	// Attempt to get the resource
	resp, err := http.Get(URL)
	if err != nil {
		user.SetFlashMessage(w, r, "alert", "Request to resource failed, not updating.")
		http.Redirect(w, r, "/admin", http.StatusTemporaryRedirect)
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		user.SetFlashMessage(w, r, "alert", "Request to resource failed, not updating.")
		http.Redirect(w, r, "/admin", http.StatusTemporaryRedirect)
		return
	}
	cache.Cache.Delete("schedule")

	// URL seems fine, updating
	config.Config.ScheduleAPIURL = URL
	user.SetFlashMessage(w, r, "success", "Schedule URL has been updated!")
	http.Redirect(w, r, "/admin", http.StatusTemporaryRedirect)
}

/*
*	User routes
 */
func userIndex(w http.ResponseWriter, r *http.Request) {
	users, err := user.All()
	if err != nil {
		log.Println(errors.Wrap(err, "admin.user.index"))
	}
	u, err := user.FromSession(r)
	if err != nil {
		log.Println(errors.Wrap(err, "admin.user.index"))
	}
	data := map[string]interface{}{
		"User":    u,
		"Users":   users,
		"Alert":   user.GetFlashMessage(w, r, "alert"),
		"Success": user.GetFlashMessage(w, r, "success"),
	}

	adminRenderer.HTML(w, http.StatusOK, "user.html", data)
}

// Store the user in the database
func userStore(w http.ResponseWriter, r *http.Request) {
	// Parse form and get the username submitted
	r.ParseForm()
	username := r.Form.Get("username")
	// Create a user object
	u := user.User{
		// Username is lowercase since that's what Twitch
		// returns through their Oauth response
		Username: strings.ToLower(username),
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

	if err := user.Create(username); err != nil || username == "" {
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

/*
*	Article handlers
 */
// articleIndex renders a paginated list of the articles in the DB
func articleIndex(w http.ResponseWriter, r *http.Request) {
	// Get current page number
	p := getArticlePage(r)

	// Retrieve articles for current page
	articles, err := article.Page(p, false)
	if err != nil {
		// If something goes wrong we render the 500-page
		log.Println(errors.Wrap(err, "admin.article.index"))
		data := getPagedata()
		renderer.HTML(w, http.StatusInternalServerError, "500.html", data)
		return
	}
	for i, a := range articles {
		a.ShortenBody()
		articles[i] = a
	}

	u, err := user.FromSession(r)
	if err != nil {
		// If something goes wrong we render the 500-page
		log.Println(errors.Wrap(err, "admin.article.index"))
		data := getPagedata()
		renderer.HTML(w, http.StatusInternalServerError, "500.html", data)
		return
	}

	// Set up the data we need
	data := map[string]interface{}{
		"User":     u,
		"Articles": articles,
		"Alert":    user.GetFlashMessage(w, r, "alert"),
		"Success":  user.GetFlashMessage(w, r, "success"),
		"NextPage": p + 1,
		"PrevPage": p - 1,
		"CurrPage": p,
	}
	// Total page count
	data["LastPage"], err = article.PageCount()

	adminRenderer.HTML(w, http.StatusOK, "article.html", data)
}

func articleCreate(w http.ResponseWriter, r *http.Request) {
	adminRenderer.HTML(w, http.StatusOK, "create_article.html", nil)
}

func articleStore(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	a := article.Article{
		Title: r.Form.Get("title"),
		Body:  r.Form.Get("body"),
	}
	a.Published = false
	if r.FormValue("published") == "1" {
		a.Published = true
	}

	u, err := user.FromSession(r)
	if err != nil {
		http.Redirect(w, r, "/admin/article", http.StatusSeeOther)
	}

	a.AddAuthor(u)

	if err := a.Create(); err != nil {
		user.SetFlashMessage(w, r, "alert", "An error occured while trying to create the article.")
		log.Println(errors.Wrap(err, "handlers.createArticle"))
		http.Redirect(w, r, "/admin/article", http.StatusSeeOther)
		return
	}

	user.SetFlashMessage(w, r, "success", "The article was saved successfully")
	http.Redirect(w, r, "/admin/article", http.StatusSeeOther)
}

func articleEdit(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	a, err := article.Get(id, nil)
	if err != nil {
		user.SetFlashMessage(w, r, "alert", "Couldn't find the article...")
		log.Println(errors.Wrap(err, "handlers.articleEdit"))
		http.Redirect(w, r, "/admin/article", http.StatusSeeOther)
		return
	}

	data := map[string]interface{}{
		"Article": a,
		"Alert":   user.GetFlashMessage(w, r, "alert"),
		"Success": user.GetFlashMessage(w, r, "success"),
	}

	adminRenderer.HTML(w, http.StatusOK, "edit_article.html", data)
}

func articleUpdate(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	a, err := article.Get(id, nil)
	if err != nil {
		user.SetFlashMessage(w, r, "alert", "Couldn't find the article...")
		log.Println(errors.Wrap(err, "handlers.articleUpdate"))
		http.Redirect(w, r, "/admin/article/", http.StatusSeeOther)
		return
	}

	u, err := user.FromSession(r)
	// No reason to do more error handling since we only use the user for author
	if err != nil {
		log.Println(errors.Wrap(err, "handlers.articleUpdate"))
	} else if !a.AuthorExists(u) {
		a.Authors = append(a.Authors, u)
	}

	r.ParseForm()
	title := r.FormValue("title")
	body := r.FormValue("body")
	a.Published = false
	if r.FormValue("published") == "1" {
		a.Published = true
	}

	if title != "" {
		a.Title = title
	}

	if body != "" {
		a.Body = body
	}

	if err = a.Update(); err != nil {
		user.SetFlashMessage(w, r, "alert", "Something went wrong while updating...")
		log.Println(errors.Wrap(err, "handlers.articleUpdate"))
		http.Redirect(w, r, "/admin/article/"+id, http.StatusSeeOther)
		return
	}

	user.SetFlashMessage(w, r, "success", "Changes have been saved")
	http.Redirect(w, r, "/admin/article/"+id, http.StatusSeeOther)
}

func articleDelete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := article.Delete(id)
	if err != nil {
		log.Println(errors.Wrap(err, "handlers.articleUpdate"))
		user.SetFlashMessage(w, r, "alert", "Couldn't find the article you tried to delete")
		http.Redirect(w, r, "/admin/article", http.StatusSeeOther)
		return
	}

	user.SetFlashMessage(w, r, "success", "The article was deleted")
	http.Redirect(w, r, "/admin/article", http.StatusSeeOther)
}
