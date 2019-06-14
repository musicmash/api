package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/musicmash/api/internal/config"
	"github.com/musicmash/api/internal/log"
	artsapi "github.com/musicmash/artists/pkg/api"
	subsapi "github.com/musicmash/subscriptions/pkg/api"
)

func getMux() *chi.Mux {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/healthz", healthz)

	r.Route("/feed", func(r chi.Router) {
		r.Get("/", getUserFeed)
	})

	r.Route("/v1", func(r chi.Router) {
		artistsProvider := artsapi.NewProvider(config.Config.Services.Artists, 1)
		subscriptionsProvider := subsapi.NewProvider(config.Config.Services.Subscriptions, 1)

		NewArtistsController(artistsProvider).Register(r)
		NewSubscriptionsController(subscriptionsProvider, artistsProvider).Register(r)
	})

	return r
}

func ListenAndServe(ip string, port int) error {
	addr := fmt.Sprintf("%s:%d", ip, port)
	log.Infof("Listening API on '%s'", addr)
	return http.ListenAndServe(addr, getMux())
}
