package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	custommiddleware "github.com/musicmash/api/internal/api/middleware"
	"github.com/musicmash/api/internal/api/middleware/auth"
	"github.com/musicmash/api/internal/log"
)

func getMux(authMiddleware custommiddleware.Middleware) *chi.Mux {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/healthz", healthz)

	r.Post("/auth", authUser)

	r.Route("/token", func(r chi.Router) {
		r.Use(authMiddleware)
		r.Delete("/", deleteToken)
	})

	r.Route("/feed", func(r chi.Router) {
		r.Use(authMiddleware)
		r.Get("/", getUserFeed)
	})

	r.Route("/subscriptions", func(r chi.Router) {
		r.Use(authMiddleware)
		r.Get("/", getUserSubscriptions)
		r.Post("/", createSubscriptions)
		r.Delete("/", deleteSubscriptions)
	})

	r.Route("/artists", func(r chi.Router) {
		r.Use(authMiddleware)
		r.Get("/", searchArtist)
		r.Get("/{artist_name}", getArtistDetails)
	})
	return r
}

func ListenAndServe(ip string, port int) error {
	authorizer := auth.NewAuthorizer(authProvider)
	authMiddleware := auth.NewMiddleware(authorizer)

	addr := fmt.Sprintf("%s:%d", ip, port)
	log.Infof("Listening API on '%s'", addr)
	return http.ListenAndServe(addr, getMux(authMiddleware))
}
