package token

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/musicmash/auth/pkg/api"
)

func GetDetails(provider *api.Provider, token string) (*Session, error) {
	url := fmt.Sprintf("%s/token/%s", provider.URL, token)
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

func Delete(provider *api.Provider, token string) error {
	url := fmt.Sprintf("%s/token/%s", provider.URL, token)
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	resp, err := provider.Client.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("got %d status code", resp.StatusCode)
	}
	return nil
}
