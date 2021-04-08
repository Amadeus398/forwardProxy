package logging

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
	"os"
	"sync"
	"time"
)

const logFilename = "logging.txt"

var (
	once    sync.Once
	logFile *os.File
)

type Logs struct {
	infoLog  *zerolog.Event
	errorLog *zerolog.Event
	warnLog  *zerolog.Event
	fatLog   *zerolog.Event
	method   string
	module   string
}

// NewLogs returns a method and a module,
// depending on which module it is called in.
//
// once.Do joins 2 io.Writer and writes
// logger messages to the file
// and to the os.Stdout.
//
// the file closing function
// must be called in the
// module main only.
func NewLogs(method, module string) (logs *Logs, f func()) {
	once.Do(func() {
		var err error
		logFile, err = os.OpenFile(logFilename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			fmt.Println("failed to open file")
			return
		}
		colorLogger := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
		log.Logger = log.Output(io.MultiWriter(zerolog.SyncWriter(logFile), colorLogger))
	})

	return &Logs{method: method, module: module}, cleanup
}

// cleanup closes file *os.File
func cleanup() {
	if err := logFile.Close(); err != nil {
		panic(err)
	}
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
