package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/musicmash/api/internal/api/validators"
)

func getUserFeed(w http.ResponseWriter, r *http.Request) {
	userName := chi.URLParam(r, "user_name")
	if err := validators.IsUserExits(w, userName); err != nil {
		return
	}

	w.WriteHeader(http.StatusHTTPVersionNotSupported)
	return
}
