package logger

import (
	"go.uber.org/zap"
)

// Logger is a wrapper type for zap.Logger
type Logger *zap.Logger

// NewLogger returns a new logger with type (Production/Development) based on
func NewLogger(config *AppConfig) (Logger, error) {

	var lg *zap.Logger

	switch config.IsProduction {
	case true:
		if lg, err := zap.NewProduction(); err != nil {
			return nil, err
		}
	case false:
		if lg, err := zap.NewDevelopment(); err != nil {
			return nil, err
		}
	}

	defer lg.Sync()
	return lg, nil
}
