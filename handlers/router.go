package handlers

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"

	"github.com/esamarathon/website/config"
	"github.com/esamarathon/website/handlers/admin"
	. "github.com/esamarathon/website/handlers/helpers"
	"github.com/rs/cors"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func init() {
	router.PathPrefix("/static").Handler(handleStatic("public", "/static"))
	router.HandleFunc("/", CSP(Index)).Methods("GET")
	router.HandleFunc("/schedule", CSP(Schedule)).Methods("GET")
	router.HandleFunc("/news", CSP(News)).Methods("GET")
	router.HandleFunc("/news/{id}", CSP(Article)).Methods("GET")
	router.HandleFunc("/auth", CSP(AuthRedirect)).Methods("GET")
	router.HandleFunc("/auth/callback", CSP(AuthCallback)).Methods("GET")
	router.HandleFunc("/login", CSP(HandleAuth)).Methods("GET")
	router.HandleFunc("/logout", CSP(HandleLogout)).Methods("GET")

	admin.AdminRoutes("/admin", router)

	router.HandleFunc("/{name}", CSP(Page)).Methods("GET")

	router.NotFoundHandler = http.HandlerFunc(CSP(HandleNotFound))
}

func handleStatic(dir, prefix string) http.HandlerFunc {
	fs := http.FileServer(http.Dir(dir))
	realHandler := http.StripPrefix(prefix, fs).ServeHTTP
	f := func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Cache-Control", "max-age=2592000")
		realHandler(w, req)
	}
	return makeGzipHandler(f)

}

func Router(version string) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{config.Config.SiteURL},
		AllowedMethods:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	return c.Handler(router)
}

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func makeGzipHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			fn(w, r)
			return
		}
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		gzr := gzipResponseWriter{Writer: gz, ResponseWriter: w}
		fn(gzr, r)
	}
}
