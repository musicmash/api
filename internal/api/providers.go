package api

import (
	"github.com/musicmash/api/internal/clients"
	"github.com/musicmash/api/internal/config"
	"github.com/musicmash/auth/pkg/api"
)

var (
	feedProvider          *clients.Provider
	artistsProvider       *clients.Provider
	subscriptionsProvider *clients.Provider
	usersProvider         *clients.Provider
	authProvider          *api.Provider
)

func InitProviders() {
	feedProvider = clients.NewProvider(config.Config.Services.Artists)
	artistsProvider = clients.NewProvider(config.Config.Services.Artists)
	subscriptionsProvider = clients.NewProvider(config.Config.Services.Subscriptions)
	usersProvider = clients.NewProvider(config.Config.Services.Users)
	authProvider = api.NewProvider(config.Config.Services.Auth)
}
