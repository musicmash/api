package clients

import "net/http"

type Provider struct {
	URL    string
	Client *http.Client
}

func NewProvider(url string) *Provider {
	return &Provider{
		URL:    url,
		Client: &http.Client{},
	}
}
