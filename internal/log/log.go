package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

func NewLogAndSetLevel(logLevel string) (*logrus.Logger, error) {
	logger := logrus.New()
	loggerLevel, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return nil, err
	}

	logger.SetLevel(loggerLevel)
	logger.SetReportCaller(true)
	logger.SetOutput(os.Stdout)
	logger.Info("Logger has been configurated")
	return logger, nil
}

func SetLevel(log *logrus.Logger, logLevel string) error {
	loggerLevel, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return err
	}

	log.SetLevel(loggerLevel)
	log.Info("logger level has been configurated")
	return nil
}
