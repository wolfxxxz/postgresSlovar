package config

import (
	"fmt"
	"postgresTakeWords/internal/apperrors"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Token        string `env:"TOKEN"`
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
		errMsg := fmt.Sprintf(" %s", err.Error())
		log.Info("gotoenv could not find .env", errMsg)
		return apperrors.EnvConfigParseError.AppendMessage(errMsg)

	}

	if err := env.Parse(v); err != nil {
		errMsg := fmt.Sprintf("%+v\n", err)
		return apperrors.EnvConfigParseError.AppendMessage(errMsg)
	}

	log.Info("Config has been parsed, succesfully!!!")
	return nil
}
