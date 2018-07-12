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
	s = strings.ToLower(s)
	s = strings.Replace(s, " ", "-", -1)
	s = url.PathEscape(s)
	return s
}

func CSP(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header("Content-Security-Policy", "default-src 'self'; script-src 'sha256-4kSj415Ktl8nD2hH1J/vYCiFDzN8b/SnwBz3WA4H0IY=' 'https://pagead2.googlesyndication.com'; img-src *")
	}
}
