package logger

import (
	"log"
	"log/slog"
	"os"
)

var _ LogUtil = (*logging)(nil)

type logging struct {
	log         *slog.Logger
	errorLogger *log.Logger
}

func NewLogger() *logging {
	handler := slog.NewJSONHandler(os.Stdout, nil)
	_logger := slog.New(handler)
	errorLogger := slog.NewLogLogger(handler, slog.LevelError)

	slog.SetDefault(_logger)

	return &logging{
		log:         _logger,
		errorLogger: errorLogger,
	}
}

func (l *logging) Info(message string, data ...LogData) {
	l.log.Info(message, "data", data)
}

func (l *logging) Debug(message string, data ...LogData) {
	l.log.Debug(message, "data", data)
}

func (l *logging) Warn(message string, data ...LogData) {
	l.log.Warn(message, "data", data)
}

func (l *logging) Error(message string, data ...LogData) {
	l.log.Error(message, "data", data)
}
