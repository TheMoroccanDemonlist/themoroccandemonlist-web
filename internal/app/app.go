package app

import (
	"html/template"

	"github.com/themoroccandemonlist/themoroccandemonlist-web/internal/config"
)

type App struct {
	Config    *config.Config
	Templates *template.Template
}

func New(config *config.Config) *App {
	templates := template.Must(template.ParseGlob("web/templates/*.html"))
	return &App{
		Config:    config,
		Templates: templates,
	}
}
