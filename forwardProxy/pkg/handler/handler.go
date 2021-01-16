package handler

import (
	"context"
	"fmt"
	"forwardProxy/pkg/logging"
	"io/ioutil"
	"net/http"
)

const monkeys = "40 тысяч обезьян в жопу сунули банан"

type MyHandler struct {
	c *http.Client
}

var (
	funcInfoLog = logging.Logs{}.GetInfo().Str("method", "serveHTTP")
	funcErrLog  = logging.Logs{}.GetError().Str("method", "serveHTTP")
)

func init() {
	logging.Logs{}.GetInfo().Str("module", "handler").Msg("starting handler")
	logging.Logs{}.GetError().Str("module", "handler").Msg("starting handler")
	//logging.Logs{}.GetWarn().Str("module", "handler").Msg("starting")
}

//
func (m MyHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	funcInfoLog.Str("when", "starting processing of request").Str("url", request.RequestURI).
		Msg("started")

	ctx := context.TODO()
	myRequest := request.Clone(ctx)
	myRequest.RequestURI = ""

	resp, err := m.getClient().Do(myRequest)
	if err != nil {
		funcErrLog.Str("when", "completed request, start response").Str("url", request.RequestURI).
			Msg("unable to get response")

		writer.Header().Set("Content-type", "text/plain; charset=utf-8")
		bytesMonkeys := []byte(monkeys)
		writer.Header().Set("Content-Length", fmt.Sprintf("%d", len(bytesMonkeys)))
		writer.WriteHeader(http.StatusInternalServerError)
		if _, err := writer.Write(bytesMonkeys); err != nil {
			funcErrLog.Str("when", "sending response").Msg("can't send response to user")
			return
		}
		return
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			funcErrLog.Str("when", "close body").
				Msg("still shit")
			return
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		funcErrLog.Str("when", "reading response body").Msg("unable to read Body")
		return
	}
	for header, headerVal := range resp.Header {
		for _, headerValue := range headerVal {
			writer.Header().Set(header, headerValue)
		}
	}
	writer.WriteHeader(resp.StatusCode)
	// TODO warnlog

	if _, err := writer.Write(body); err != nil {
		funcErrLog.Str("when", "writing body").Msg("unable to write body")
		return
	}

}

// Check the m.c, if m.c nil,
// to create &http.Client{}
func (m MyHandler) getClient() *http.Client {
	if m.c == nil {
		m.c = &http.Client{}
	}
	return m.c
}
