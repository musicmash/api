package api

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/musicmash/api/internal/api/validators"
)

func searchArtist(w http.ResponseWriter, r *http.Request) {
	userName := chi.URLParam(r, "user_name")
	if err := validators.IsUserExits(w, userName); err != nil {
		return
	}

	name := strings.TrimSpace(r.URL.Query().Get("name"))
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusHTTPVersionNotSupported)
	return
}

func getArtistDetails(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimSpace(chi.URLParam(r, "artist_name"))
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusHTTPVersionNotSupported)
	return
}
