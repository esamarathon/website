package helpers

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/dannyvankooten/grender"
)

var Renderer = grender.New(grender.Options{
	TemplatesGlob: "templates_minified/*.html",
	PartialsGlob:  "templates_minified/partials/*.html",
})

// Extracts the page query param if present
func GetPagination(r *http.Request) int {
	// Default page number
	page := "0"

	// If a page is specified we use that instead
	if query := r.URL.Query()["page"]; len(query) != 0 {
		page = query[0]
	}
	p, err := strconv.Atoi(page)

	// If we failed to get the page number
	// we just set it to 0 (first page)
	if err != nil {
		p = 0
	}
	return p
}

func Urlify(s string) string {
	s = strings.Replace(s, " ", "-", -1)
	s = url.PathEscape(s)
	return s
}