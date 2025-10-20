package main

import (
	"github.com/themoroccandemonlist/themoroccandemonlist-web/internal/app"
	"github.com/themoroccandemonlist/themoroccandemonlist-web/internal/config"
	"github.com/themoroccandemonlist/themoroccandemonlist-web/internal/router"
)

func main() {
	config.LoadEnv()
	config.LoadDatabase()
	config := config.Load()
	application := app.New(config)

	router.New(config, application)
}
