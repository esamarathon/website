package viewmodels

import (
	"log"
	"net/http"

	"github.com/esamarathon/website/config"
	"github.com/esamarathon/website/models/article"
	"github.com/esamarathon/website/models/menu"
	"github.com/esamarathon/website/models/page"
	"github.com/esamarathon/website/models/social"
	"github.com/esamarathon/website/models/user"
	"github.com/pkg/errors"
)

type AdminView struct {
	User    user.User
	Alert   string // Alert message
	Success string // Success message
}

type pagination struct {
	NextPage int
	PrevPage int
	CurrPage int
	LastPage int
}

type adminIndexView struct {
	AdminView
	Livemode       bool
	ScheduleAPIURL string
	ShowSchedule   bool
	Frontpage      frontPage
	SocialLinks    social.SocialLinks
}

type adminUserIndexView struct {
	AdminView
	Users []user.User
}

type adminArticleIndexView struct {
	AdminView
	pagination
	Articles []article.Article
}

type adminArticleCreateView struct {
	AdminView
}

type adminArticleEditView struct {
	AdminView
	Article article.Article
}

type adminMenuIndexView struct {
	AdminView
	Menu menu.Menu
}

type adminPageIndexView struct {
	AdminView
	pagination
	Pages []page.Page
}

type adminPageCreateView struct {
	AdminView
}

type adminPageEditView struct {
	AdminView
	Page page.Page
}

func getUser(r *http.Request) user.User {
	u, userErr := user.FromSession(r)
	if userErr != nil {
		log.Println(errors.Wrap(userErr, "getUser"))
	}
	return u
}

func getAdminView(w http.ResponseWriter, r *http.Request) AdminView {
	return AdminView{
		User:    getUser(r),
		Alert:   user.GetFlashMessage(w, r, "alert"),
		Success: user.GetFlashMessage(w, r, "success"),
	}
}

func AdminIndex(w http.ResponseWriter, r *http.Request) adminIndexView {
	view := adminIndexView{
		AdminView:      getAdminView(w, r),
		Livemode:       config.Config.LiveMode,
		ScheduleAPIURL: config.Config.ScheduleAPIURL,
		ShowSchedule:   config.Config.ShowSchedule,
		Frontpage:      getFrontpage(),
		SocialLinks:    social.Get(),
	}

	return view
}

func AdminUserIndex(w http.ResponseWriter, r *http.Request) adminUserIndexView {
	users, err := user.All()
	if err != nil {
		log.Println(errors.Wrap(err, "admin.user.index"))
	}
	view := adminUserIndexView{
		AdminView: getAdminView(w, r),
		Users:     users,
	}

	return view
}

func AdminArticleIndex(w http.ResponseWriter, r *http.Request) adminArticleIndexView {
	view := adminArticleIndexView{
		AdminView: getAdminView(w, r),
	}
	return view
}

func AdminArticleCreate(w http.ResponseWriter, r *http.Request) adminArticleCreateView {
	view := adminArticleCreateView{
		AdminView: getAdminView(w, r),
	}
	return view
}

func AdminArticleEdit(w http.ResponseWriter, r *http.Request) adminArticleEditView {
	view := adminArticleEditView{
		AdminView: getAdminView(w, r),
	}
	return view
}

func AdminMenuIndex(w http.ResponseWriter, r *http.Request) adminMenuIndexView {
	view := adminMenuIndexView{
		Menu:      menu.Get(),
		AdminView: getAdminView(w, r),
	}
	return view
}

func AdminPageIndex(w http.ResponseWriter, r *http.Request) adminPageIndexView {
	view := adminPageIndexView{
		AdminView: getAdminView(w, r),
	}
	return view
}

func AdminPageCreate(w http.ResponseWriter, r *http.Request) adminPageCreateView {
	view := adminPageCreateView{
		AdminView: getAdminView(w, r),
	}
	return view
}

func AdminPageEdit(w http.ResponseWriter, r *http.Request) adminPageEditView {
	view := adminPageEditView{
		AdminView: getAdminView(w, r),
	}
	return view
}
