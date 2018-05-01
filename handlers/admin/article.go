package admin

import (
	"log"
	"net/http"

	"github.com/pkg/errors"

	. "github.com/esamarathon/website/handlers/helpers"
	"github.com/esamarathon/website/models/article"
	"github.com/esamarathon/website/models/user"
	"github.com/esamarathon/website/viewmodels"

	"github.com/gorilla/mux"
)

/*
*	Article handlers
*/
// articleIndex renders a paginated list of the articles in the DB
func articleIndex(w http.ResponseWriter, r *http.Request) {
	// Get current page number
	p := GetPagination(r)
	view := viewmodels.AdminArticleIndex(w, r)

	// Retrieve articles for current page
	articles, err := article.Page(p, false)
	if err != nil {
		// If something goes wrong we render the 500-page
		log.Println(errors.Wrap(err, "admin.article.index"))
		HandleInternalError(w)
		return
	}
	for i, a := range articles {
		a.ShortenBody()
		articles[i] = a
	}

	// Total page count
	count, err := article.PageCount(false)
	if err != nil {
		// If something goes wrong we render the 500-page
		log.Println(errors.Wrap(err, "admin.article.index"))
		HandleInternalError(w)
		return
	}

	// Set all the necessary values
	view.Articles = articles
	view.NextPage = p + 1
	view.PrevPage = p - 1
	view.CurrPage = p
	view.LastPage = count

	adminRenderer.HTML(w, http.StatusOK, "article.html", view)
}

func articleCreate(w http.ResponseWriter, r *http.Request) {
	adminRenderer.HTML(w, http.StatusOK, "create_article.html", viewmodels.AdminArticleCreate(w, r))
}

func articleStore(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	a := article.Article{
		Title: r.Form.Get("title"),
		Body:  r.Form.Get("body"),
	}
	if a.Title == "" || a.Body == "" {
		r.Method = "GET"
		user.SetFlashMessage(w, r, "alert", "Invalid input data.")
		log.Println("Missing input data, handlers.createArticle")
		http.Redirect(w, r, "/admin/article/create", http.StatusSeeOther)
		return
	}
	a.Published = false
	if r.FormValue("published") == "1" {
		a.Published = true
	}

	u, err := user.FromSession(r)
	if err != nil {
		user.SetFlashMessage(w, r, "alert", "An error occured while retriving the user.")
		log.Println(errors.Wrap(err, "handlers.createArticle"))
		http.Redirect(w, r, "/admin/article", http.StatusSeeOther)
		return
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
	view := viewmodels.AdminArticleEdit(w, r)

	a, err := article.Get(id, nil)
	if err != nil {
		user.SetFlashMessage(w, r, "alert", "Couldn't find the article...")
		log.Println(errors.Wrap(err, "handlers.articleEdit"))
		http.Redirect(w, r, "/admin/article", http.StatusSeeOther)
		return
	}
	view.Article = a

	adminRenderer.HTML(w, http.StatusOK, "edit_article.html", view)
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
