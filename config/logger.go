package config

import (
	"context"
	"os"
	"sync/atomic"

	"github.com/sirupsen/logrus"
)

type customTextFormatter struct {
	logrus.TextFormatter
}

var logger *logrus.Logger
var currentMode atomic.Value

// Carrega o logger
func loadLogger(ctx context.Context) *logrus.Logger {
	mode := ctx.Value(aplicationModeKey).(applicationMode)

	currentMode.Store(mode)

	logger = logrus.New()
	
	logger.SetFormatter(&customTextFormatter{
		TextFormatter: logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "15:04:05.000",
			ForceColors:     true,
		},
	})

	logger.SetOutput(os.Stdout)

	return logger
}
