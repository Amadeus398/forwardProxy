package logging

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Logs struct {
	infoLog  *zerolog.Event
	errorLog *zerolog.Event
	warnLog  *zerolog.Event
	fatLog   *zerolog.Event
	method   string
	module   string
}

func NewLogs(method, module string) *Logs {
	return &Logs{method: method, module: module}
}

func (l Logs) addMetadata(e *zerolog.Event) *zerolog.Event {
	return e.Str("module", l.module).Str("method", l.method)
}

func (l Logs) GetInfo() *zerolog.Event {
	if l.infoLog == nil {
		l.infoLog = l.addMetadata(log.Info())
	}
	return l.infoLog
}

func (l Logs) GetError() *zerolog.Event {
	if l.errorLog == nil {
		l.errorLog = l.addMetadata(log.Error())
	}
	return l.errorLog
}

func (l Logs) GetWarn() *zerolog.Event {
	if l.warnLog == nil {
		l.warnLog = l.addMetadata(log.Warn())
	}
	return l.warnLog
}

func (l Logs) GetFatal() *zerolog.Event {
	if l.fatLog == nil {
		l.fatLog = l.addMetadata(log.Fatal())
	}
	return l.fatLog
}
