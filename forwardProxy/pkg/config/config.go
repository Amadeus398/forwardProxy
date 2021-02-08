package config

import "github.com/rs/zerolog"

type EnvCache struct {
	Port     string `envconfig:"PORT" required:true`
	LogLevel string `envconfig:"LEVEL" required:true`
}

// GetLogLevel determines which level
// the LogLevel environment
// corresponds to.
func (c EnvCache) GetLogLevel() zerolog.Level {
	switch c.LogLevel {
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	default:
		return zerolog.Disabled
	}
}
