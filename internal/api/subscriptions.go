package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/musicmash/api/internal/clients/subscriptions"
)

func createSubscriptions(w http.ResponseWriter, r *http.Request) {
	userName := chi.URLParam(r, "user_name")
	if err := IsUserExits(w, userName); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userArtists := []string{}
	if err := json.NewDecoder(r.Body).Decode(&userArtists); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := subscriptions.Subscribe(subscriptionsProvider, userName, userArtists); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

func deleteSubscriptions(w http.ResponseWriter, r *http.Request) {
	userName := chi.URLParam(r, "user_name")
	if err := IsUserExits(w, userName); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userArtists := []string{}
	if err := json.NewDecoder(r.Body).Decode(&userArtists); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := subscriptions.UnSubscribe(subscriptionsProvider, userName, userArtists); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func getUserSubscriptions(w http.ResponseWriter, r *http.Request) {
	userName := chi.URLParam(r, "user_name")
	if err := IsUserExits(w, userName); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userSubscriptions, err := subscriptions.Get(subscriptionsProvider, userName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(&userSubscriptions)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(body)
}
