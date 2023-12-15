package config

import (
	"postgresTakeWords/internal/apperrors"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	LogLevel     string `env:"LOGGER_LEVEL"`
	SqlHost      string `env:"SQLHost"`
	SqlPort      string `env:"SQL_PORT"`
	SqlType      string `env:"SQL_TYPE"`
	SqlMode      string `env:"SQL_MODE"`
	UserName     string `env:"USER_NAME"`
	Password     string `env:"PASSWORD"`
	DBName       string `env:"DB_NAME"`
	TimeoutQuery string `env:"TIMEOUT_QUERY"`
}

func NewConfig() *Config {
	return &Config{}
}

func (v *Config) ParseConfig(path string, log *logrus.Logger) error {
	err := godotenv.Load(path)
	if err != nil {
		appErr := apperrors.EnvConfigLoadError.AppendMessage(err)
		log.Error(appErr)
		return appErr

	}

	if err := env.Parse(v); err != nil {
		appErr := apperrors.EnvConfigParseError.AppendMessage(err)
		log.Error(appErr)
		return appErr
	}

	log.Info("Config has been parsed")
	return nil
}
