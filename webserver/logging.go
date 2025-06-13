package main

import (
	"learn-go/v2/internal/interfaces"
	"log"
	"log/slog"
	"os"
)

var _ interfaces.LogUtil = (*Logger)(nil)

type Logger struct {
	logger      *slog.Logger
	ErrorLogger *log.Logger
}

func NewLogger() *Logger {
	handler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(handler)
	errorLogger := slog.NewLogLogger(handler, slog.LevelError)

	slog.SetDefault(logger)

	return &Logger{
		logger:      logger,
		ErrorLogger: errorLogger,
	}
}

func (l *Logger) Info(message string, data ...interfaces.LogData) {
	l.logger.Info(message, "data", data)
}

func (l *Logger) Debug(message string, data ...interfaces.LogData) {
	l.logger.Debug(message, "data", data)
}

func (l *Logger) Warn(message string, data ...interfaces.LogData) {
	l.logger.Warn(message, "data", data)
}

func (l *Logger) Error(message string, data ...interfaces.LogData) {
	l.logger.Error(message, "data", data)
}
