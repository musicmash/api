package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/musicmash/subscriptions/pkg/api"
	"github.com/musicmash/subscriptions/pkg/api/subscriptions"
)

const PathSubscriptions = "/subscriptions"

type SubscriptionsController struct {
	provider *api.Provider
}

func NewSubscriptionsController(subscriptionsProvider *api.Provider) *SubscriptionsController {
	return &SubscriptionsController{provider: subscriptionsProvider}
}

func (s *SubscriptionsController) Register(router chi.Router) {
	router.Route(PathSubscriptions, func(r chi.Router) {
		r.Get("/", s.getUserSubscriptions)
		r.Post("/", s.createSubscriptions)
		r.Delete("/", s.deleteSubscriptions)
	})
}

func (s *SubscriptionsController) createSubscriptions(w http.ResponseWriter, r *http.Request) {
	userName := GetUserName(r)
	userArtists := []int64{}
	if err := json.NewDecoder(r.Body).Decode(&userArtists); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(userArtists) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(userArtists) > 100 {
		userArtists = userArtists[0:100]
	}

	if err := subscriptions.Create(s.provider, userName, userArtists); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

func (s *SubscriptionsController) deleteSubscriptions(w http.ResponseWriter, r *http.Request) {
	userName := GetUserName(r)
	userArtists := []int64{}
	if err := json.NewDecoder(r.Body).Decode(&userArtists); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(userArtists) > 100 {
		userArtists = userArtists[0:100]
	}

	if err := subscriptions.Delete(s.provider, userName, userArtists); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *SubscriptionsController) getUserSubscriptions(w http.ResponseWriter, r *http.Request) {
	userName := GetUserName(r)
	userSubscriptions, err := subscriptions.Get(s.provider, userName)
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
