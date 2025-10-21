package config

import (
	"context"
	"crypto/rand"
	"log"
	"net/http"
	"os"

	"github.com/boj/redistore"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Config struct {
	OAuthConfig *oauth2.Config
	Session     sessions.Store
	Environment string
	SessionKey  []byte
	DB          *pgx.Conn
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
}

func LoadDatabase() *pgx.Conn {
	connection, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal("Unable to connect to PostgreSQL: ", err)
	}
	return connection
}

func Load() *Config {
	environment := os.Getenv("ENVIRONMENT")

	oAuthConfig := &oauth2.Config{
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}

	sessionKey := make([]byte, 32)
	if _, err := rand.Read(sessionKey); err != nil {
		log.Fatal("Failed to generate session key: ", err)
	}

	store, err := redistore.NewRediStore(10, "tcp", os.Getenv("REDIS_ADDRESS"), os.Getenv("REDIS_USERNAME"), os.Getenv("REDIS_PASSWORD"), sessionKey)
	if err != nil {
		log.Fatal("Error creating Redis Store: ", err)
	}

	secureCookies := environment == "production"

	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
		Secure:   secureCookies,
		SameSite: http.SameSiteLaxMode,
	}

	databaseConnection := LoadDatabase()

	return &Config{
		OAuthConfig: oAuthConfig,
		Session:     store,
		Environment: environment,
		SessionKey:  sessionKey,
		DB:          databaseConnection,
	}
}
