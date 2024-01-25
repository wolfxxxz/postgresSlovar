package database

import (
	"context"
	"fmt"
	"postgresTakeWords/internal/apperrors"
	"postgresTakeWords/internal/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func InitClientPostgress(ctx context.Context, conf *config.Config, log *logrus.Logger) (*sqlx.DB, error) {
	connect := fmt.Sprintf("%v://%v:%v/%v?sslmode=%v&user=%v&password=%v&database=%v", conf.SqlType, conf.SqlHost,
		conf.SqlPort, conf.SqlType, conf.SqlMode, conf.UserName, conf.Password, conf.DBName)

	db, err := sqlx.Connect(conf.SqlType, connect)
	if err != nil {
		appErr := apperrors.InitPostgressErr.AppendMessage(err)
		log.Error(appErr)
		return nil, appErr
	}

	if err := db.Ping(); err != nil {
		appErr := apperrors.InitPostgressErr.AppendMessage(err)
		log.Error(appErr)
		return nil, appErr
	}

	log.Info("DB Postgres has been connected, DB.Ping success ")
	return db, nil
}
