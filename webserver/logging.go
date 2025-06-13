package main

import (
	"learn-go/v2/internal/interfaces"
	"log"
	"log/slog"
	"os"
)

var _ interfaces.LogUtil = (*logging)(nil)

type logging struct {
	logger      *slog.Logger
	ErrorLogger *log.Logger
}

func NewLogger() *logging {
	handler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(handler)
	errorLogger := slog.NewLogLogger(handler, slog.LevelError)

	slog.SetDefault(logger)

	return &logging{
		logger:      logger,
		ErrorLogger: errorLogger,
	}
}

func (l *logging) Info(message string, data ...interfaces.LogData) {
	l.logger.Info(message, "data", data)
}

func (l *logging) Debug(message string, data ...interfaces.LogData) {
	l.logger.Debug(message, "data", data)
}

func (l *logging) Warn(message string, data ...interfaces.LogData) {
	l.logger.Warn(message, "data", data)
}

func (l *logging) Error(message string, data ...interfaces.LogData) {
	l.logger.Error(message, "data", data)
}
