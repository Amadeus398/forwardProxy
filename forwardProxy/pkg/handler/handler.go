package handler

import (
	"context"
	"fmt"
	"forwardProxy/pkg/logging"
	"io/ioutil"
	"net/http"
)

const message = "If you see this page, an error has occurred"

type MyHandler struct {
	c *http.Client
	l *logging.Logs
}

func (m MyHandler) getLoggers() *logging.Logs {
	if m.l == nil {
		m.l, _ = logging.NewLogs("serveHTTP", "handler")
	}
	return m.l
}

func (m MyHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	m.getLoggers().GetInfo().Str("when", "starting processing of request").Str("url", request.RequestURI).
		Msg("started handler")

	ctx := context.TODO()
	myRequest := request.Clone(ctx)
	myRequest.RequestURI = ""

	resp, err := m.getClient().Do(myRequest)
	if err != nil {
		m.getLoggers().GetError().Str("when", "completed request, start response").
			Str("url", request.RequestURI).Msg("unable to get response")

		writer.Header().Set("Content-type", "text/plain; charset=utf-8")
		bytesMonkeys := []byte(message)
		writer.Header().Set("Content-Length", fmt.Sprintf("%d", len(bytesMonkeys)))
		writer.WriteHeader(http.StatusInternalServerError)
		if _, err := writer.Write(bytesMonkeys); err != nil {
			m.getLoggers().GetError().Str("when", "sending response").Msg("can't send response to userName")
			return
		}
		return
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			m.getLoggers().GetError().Str("when", "close body").
				Msg("unable to close Body")
			return
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		m.getLoggers().GetError().Str("when", "reading response body").Msg("unable to read Body")
		return
	}
	for header, headerVal := range resp.Header {
		for _, headerValue := range headerVal {
			writer.Header().Set(header, headerValue)
		}
	}
	writer.WriteHeader(resp.StatusCode)
	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		m.getLoggers().GetWarn().Str("when", "status code response").Msg("status code 4xx")
	} else if resp.StatusCode >= 500 {
		m.getLoggers().GetWarn().Str("when", "status code response").Msg("status code 5xx")
	}

	if _, err := writer.Write(body); err != nil {
		m.getLoggers().GetError().Str("when", "writing body").Msg("unable to write body")
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
