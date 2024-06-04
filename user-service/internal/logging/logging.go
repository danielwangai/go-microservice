package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

const (
	defaultTimeStampFormat = "01-02-2006 15:04:05"
	defaultLogLevel        = "info"
)

func NewLogger() *logrus.Logger {
	level, _ := logrus.ParseLevel(defaultLogLevel)
	logger := &logrus.Logger{
		Out:          os.Stdout,
		Level:        level,
		ReportCaller: true,
	}
	return logger
}

// SetJSONFormatter overrides the default formatter so as to
// support logging with json
func SetJSONFormatter(log *logrus.Logger) *logrus.Logger {
	log.Formatter = &logrus.JSONFormatter{
		TimestampFormat: defaultTimeStampFormat,
	}
	return log
}
