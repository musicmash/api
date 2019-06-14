package api

import (
	"github.com/musicmash/api/internal/clients"
	"github.com/musicmash/api/internal/config"
)

var (
	feedProvider *clients.Provider
)

func InitProviders() {
	feedProvider = clients.NewProvider(config.Config.Services.Artists)
}
