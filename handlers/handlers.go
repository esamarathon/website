package handlers

import (
	"fmt"
	"net/http"

	"github.com/dannyvankooten/grender"
	"github.com/olenedr/esamarathon/article"
	"github.com/olenedr/esamarathon/db"
)

var renderer = grender.New(grender.Options{
	TemplatesGlob: "templates/*.html",
})

// Index returns index view
func Index(w http.ResponseWriter, r *http.Request) {
	m := meta{
		"ESA Marathon",
		"Welcome to European Speedrunner Assembly!",
		"http://www.esamarathon.com/images/esa/europeanspeedrunnerassembly.png",
	}
	c := content{
		"ESA Marathon",
		"Lorem ipsum",
	}

	p := page{
		&m,
		&c,
	}
	renderer.HTML(w, http.StatusOK, "index.html", p)
}

// Test is a test handler for Ole to debug stuff
func Test(w http.ResponseWriter, r *http.Request) {
	m := meta{
		"ESA Marathon",
		"Welcome to European Speedrunner Assembly!",
		"http://www.esamarathon.com/images/esa/europeanspeedrunnerassembly.png",
	}
	c := content{
		"ESA Marathon",
		"Lorem ipsum",
	}

	p := page{
		&m,
		&c,
	}

	if err := db.Connection.C("articles").Insert(p); err != nil {
		fmt.Println("Failed to insert document to DB")
	}

	// count, _ := db.Connection.C("articles").Count()
	a := article.Article{}

	articles, _ := a.All()
	fmt.Printf("Articles in the DB: %v\n", articles)

	renderer.HTML(w, http.StatusOK, "index.html", p)
}
