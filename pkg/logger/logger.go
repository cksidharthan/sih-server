package logger

import (
	"fmt"
	"github.com/cksidharthan/sih-server/pkg/config"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
	_ "time/tzdata"
)

type UTCFormatter struct {
	logrus.Formatter
}

// Format - Creates the format of time in America/New York timezone for the logging timestamp
func (u UTCFormatter) Format(e *logrus.Entry) ([]byte, error) {
	// set nyc timezone
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		return nil, fmt.Errorf("error loading location: %w", err)
	}
	e.Time = e.Time.In(loc)
	return u.Formatter.Format(e)
}

// New - Creates a new Logger object to be used by the application
func New(envCfg *config.Config) *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&UTCFormatter{&logrus.TextFormatter{}})
	if envCfg.LogLevel == "debug" {
		logger.SetLevel(logrus.DebugLevel)
	} else if envCfg.LogLevel == "info" {
		logger.SetLevel(logrus.InfoLevel)
	} else {
		logger.SetLevel(logrus.WarnLevel)
	}

	logger.SetOutput(io.MultiWriter(os.Stdout))

	return logger
}
