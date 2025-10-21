package handlers

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"

	"github.com/themoroccandemonlist/themoroccandemonlist-web/internal/app"
)

type Handler struct {
	App *app.App
}

func (handler *Handler) GoogleAuth(w http.ResponseWriter, r *http.Request) {
	stateBytes := make([]byte, 32)
	if _, err := rand.Read(stateBytes); err != nil {
		log.Printf("Failed to generate state: %v", err)
		http.Error(w, "Failed to generate state", http.StatusInternalServerError)
		return
	}
	state := base64.URLEncoding.EncodeToString(stateBytes)

	session, _ := handler.App.Config.Session.Get(r, "session")
	session.Values["oauth_state"] = state
	session.Save(r, w)

	url := handler.App.Config.OAuthConfig.AuthCodeURL(string(state))
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (handler *Handler) GoogleCallback(w http.ResponseWriter, r *http.Request) {
	session, _ := handler.App.Config.Session.Get(r, "session")
	storedState, _ := session.Values["oauth_state"].(string)
	if r.URL.Query().Get("state") != storedState {
		http.Error(w, "Invalid state parameter", http.StatusBadRequest)
		return
	}

	token, err := handler.App.Config.OAuthConfig.Exchange(context.Background(), r.URL.Query().Get("code"))
	if err != nil {
		http.Error(w, "Token exchange failed", http.StatusInternalServerError)
		return
	}

	client := handler.App.Config.OAuthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		http.Error(w, "Failed to get user info", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var user map[string]any
	json.NewDecoder(resp.Body).Decode(&user)
	session.Values["user"] = user
	session.Values["authenticated"] = true
	session.Save(r, w)

	http.Redirect(w, r, "/profile", http.StatusSeeOther)
}

func (handler *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	session, err := handler.App.Config.Session.Get(r, "session")
	if err != nil {
		http.Error(w, "Failed to retrieve session", http.StatusInternalServerError)
		return
	}

	session.Options.MaxAge = -1
	session.Values = make(map[any]any)

	if err := session.Save(r, w); err != nil {
		http.Error(w, "Failed to save session", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
