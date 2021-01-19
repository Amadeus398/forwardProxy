package main

import (
	"forwardProxy/pkg/handler"
	"forwardProxy/pkg/logging"
	"net/http"
)

var (
	fatalLogger = logging.Logs{}.GetFatal().Str("module", "main")
	infoLogger  = logging.Logs{}.GetInfo().Str("module", "main")
)

func main() {
	infoLogger.Msg("starting forwardProxy")

	proxy := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: handler.MyHandler{},
	}

	if err := proxy.ListenAndServe(); err != nil {
		defer fatalLogger.Str("when", "start proxy").
			Err(err).Msg("exiting")
		panic("The output from the program, shutdown")
	}

}
