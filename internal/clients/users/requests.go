package users

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/musicmash/api/internal/clients"
)

func Get(provider *clients.Provider, userName string) error {
	url := fmt.Sprintf("%s/users/%s", provider.URL, userName)
	resp, err := provider.Client.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode >= http.StatusBadRequest {
		return ErrUserNotFound
	}
	if resp.StatusCode >= http.StatusInternalServerError {
		return errors.New("internal error")
	}
	return nil
}
