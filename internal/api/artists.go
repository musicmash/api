package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/musicmash/api/internal/clients/artists"
	"github.com/musicmash/api/internal/log"
)

func searchArtist(w http.ResponseWriter, r *http.Request) {
	userName := chi.URLParam(r, "user_name")
	if err := IsUserExits(w, userName); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	artistName := strings.TrimSpace(r.URL.Query().Get("name"))
	if artistName == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	artists, err := artists.Search(artistsProvider, userName, artistName)
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
	w.Write(buffer)
}

func getArtistDetails(w http.ResponseWriter, r *http.Request) {
	userName := chi.URLParam(r, "user_name")
	if err := IsUserExits(w, userName); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	artistName := chi.URLParam(r, "artist_name")
	if artistName == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	artists, err := artists.GetDetails(artistsProvider, userName, artistName)
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
	w.Write(buffer)
}
