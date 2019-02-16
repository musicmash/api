package api

import (
	"encoding/json"
	"net/http"

	"github.com/musicmash/api/internal/clients/subscriptions"
)

func createSubscriptions(w http.ResponseWriter, r *http.Request) {
	userName := getUserName(r)
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
	userName := getUserName(r)
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
	userName := getUserName(r)
	userSubscriptions, err := subscriptions.Get(subscriptionsProvider, userName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	buffer, err := json.Marshal(&userSubscriptions)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(buffer)
}
