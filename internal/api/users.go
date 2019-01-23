package api

import (
	"net/http"
	"strings"

	"github.com/musicmash/api/internal/clients/users"
)

func IsUserExits(w http.ResponseWriter, name string) error {
	if strings.TrimSpace(name) == "" {
		return users.ErrUserNotFound
	}

	return users.Get(usersProvider, name)
}
