package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/musicmash/api/internal/config"
	"github.com/musicmash/api/internal/log"
	"github.com/musicmash/subscriptions/pkg/api"
)

func getMux() *chi.Mux {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/healthz", healthz)

	r.Post("/auth", authUser)

	r.Route("/token", func(r chi.Router) {
		r.Delete("/", deleteToken)
	})

	r.Route("/feed", func(r chi.Router) {
		r.Get("/", getUserFeed)
	})

	r.Route("/v1", func(r chi.Router) {
		NewSubscriptionsController(api.NewProvider(config.Config.Services.Subscriptions, 1)).Register(r)
	})

	r.Route("/artists", func(r chi.Router) {
		r.Get("/", searchArtist)
		r.Get("/{artist_name}", getArtistDetails)
	})
	return r
}

func ListenAndServe(ip string, port int) error {
	addr := fmt.Sprintf("%s:%d", ip, port)
	log.Infof("Listening API on '%s'", addr)
	return http.ListenAndServe(addr, getMux())
}
