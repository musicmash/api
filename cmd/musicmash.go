package main

import (
	"flag"

	"github.com/musicmash/musicmash/internal/api"
	"github.com/musicmash/musicmash/internal/config"
	"github.com/musicmash/musicmash/internal/cron"
	"github.com/musicmash/musicmash/internal/db"
	"github.com/musicmash/musicmash/internal/log"
	tasks "github.com/musicmash/musicmash/internal/tasks/subscribe"
)

func init() {
	log.SetLogFormatter(&log.DefaultFormatter)
	configPath := flag.String("config", "/etc/musicmash/musicmash.yaml", "Path to musicmash.yaml config")
	logLevel := flag.String("log-level", "info", "log level {debug,info,warning,error}")
	flag.Parse()

	if err := config.InitConfig(*configPath); err != nil {
		panic(err)
	}

	if *logLevel != "info" || config.Config.Log.Level == "" {
		// Priority to command-line
		log.ConfigureStdLogger(*logLevel)
	} else {
		// Priority to config
		if config.Config.Log.Level != "" {
			log.ConfigureStdLogger(config.Config.Log.Level)
		}
	}

	tasks.InitWorkerPool()
	db.DbMgr = db.NewMainDatabaseMgr()
}

func main() {
	log.Info("Running musicmash..")
	go cron.Run()
	log.Panic(api.ListenAndServe(config.Config.HTTP.IP, config.Config.HTTP.Port))
}
