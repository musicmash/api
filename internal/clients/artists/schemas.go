package artists

import "github.com/musicmash/api/internal/clients/releases"

type Artist struct {
	Name   string `json:"name"`
	Poster string `json:"poster"`
}

type ArtistInfo struct {
	Artist
	Stores   []*ArtistStoreInfo     `json:"stores"`
	Releases *ArtistDetailsReleases `json:"releases"`
}

type ArtistStoreInfo struct {
	StoreName string `json:"name"`
	StoreID   string `json:"id"`
}

type ArtistDetailsReleases struct {
	Announced []*releases.Release `json:"announced"`
	Recent    []*releases.Release `json:"released"`
}
