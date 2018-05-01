package admin

import (
	"log"
	"net/http"

	"github.com/pkg/errors"

	. "github.com/esamarathon/website/handlers/helpers"
	"github.com/esamarathon/website/models/article"
	"github.com/esamarathon/website/models/page"
	"github.com/esamarathon/website/models/user"
	"github.com/esamarathon/website/viewmodels"
	
	"github.com/gorilla/mux"
)

/*
* Generic Page handlers
*/
func pageIndex(w http.ResponseWriter, r *http.Request) {
	p := GetPagination(r)
	view := viewmodels.AdminPageIndex(w, r)

	pages, err := page.Pagination(p, false)
	if err != nil {
		// If something goes wrong we render the 500-page
		log.Println(errors.Wrap(err, "admin.page.index"))
		HandleInternalError(w)
		return
	}
	for i, p := range pages {
		p.ShortenBody()
		pages[i] = p
	}

	// Total page count
	count, err := page.PageCount(false)
	if err != nil {
		// If something goes wrong we render the 500-page
		log.Println(errors.Wrap(err, "admin.page.index"))
		HandleInternalError(w)
		return
	}

	// Set all the necessary values
	view.Pages = pages
	view.NextPage = p + 1
	view.PrevPage = p - 1
	view.CurrPage = p
	view.LastPage = count

	adminRenderer.HTML(w, http.StatusOK, "page.html", view)
}

func pageCreate(w http.ResponseWriter, r *http.Request) {
	adminRenderer.HTML(w, http.StatusOK, "create_page.html", viewmodels.AdminPageCreate(w, r))
}

func pageStore(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	p := page.Page {
		FriendlyName: r.Form.Get("name"),
		Article: article.Article {
			Title: r.Form.Get("title"),
			Body:  r.Form.Get("body"),
		},
	}

	if p.Title == "" || p.Body == "" {
		r.Method = "GET"
		user.SetFlashMessage(w, r, "alert", "Invalid input data.")
		log.Println("Missing input data, handlers.createPage")
		http.Redirect(w, r, "/admin/page/create", http.StatusSeeOther)
		return
	}

	if p.FriendlyName == "" {
		p.FriendlyName = p.Title
	}

	p.FriendlyName = Urlify(p.FriendlyName)

	p.Published = false
	if r.FormValue("published") == "1" {
		p.Published = true
	}

	u, err := user.FromSession(r)
	if err != nil {
		user.SetFlashMessage(w, r, "alert", "An error occured while retriving the user.")
		log.Println(errors.Wrap(err, "handlers.createPage"))
		http.Redirect(w, r, "/admin/page", http.StatusSeeOther)
		return
	}

	p.AddAuthor(u)

	if err := p.Create(); err != nil {
		user.SetFlashMessage(w, r, "alert", "An error occured while trying to create the page.")
		log.Println(errors.Wrap(err, "handlers.createPage"))
		http.Redirect(w, r, "/admin/page", http.StatusSeeOther)
		return
	}

	user.SetFlashMessage(w, r, "success", "The page was saved successfully")
	http.Redirect(w, r, "/admin/page", http.StatusSeeOther)
}

func pageEdit(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	view := viewmodels.AdminPageEdit(w, r)

	p, err := page.Get(id, nil)
	if err != nil {
		user.SetFlashMessage(w, r, "alert", "Couldn't find the page...")
		log.Println(errors.Wrap(err, "handlers.pageEdit"))
		http.Redirect(w, r, "/admin/page", http.StatusSeeOther)
		return
	}
	view.Page = p

	adminRenderer.HTML(w, http.StatusOK, "edit_page.html", view)
}

func pageUpdate(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	p, err := page.Get(id, nil)
	if err != nil {
		user.SetFlashMessage(w, r, "alert", "Couldn't find the page...")
		log.Println(errors.Wrap(err, "handlers.pageUpdate"))
		http.Redirect(w, r, "/admin/page/", http.StatusSeeOther)
		return
	}

	u, err := user.FromSession(r)
	// No reason to do more error handling since we only use the user for author
	if err != nil {
		log.Println(errors.Wrap(err, "handlers.pageUpdate"))
	} else if !p.AuthorExists(u) {
		p.Authors = append(p.Authors, u)
	}

	r.ParseForm()
	title := r.FormValue("title")
	name := r.FormValue("name")
	body := r.FormValue("body")
	p.Published = false
	if r.FormValue("published") == "1" {
		p.Published = true
	}

	if title != "" {
		p.Title = title
	}

	if name != "" {
		p.FriendlyName = Urlify(name)
	}

	if body != "" {
		p.Body = body
	}

	if err = p.Update(); err != nil {
		user.SetFlashMessage(w, r, "alert", "Something went wrong while updating...")
		log.Println(errors.Wrap(err, "handlers.pageUpdate"))
		http.Redirect(w, r, "/admin/page/"+id, http.StatusSeeOther)
		return
	}

	user.SetFlashMessage(w, r, "success", "Changes have been saved")
	http.Redirect(w, r, "/admin/page/"+id, http.StatusSeeOther)
}

func pageDelete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := page.Delete(id)
	if err != nil {
		log.Println(errors.Wrap(err, "handlers.pageDelete"))
		user.SetFlashMessage(w, r, "alert", "Couldn't find the page you tried to delete")
		http.Redirect(w, r, "/admin/page", http.StatusSeeOther)
		return
	}

	user.SetFlashMessage(w, r, "success", "The page was deleted")
	http.Redirect(w, r, "/admin/page", http.StatusSeeOther)
}
