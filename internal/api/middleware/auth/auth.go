package auth

import (
	"github.com/musicmash/api/internal/log"
	"github.com/musicmash/auth/pkg/api"
	"github.com/musicmash/auth/pkg/api/info"
	"github.com/pkg/errors"
)

var ErrNotAuthorized = errors.New("user not authorized")

type Authorizer interface {
	Authorize(token string) (username string, err error)
}

type MusicAuthorizer struct {
	Provider *api.Provider
}

func NewAuthorizer(provider *api.Provider) Authorizer {
	return &MusicAuthorizer{Provider: provider}
}

func (m *MusicAuthorizer) Authorize(token string) (username string, err error) {
	session, err := info.Get(m.Provider, token)
	if err != nil {
		log.Debugf("can't find session with provided token '%s'", token)
		return "", ErrNotAuthorized
	}
	return session.UserName, nil
}
