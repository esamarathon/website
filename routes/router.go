package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/olenedr/esamarathon/handlers"
	"github.com/olenedr/esamarathon/middleware"
)

// GetRouter returns an instance of the Mux router
func GetRouter() *mux.Router {
	router := mux.NewRouter()

	requiresAuth := middleware.AuthMiddleware

	router.PathPrefix("/static").Handler(handleStatic("public", "/static"))
	router.HandleFunc("/", handlers.Index).Methods("GET", "OPTIONS")
	router.HandleFunc("/schedule", handlers.Schedule).Methods("GET", "OPTIONS")
	router.HandleFunc("/news", handlers.News).Methods("GET", "OPTIONS")
	router.HandleFunc("/test", handlers.Test).Methods("GET", "OPTIONS")
	router.HandleFunc("/auth", handlers.AuthRedirect).Methods("GET")
	router.HandleFunc("/auth/callback", handlers.AuthCallback).Methods("GET")
	router.HandleFunc("/login", handlers.HandleAuth).Methods("GET")
	router.HandleFunc("/logout", handlers.HandleLogout).Methods("GET")

	//Admin routes
	router.HandleFunc("/admin", requiresAuth(handlers.AdminIndex))
	router.HandleFunc("/admin/toggle", requiresAuth(handlers.AdminToggleLive))
	router.HandleFunc("/admin/user", requiresAuth(handlers.AdminUserIndex))
	router.HandleFunc("/admin/article", requiresAuth(handlers.AdminArticleIndex))

	return router
}

func handleStatic(dir, prefix string) http.HandlerFunc {
	fs := http.FileServer(http.Dir(dir))
	realHandler := http.StripPrefix(prefix, fs).ServeHTTP
	return func(w http.ResponseWriter, req *http.Request) {
		realHandler(w, req)
	}
}
