package main

import (
	"flag"

	raven "github.com/getsentry/raven-go"
	"github.com/musicmash/api/internal/api"
	"github.com/musicmash/api/internal/config"
	"github.com/musicmash/api/internal/log"
	"github.com/pkg/errors"
)

func main() {
	configPath := flag.String("config", "/etc/musicmash-api/musicmash-api.yaml", "Path to musicmash-api.yaml config")
	flag.Parse()

	if err := config.InitConfig(*configPath); err != nil {
		panic(err)
	}
	if config.Config.Log.Level == "" {
		config.Config.Log.Level = "info"
	}

	log.SetLogFormatter(&log.DefaultFormatter)
	log.ConfigureStdLogger(config.Config.Log.Level)

	if config.Config.Sentry.Enabled {
		if err := raven.SetDSN(config.Config.Sentry.Key); err != nil {
			panic(errors.Wrap(err, "tried to setup sentry client"))
		}
	}

	log.Info("Running musicmash-api..")
	api.InitProviders()
	log.Panic(api.ListenAndServe(config.Config.HTTP.IP, config.Config.HTTP.Port))
}
