package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/musicmash/auth/pkg/api"
)

func Auth(provider *api.Provider, payload *Payload) (*ServiceToken, error) {
	url := fmt.Sprintf("%s/auth", provider.URL)
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	resp, err := provider.Client.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("got %d status code", resp.StatusCode)
	}

	token := ServiceToken{}
	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
		return nil, err
	}
	return &token, nil
}
