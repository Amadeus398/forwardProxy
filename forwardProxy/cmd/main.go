package main

import (
	"forwardProxy/pkg/config"
	"forwardProxy/pkg/handler"
	"forwardProxy/pkg/logging"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"net/http"
)

func main() {
	// Initialization loggers for
	// the main module and
	// the file closing function
	//for writing loggers.
	loggers, cleanup := logging.NewLogs("main", "cmd")
	defer cleanup()
	loggers.GetInfo().Msg("starting forwardProxy")

	cache := &config.EnvCache{}
	if err := envconfig.Process("", cache); err != nil {
		loggers.GetFatal().Err(err).Msg("unable to parse the environment")
	}

	zerolog.SetGlobalLevel(cache.GetLogLevel())

	proxy := http.Server{
		Addr:    cache.Port,
		Handler: handler.MyHandler{},
	}

	if err := proxy.ListenAndServe(); err != nil {
		defer loggers.GetFatal().Str("when", "start proxy").
			Err(err).Msg("exiting")
	}

}
