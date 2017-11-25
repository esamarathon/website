package viewmodels

import (
	"log"
	"net/http"

	"github.com/olenedr/esamarathon/config"
	"github.com/olenedr/esamarathon/models/article"
	"github.com/olenedr/esamarathon/models/user"
	"github.com/pkg/errors"
)

type adminIndexView struct {
	User           user.User
	Livemode       bool
	ScheduleAPIURL string
	Alert          string // Alert message
	Success        string // Success message
}

type adminUserIndexView struct {
	User    user.User
	Users   []user.User
	Alert   string // Alert message
	Success string // Success message
}

type adminArticleIndexView struct {
	User     user.User
	Articles []article.Article
	Alert    string // Alert message
	Success  string // Success message
	NextPage int
	PrevPage int
	CurrPage int
	LastPage int
}

type adminArticleCreateView struct {
	User    user.User
	Alert   string // Alert message
	Success string // Success message
}

type adminArticleEditView struct {
	User    user.User
	Article article.Article
	Alert   string // Alert message
	Success string // Success message
}

func getUser(r *http.Request) user.User {
	u, userErr := user.FromSession(r)
	if userErr != nil {
		log.Println(errors.Wrap(userErr, "getUser"))
	}
	return u
}

func AdminIndex(w http.ResponseWriter, r *http.Request) adminIndexView {
	view := adminIndexView{
		User:           getUser(r),
		Livemode:       config.Config.LiveMode,
		ScheduleAPIURL: config.Config.ScheduleAPIURL,
		Alert:          user.GetFlashMessage(w, r, "alert"),
		Success:        user.GetFlashMessage(w, r, "success"),
	}

	return view
}

func AdminUserIndex(w http.ResponseWriter, r *http.Request) adminUserIndexView {
	users, err := user.All()
	if err != nil {
		log.Println(errors.Wrap(err, "admin.user.index"))
	}
	view := adminUserIndexView{
		User:    getUser(r),
		Users:   users,
		Alert:   user.GetFlashMessage(w, r, "alert"),
		Success: user.GetFlashMessage(w, r, "success"),
	}

	return view
}

func AdminArticleIndex(w http.ResponseWriter, r *http.Request) adminArticleIndexView {
	view := adminArticleIndexView{
		User:    getUser(r),
		Alert:   user.GetFlashMessage(w, r, "alert"),
		Success: user.GetFlashMessage(w, r, "success"),
	}
	return view
}

func AdminArticleCreate(w http.ResponseWriter, r *http.Request) adminArticleCreateView {
	view := adminArticleCreateView{
		User:    getUser(r),
		Alert:   user.GetFlashMessage(w, r, "alert"),
		Success: user.GetFlashMessage(w, r, "success"),
	}
	return view
}

func AdminArticleEdit(w http.ResponseWriter, r *http.Request) adminArticleEditView {
	view := adminArticleEditView{
		User:    getUser(r),
		Alert:   user.GetFlashMessage(w, r, "alert"),
		Success: user.GetFlashMessage(w, r, "success"),
	}
	return view
}
