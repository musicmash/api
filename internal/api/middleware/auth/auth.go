package auth

import (
	"github.com/musicmash/api/internal/log"
	"github.com/musicmash/auth/pkg/api"
	"github.com/musicmash/auth/pkg/api/token"
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

func (m *MusicAuthorizer) Authorize(uuid string) (username string, err error) {
	session, err := token.GetDetails(m.Provider, uuid)
	if err != nil {
		log.Debugf("can't find session with provided token '%s'", uuid)
		return "", ErrNotAuthorized
	}
	return session.UserName, nil
}
