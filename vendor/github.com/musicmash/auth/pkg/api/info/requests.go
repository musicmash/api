package info

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/musicmash/auth/pkg/api"
)

func Get(provider *api.Provider, token string) (*Session, error) {
	url := fmt.Sprintf("%s/info/%s", provider.URL, token)
	resp, err := provider.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("got %d status code", resp.StatusCode)
	}

	session := Session{}
	if err := json.NewDecoder(resp.Body).Decode(&session); err != nil {
		return nil, err
	}
	return &session, nil
}
