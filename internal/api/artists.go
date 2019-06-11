package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/musicmash/api/internal/clients/artists"
	"github.com/musicmash/api/internal/log"
)

func searchArtist(w http.ResponseWriter, r *http.Request) {
	userName := GetUserName(r)
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
