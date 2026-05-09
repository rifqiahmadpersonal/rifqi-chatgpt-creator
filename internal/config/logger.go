package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

func NewLogger(cfg *LoggingConfig) *logrus.Logger {
	logger := logrus.New()

	level, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		level = logrus.DebugLevel
	}
	logger.SetLevel(level)

	if cfg.Format == "json" {
		logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02T15:04:05.000Z07:00",
		})
	} else {
		logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
		})
	}

	if cfg.Output == "stderr" {
		logger.SetOutput(os.Stderr)
	} else {
		logger.SetOutput(os.Stdout)
	}

	return logger
}
