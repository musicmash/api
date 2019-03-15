package auth

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/musicmash/api/internal/api/middleware"
	"github.com/musicmash/api/internal/log"
)

func NewMiddleware(authorizer Authorizer) middleware.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get(HeaderToken)
			if _, err := uuid.Parse(token); err != nil {
				log.Debugf("can't parse uuid '%s'", token)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			userName, err := authorizer.Authorize(token)
			if err != nil {
				log.Debugf("can't find session with provided token '%s'", token)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			setUserName(r, userName)
			next.ServeHTTP(w, r)
		})
	}
}
