package router

import (
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/themoroccandemonlist/themoroccandemonlist-web/internal/app"
	"github.com/themoroccandemonlist/themoroccandemonlist-web/internal/config"
	"github.com/themoroccandemonlist/themoroccandemonlist-web/internal/handlers"
	"github.com/themoroccandemonlist/themoroccandemonlist-web/internal/middleware"
)

func New(config *config.Config, application *app.App) *mux.Router {
	handler := &handlers.Handler{App: application}

	router := mux.NewRouter()
	router.Use(csrf.Protect(config.SessionKey))
	router.Use(middleware.Logging)
	router.Use(middleware.CSP)

	router.HandleFunc("/auth", handler.GoogleAuth).Methods("GET")
	router.HandleFunc("/auth/google/callback", handler.GoogleCallback).Methods("GET")
	router.HandleFunc("/logout", handler.Logout).Methods("GET")

	return router
}
