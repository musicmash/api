package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/musicmash/api/internal/log"
)

func getMux() *chi.Mux {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/auth", authUser)

	r.Route("/feed", func(r chi.Router) {
		r.Use(Auth)
		r.Get("/", getUserFeed)
	})

	r.Route("/subscriptions", func(r chi.Router) {
		r.Use(Auth)
		r.Get("/", getUserSubscriptions)
		r.Post("/", createSubscriptions)
		r.Delete("/", deleteSubscriptions)
	})

	r.Route("/artists", func(r chi.Router) {
		r.Use(Auth)
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
