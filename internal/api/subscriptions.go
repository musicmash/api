package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/musicmash/api/internal/log"
	artsapi "github.com/musicmash/artists/pkg/api"
	"github.com/musicmash/artists/pkg/api/artists"
	subsapi "github.com/musicmash/subscriptions/pkg/api"
	"github.com/musicmash/subscriptions/pkg/api/subscriptions"
)

const PathSubscriptions = "/subscriptions"

type SubscriptionsController struct {
	subscriptionsProvider *subsapi.Provider
	artistsProvider       *artsapi.Provider
}

func NewSubscriptionsController(subscriptionsProvider *subsapi.Provider, artistsProvider *artsapi.Provider) *SubscriptionsController {
	return &SubscriptionsController{
		artistsProvider:       artistsProvider,
		subscriptionsProvider: subscriptionsProvider,
	}
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

	validatedArtists, err := artists.Validate(s.artistsProvider, userArtists)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(validatedArtists) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := subscriptions.Create(s.subscriptionsProvider, userName, validatedArtists); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *SubscriptionsController) deleteSubscriptions(w http.ResponseWriter, r *http.Request) {
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

	if err := subscriptions.Delete(s.subscriptionsProvider, userName, userArtists); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *SubscriptionsController) getUserSubscriptions(w http.ResponseWriter, r *http.Request) {
	userName := GetUserName(r)
	userSubscriptions, err := subscriptions.Get(s.subscriptionsProvider, userName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(userSubscriptions) == 0 {
		w.Header().Set("content-type", "application/json")
		w.Write([]byte(`[]`))
		return
	}

	artists, err := artists.GetFullInfo(s.artistsProvider, userSubscriptions)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	buffer, err := json.Marshal(&artists)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(buffer)
}
