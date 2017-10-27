package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/olenedr/esamarathon/models/article"

	"github.com/dannyvankooten/grender"
	"github.com/pkg/errors"

	"github.com/olenedr/esamarathon/models/article"
	"github.com/olenedr/esamarathon/models/setting"
	"github.com/olenedr/esamarathon/models/user"

	"github.com/gorilla/mux"
	"github.com/olenedr/esamarathon/middleware"
)

func AdminRoutes(base string, router *mux.Router) {
	requireAuth := middleware.AuthMiddleware
<<<<<<< Updated upstream
	router.HandleFunc(base, requireAuth(index)).Methods("GET")
	router.HandleFunc(base+"/toggle", requireAuth(toggleLivemode)).Methods("GET")
	router.HandleFunc(base+"/user", requireAuth(userIndex)).Methods("GET")
	router.HandleFunc(base+"/user", requireAuth(userCreate)).Methods("POST")
	router.HandleFunc(base+"/article", requireAuth(articleIndex)).Methods("GET")
=======
	router.HandleFunc(base+"/user/create", requireAuth(createUser)).Methods("POST")
	router.HandleFunc(base+"/article/create", requireAuth(createArticle)).Methods("POST")
	router.HandleFunc(base+"/article/create", requireAuth(createArticleIndex)).Methods("GET")
	router.HandleFunc(base+"/article/edit/{id}", requireAuth(editArticleIndex)).Methods("GET")
	router.HandleFunc(base, requireAuth(index)).Methods("GET", "POST")
>>>>>>> Stashed changes
}

var adminRenderer = grender.New(grender.Options{
	TemplatesGlob: "templates_admin/*.html",
	PartialsGlob:  "templates_admin/partials/*.html",
})

func editArticleIndex(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	a, err := article.Get(id)
	if err != nil {
		adminRenderer.HTML(w, http.StatusInternalServerError, "index.html", nil)
		return
	}

	Page["article"] = a
	adminRenderer.HTML(w, http.StatusOK, "edit_article.html", Page)
}

func index(w http.ResponseWriter, r *http.Request) {
	// Change with actual status from DB
	u, userErr := user.UserFromSession(r)
	s, settingErr := setting.GetLiveMode().AsBool()
	if settingErr != nil {
		log.Println(errors.Wrap(settingErr, "admin.index"))
	}
	if userErr != nil {
		log.Println(errors.Wrap(userErr, "admin.index"))
	}
	data := map[string]interface{}{
		"User":   u,
		"Status": s,
	}

	adminRenderer.HTML(w, http.StatusOK, "index.html", data)
}

func userIndex(w http.ResponseWriter, r *http.Request) {
	users, err := user.All()
	if err != nil {
		log.Println(errors.Wrap(err, "admin.user.index"))
	}
	u, err := user.UserFromSession(r)
	if err != nil {
		log.Println(errors.Wrap(err, "admin.user.index"))
	}
	data := map[string]interface{}{
		"User":  u,
		"Users": users,
	}

	adminRenderer.HTML(w, http.StatusOK, "user.html", data)
}

<<<<<<< Updated upstream
func userCreate(w http.ResponseWriter, r *http.Request) {
=======
func createUser(w http.ResponseWriter, r *http.Request) {
>>>>>>> Stashed changes
	r.ParseForm()
	userName := r.Form.Get("username")

	if err := user.Insert(userName); err != nil || userName == "" {
		// TODO:Handle error better
		fmt.Fprint(w, err)
		return
	}

<<<<<<< Updated upstream
	log.Println("redirect")
	http.Redirect(w, r, "/admin/user", http.StatusSeeOther)
}

func articleIndex(w http.ResponseWriter, r *http.Request) {
	// Change with actual articledata
	timestamp := time.Now()
	body := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
	articles := []article.Article{
		{
			ID:        "1",
			Title:     "Lorem ipsum",
			Body:      body,
			CreatedAt: timestamp,
			UpdatedAt: timestamp,
		},
		{
			ID:        "2",
			Title:     "Dolor sit amet",
			Body:      body,
			CreatedAt: timestamp,
			UpdatedAt: timestamp,
		},
	}
	u, err := user.UserFromSession(r)
	if err != nil {
		log.Println(errors.Wrap(err, "admin.article.index"))
	}
	data := map[string]interface{}{
		"User":     u,
		"Articles": articles,
	}

	adminRenderer.HTML(w, http.StatusOK, "article.html", data)
}

func toggleLivemode(w http.ResponseWriter, r *http.Request) {
	setting.GetLiveMode().Toggle()
=======
	http.Redirect(w, r, "/admin", http.StatusTemporaryRedirect)
}

func createArticleIndex(w http.ResponseWriter, r *http.Request) {
	adminRenderer.HTML(w, http.StatusOK, "article.html", nil)
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("title", r.Form.Get("title"))
	a := article.Article{
		Title: r.Form.Get("title"),
		Body:  r.Form.Get("body"),
	}

	//TODO: if something needs to verified, this should be done here
	if err := a.Create(); err != nil {
		// TODO: Handle failure better
		log.Println(errors.Wrap(err, "handlers.createArticle"))
		fmt.Fprint(w, err)
		return
	}
>>>>>>> Stashed changes

	http.Redirect(w, r, "/admin", http.StatusTemporaryRedirect)
}
