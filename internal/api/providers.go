package api

import (
	"github.com/musicmash/api/internal/clients"
	"github.com/musicmash/api/internal/config"
)

var (
	feedProvider          *clients.Provider
	artistsProvider       *clients.Provider
	subscriptionsProvider *clients.Provider
	usersProvider         *clients.Provider
)

func InitProviders() {
	feedProvider = clients.NewProvider(config.Config.Services.Artists)
	artistsProvider = clients.NewProvider(config.Config.Services.Artists)
	subscriptionsProvider = clients.NewProvider(config.Config.Services.Subscriptions)
	usersProvider = clients.NewProvider(config.Config.Services.Users)
}
