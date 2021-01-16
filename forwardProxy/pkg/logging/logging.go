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
}

func (l Logs) GetInfo() *zerolog.Event {
	if l.infoLog == nil {
		l.infoLog = &zerolog.Event{}
	}
	l.infoLog = log.Info()
	return l.infoLog
}

func (l Logs) GetError() *zerolog.Event {
	if l.errorLog == nil {
		l.errorLog = &zerolog.Event{}
	}
	l.errorLog = log.Error()
	return l.errorLog
}

func (l Logs) GetWarn() *zerolog.Event {
	if l.warnLog == nil {
		l.warnLog = &zerolog.Event{}
	}
	l.warnLog = log.Warn()
	return l.warnLog
}

func (l Logs) GetFatal() *zerolog.Event {
	if l.fatLog == nil {
		l.fatLog = &zerolog.Event{}
	}
	l.fatLog = log.Fatal()
	return l.fatLog
}
