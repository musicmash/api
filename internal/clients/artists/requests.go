package artists

import (
	"encoding/json"
	"fmt"

	"github.com/musicmash/api/internal/clients"
)

func Search(provider *clients.Provider, userName, artistName string) ([]*Artist, error) {
	url := fmt.Sprintf("%s/%s/artists?name=%s", provider.URL, userName, artistName)
	resp, err := provider.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	artists := []*Artist{}
	if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
		return nil, err
	}
	return artists, nil
}

func GetDetails(provider *clients.Provider, userName, artistName string) (*ArtistInfo, error) {
	url := fmt.Sprintf("%s/%s/artists/%s", provider.URL, userName, artistName)
	resp, err := provider.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	details := ArtistInfo{}
	if err := json.NewDecoder(resp.Body).Decode(&details); err != nil {
		return nil, err
	}
	return &details, nil
}
