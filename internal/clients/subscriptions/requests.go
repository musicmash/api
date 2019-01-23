package subscriptions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/musicmash/api/internal/clients"
)

func Get(provider *clients.Provider, userName string) ([]*Subscription, error) {
	url := fmt.Sprintf("%s/%s/subscriptions", provider.URL, userName)
	resp, err := provider.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	subscriptions := []*Subscription{}
	if err := json.NewDecoder(resp.Body).Decode(&subscriptions); err != nil {
		return nil, err
	}
	return subscriptions, nil
}

func Subscribe(provider *clients.Provider, userName string, artists []string) error {
	url := fmt.Sprintf("%s/%s/subscriptions", provider.URL, userName)
	b, err := json.Marshal(&artists)
	if err != nil {
		return err
	}

	resp, err := provider.Client.Post(url, "application/json", bytes.NewReader(b))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusAccepted {
		return nil
	}
	return fmt.Errorf("wanna 202 status code, but got %v", resp.StatusCode)
}

func UnSubscribe(provider *clients.Provider, userName string, artists []string) error {
	url := fmt.Sprintf("%s/%s/subscriptions", provider.URL, userName)
	b, err := json.Marshal(&artists)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodDelete, url, bytes.NewReader(b))
	if err != nil {
		return err
	}

	resp, err := provider.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return nil
	}
	return fmt.Errorf("wanna 200 status code, but got %v", resp.StatusCode)
}
