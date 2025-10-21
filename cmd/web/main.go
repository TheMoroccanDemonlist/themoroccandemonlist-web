package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/themoroccandemonlist/themoroccandemonlist-web/internal/app"
	"github.com/themoroccandemonlist/themoroccandemonlist-web/internal/config"
	"github.com/themoroccandemonlist/themoroccandemonlist-web/internal/router"
)

func main() {
	config.LoadEnv()
	config := config.Load()
	application := app.New(config)

	router := router.New(config, application)

	address := ":8080"
	if config.Environment == "production" {
		address = ":443"
	}

	server := &http.Server{
		Addr:    address,
		Handler: router,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		log.Println("Server is shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Close database connection
		if err := config.DB.Close(ctx); err != nil {
			log.Printf("Error closing DB: %v", err)
		}

		// Shutdown server gracefully
		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Shutdown error: %v", err)
		}
	}()

	log.Printf("Starting server on %s...", address)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server error: %v", err)
	}
}
