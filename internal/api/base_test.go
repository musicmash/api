package api

import (
	"net/http/httptest"

	"github.com/musicmash/api/internal/config"
	"github.com/musicmash/api/internal/testutils"
	"github.com/musicmash/api/pkg/api"
)

var (
	server *httptest.Server
	env    *testutils.Env
	client *api.Provider
)

func setup() {
	env = testutils.Setup()
	config.Config = &config.AppConfig{
		Services: config.Services{
			Artists:       env.Server.URL,
			Subscriptions: env.Server.URL,
		},
	}
	server = httptest.NewServer(getMux())
	client = api.NewProvider(server.URL, 1)
}

func teardown() {
	env.TearDown()
	server.Close()
}
