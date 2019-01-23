package feed

import (
	"time"

	"github.com/musicmash/api/internal/clients/releases"
)

type Feed struct {
	Date      time.Time           `json:"date"`
	Announced []*releases.Release `json:"announced"`
	Released  []*releases.Release `json:"released"`
}
