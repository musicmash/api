package api

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/musicmash/auth/pkg/api/info"
)

const (
	HeaderToken    = "x-auth-token"
	HeaderUserName = "user_name"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := w.Header().Get(HeaderToken)
		if _, err := uuid.Parse(token); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		session, err := info.Get(authProvider, token)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		setUserName(r, session.UserName)
		next.ServeHTTP(w, r)
	})
}

func getUserName(r *http.Request) string {
	return r.Header.Get(HeaderUserName)
}

func setUserName(r *http.Request, userName string) {
	r.Header.Set(HeaderUserName, userName)
}
