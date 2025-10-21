package router

import (
	"log"
	"net/http"

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

	if environment := config.Environment; environment == "development" {
		log.Println("Server starting on :8080")
		log.Fatal(http.ListenAndServe(":8080", router))
	} else if environment == "production" {
		log.Println("Server starting on :443")
	} else {
		log.Println("Server can't be started in this environment")
	}
	return router
}
