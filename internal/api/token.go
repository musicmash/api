package api

import (
	"net/http"

	"github.com/musicmash/api/internal/api/middleware/auth"
	"github.com/musicmash/api/internal/log"
	"github.com/musicmash/auth/pkg/api/token"
)

func deleteToken(w http.ResponseWriter, r *http.Request) {
	if err := token.Delete(authProvider, auth.GetToken(r)); err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
