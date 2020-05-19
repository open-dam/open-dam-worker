package opendam

import (
	"os"

	"github.com/sirupsen/logrus"
)

const defaultLogLevel = logrus.InfoLevel

func Logger() *logrus.Entry {
	logger := logrus.New()

	logLevel := defaultLogLevel
	if envLevel, ok := os.LookupEnv("LOG_LEVEL"); ok {
		reqLogLevel, err := logrus.ParseLevel(envLevel)
		if err == nil {
			logLevel = reqLogLevel
		}
	}
	logger.SetLevel(logLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})

	hostname, _ := os.Hostname()
	environment := os.Getenv("ENVIRONMENT")

	return logger.WithFields(logrus.Fields{
		"host":        hostname,
		"environment": environment,
	})
}
