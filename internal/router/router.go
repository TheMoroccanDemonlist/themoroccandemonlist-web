package router

import (
	"log"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/themoroccandemonlist/themoroccandemonlist-web/internal/app"
	"github.com/themoroccandemonlist/themoroccandemonlist-web/internal/config"
)

func New(config *config.Config, application *app.App) *mux.Router {
	router := mux.NewRouter()
	router.Use(csrf.Protect(config.SessionKey))

	// router.HandleFunc("/", application.HomeHandler).Methods("GET")

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
