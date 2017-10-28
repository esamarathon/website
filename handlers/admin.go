package handlers

import (
	"fmt"
	"log"
	"net/http"

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
	router.HandleFunc(base, requireAuth(index)).Methods("GET")
	router.HandleFunc(base+"/toggle", requireAuth(toggleLivemode)).Methods("GET")
	router.HandleFunc(base+"/user", requireAuth(userIndex)).Methods("GET")
	router.HandleFunc(base+"/user", requireAuth(userStore)).Methods("POST")
	router.HandleFunc(base+"/article", requireAuth(articleIndex)).Methods("GET")
	router.HandleFunc(base+"/article/create", requireAuth(articleCreate)).Methods("GET")
	router.HandleFunc(base+"/article/create", requireAuth(articleStore)).Methods("POST")
	router.HandleFunc(base+"/article/edit/{id}", requireAuth(editArticleIndex)).Methods("GET")
	router.HandleFunc(base+"/article/update/{id}", requireAuth(updateArticle)).Methods("POST")
	router.HandleFunc(base, requireAuth(index)).Methods("GET", "POST")
}

var adminRenderer = grender.New(grender.Options{
	TemplatesGlob: "templates_admin/*.html",
	PartialsGlob:  "templates_admin/partials/*.html",
})

func updateArticle(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	a, err := article.Get(id)
	if err != nil {
		log.Println(errors.Wrap(err, "handlers.updateArticle"))
		// TODO: Add flash message letting the user know what went wrong
		http.Redirect(w, r, "/admin/article/"+id, http.StatusSeeOther)
	}

	u, err := user.UserFromSession(r)
	// No reason to do more error handling since we only use the user for author
	if err != nil {
		log.Println(errors.Wrap(err, "handlers.updateArticle"))
	} else if !a.AuthorExists(u) {
		a.Authors = append(a.Authors, u)
	}

	r.ParseForm()
	title := r.FormValue("title")
	body := r.FormValue("body")

	if title != "" {
		a.Title = title
	}

	if body != "" {
		a.Body = body
	}

	if err = a.Update(); err != nil {
		log.Println(errors.Wrap(err, "handlers.updateArticle"))
		// TODO: Add flash message letting the user know that saving the article failed
		http.Redirect(w, r, "/admin/article/"+id, http.StatusSeeOther)
	}

	http.Redirect(w, r, "/admin/article", http.StatusSeeOther)
}

func editArticleIndex(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	a, err := article.Get(id)
	if err != nil {
		log.Println(errors.Wrap(err, "handlers.editArticleIndex"))
		adminRenderer.HTML(w, http.StatusInternalServerError, "index.html", nil)
		return
	}

	data := map[string]interface{}{
		"Article": a,
	}

	adminRenderer.HTML(w, http.StatusOK, "edit_article.html", data)
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

func userStore(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userName := r.Form.Get("username")

	if err := user.Insert(userName); err != nil || userName == "" {
		// TODO:Handle error better
		fmt.Fprint(w, err)
		return
	}

	http.Redirect(w, r, "/admin/user", http.StatusSeeOther)
}

func articleIndex(w http.ResponseWriter, r *http.Request) {
	// Change with actual articledata
	articles, err := article.All()
	if err != nil {
		log.Println(errors.Wrap(err, "admin.article.index"))
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
	http.Redirect(w, r, "/admin", http.StatusTemporaryRedirect)
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

	u, err := user.UserFromSession(r)
	if err != nil {
		http.Redirect(w, r, "/admin/article", http.StatusSeeOther)
	}

	a.AddAuthor(u)

	//TODO: if something needs to verified, this should be done here
	if err := a.Create(); err != nil {
		// TODO: Handle failure better
		log.Println(errors.Wrap(err, "handlers.createArticle"))
		fmt.Fprint(w, err)
		return
	}

	http.Redirect(w, r, "/admin/article", http.StatusSeeOther)
}
