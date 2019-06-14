package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/musicmash/api/internal/log"
	"github.com/musicmash/artists/pkg/api"
	"github.com/musicmash/artists/pkg/api/search"
)

const PathArtists = "/artists"

type ArtistsController struct {
	provider *api.Provider
}

func NewArtistsController(artistsProvider *api.Provider) *ArtistsController {
	return &ArtistsController{provider: artistsProvider}
}

func (a *ArtistsController) Register(router chi.Router) {
	router.Route(PathArtists, func(r chi.Router) {
		r.Get("/search", a.doSearch)
	})
}

func (a *ArtistsController) doSearch(w http.ResponseWriter, r *http.Request) {
	artistName := r.URL.Query().Get("name")
	if len(artistName) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	artists, err := search.Do(a.provider, artistName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error(err)
		return
	}

	buffer, err := json.Marshal(&artists)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error(err)
		return
	}
	w.Header().Set("content-type", "application/json")
	_, _ = w.Write(buffer)
}
