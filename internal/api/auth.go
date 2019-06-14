package api

import (
	"encoding/json"
	"net/http"

	"github.com/musicmash/api/internal/log"
	"github.com/musicmash/auth/pkg/api/auth"
)

func authUser(w http.ResponseWriter, r *http.Request) {
	payload := auth.Payload{}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if payload.Service != "google" || payload.Token == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO (m.kalinin): handle raw response and write body directly
	token, err := auth.Auth(authProvider, &payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := json.Marshal(&token)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	_, _ = w.Write(body)
}
