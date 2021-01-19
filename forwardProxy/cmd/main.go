package main

import (
	"forwardProxy/pkg/config"
	"forwardProxy/pkg/handler"
	"forwardProxy/pkg/logging"
	"github.com/kelseyhightower/envconfig"
	"net/http"
)

func main() {
	loggers := logging.NewLogs("main", "cmd")

	loggers.GetInfo().Msg("starting forwardProxy")
	cache := &config.EnvCache{}
	if err := envconfig.Process("", cache); err != nil {
		loggers.GetFatal().Err(err).Msg("unable to parse the environment")
		panic("unable to parse the environment")
	}

	proxy := http.Server{
		Addr:    cache.Port,
		Handler: handler.MyHandler{},
	}

	if err := proxy.ListenAndServe(); err != nil {
		defer loggers.GetFatal().Str("when", "start proxy").
			Err(err).Msg("exiting")
		panic("The output from the program, shutdown")
	}

}
