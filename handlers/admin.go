package handlers

import (
	"net/http"
	"time"

	"github.com/dannyvankooten/grender"
	"github.com/olenedr/esamarathon/article"
	"github.com/olenedr/esamarathon/models/setting"
	"github.com/olenedr/esamarathon/user"
)

var adminRenderer = grender.New(grender.Options{
	TemplatesGlob: "templates_admin/*.html",
	PartialsGlob:  "templates_admin/partials/*.html",
})

func AdminIndex(w http.ResponseWriter, r *http.Request) {
	// Change with actual status from DB
	u, userErr := user.UserFromSession(r)
	s, settingErr := setting.GetLiveMode().AsBool()
	if userErr != nil || settingErr != nil {
		http.Redirect(w, r, "500.html", http.StatusTemporaryRedirect)
	}
	v := map[string]interface{}{
		"User":   u,
		"Status": s,
	}

	adminRenderer.HTML(w, http.StatusOK, "index.html", v)
}

func AdminUserIndex(w http.ResponseWriter, r *http.Request) {
	// Change with actual userdata
	users := []user.User{
		{
			Username: "Korkn",
		},
		{
			Username: "egreb__",
		},
	}
	u, _ := user.UserFromSession(r)
	v := map[string]interface{}{
		"User":  u,
		"Users": users,
	}

	adminRenderer.HTML(w, http.StatusOK, "user.html", v)
}

func AdminArticleIndex(w http.ResponseWriter, r *http.Request) {
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
	u, _ := user.UserFromSession(r)
	v := map[string]interface{}{
		"User":     u,
		"Articles": articles,
	}

	adminRenderer.HTML(w, http.StatusOK, "article.html", v)
}

func AdminToggleLive(w http.ResponseWriter, r *http.Request) {
	setting.GetLiveMode().Toggle()

	http.Redirect(w, r, "/admin", http.StatusTemporaryRedirect)
}
