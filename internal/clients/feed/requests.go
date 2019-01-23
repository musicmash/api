package feed

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/musicmash/api/internal/clients"
)

type Options struct {
	Since *time.Time
}

func Get(provider *clients.Provider, userName string, opts *Options) (*Feed, error) {
	url := fmt.Sprintf("%s/%s/feed", provider.URL, userName)
	if opts != nil {
		if opts.Since != nil {
			url = url + "?since=" + opts.Since.Format("2006-01-02")
		}
	}
	resp, err := provider.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	feed := Feed{}
	if err := json.NewDecoder(resp.Body).Decode(&feed); err != nil {
		return nil, err
	}
	return &feed, nil
}
