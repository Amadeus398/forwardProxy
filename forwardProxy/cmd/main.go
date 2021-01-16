package main

import (
	"forwardProxy/pkg/handler"
	"forwardProxy/pkg/logging"
	"net/http"
)

var (
	funcFatalLog = logging.Logs{}.GetFatal()
	funcInfoLog  = logging.Logs{}.GetInfo()
)

func main() {
	funcInfoLog.Str("module", "main").Msg("starting forwardProxy")

	proxy := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: handler.MyHandler{},
	}

	if err := proxy.ListenAndServe(); err != nil {
		defer funcFatalLog.Str("method", "main").Str("when", "start proxy").
			Err(err).Msg("exiting")
		panic("все пошло по пизде")
	}

}
