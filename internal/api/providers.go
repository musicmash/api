package api

import (
	"github.com/musicmash/api/internal/clients"
	"github.com/musicmash/api/internal/config"
	"github.com/musicmash/auth/pkg/api"
)

var (
	feedProvider *clients.Provider
	authProvider *api.Provider
)

func InitProviders() {
	feedProvider = clients.NewProvider(config.Config.Services.Artists)
	authProvider = api.NewProvider(config.Config.Services.Auth)
}
