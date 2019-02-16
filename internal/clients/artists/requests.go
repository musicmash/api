package artists

import (
	"encoding/json"
	"fmt"
	"net/http"

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

	if resp.StatusCode >= http.StatusBadRequest {
		if resp.StatusCode == http.StatusNotFound {
			return nil, ErrArtistNotFound
		}
		return nil, fmt.Errorf("got %d status code from musicmash/artists", resp.StatusCode)
	}

	details := ArtistInfo{}
	if err := json.NewDecoder(resp.Body).Decode(&details); err != nil {
		return nil, err
	}
	return &details, nil
}
