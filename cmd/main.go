package main

import (
	"context"
	"fmt"
	"postgresTakeWords/internal/apperrors"
	"postgresTakeWords/internal/competition"
	"postgresTakeWords/internal/config"
	"postgresTakeWords/internal/infrastructure/database"
	"postgresTakeWords/internal/log"

	_ "github.com/lib/pq"
)

const (
	stat     = "statistic.txt"
	libJson  = "save/library.json"
	libTxt   = "save/library.txt"
	newWords = "save/newWords.txt"
)

func main() {
	logger, err := log.NewLogAndSetLevel("info")
	if err != nil {
		logger.Fatal(err)
	}

	conf := config.NewConfig()
	err = conf.ParseConfig("config/.env", logger)
	if err != nil {
		logger.Fatal(apperrors.EnvConfigLoadError.AppendMessage(err))
	}

	if err = log.SetLevel(logger, conf.LogLevel); err != nil {
		logger.Fatal(err)
	}

	ctx := context.Background()
	psqlDB, err := database.InitClientPostgress(ctx, conf, logger)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println("Conf connect")

	compet := competition.NewCompetition(stat, libJson, libTxt, newWords, psqlDB, logger)
	err = compet.StartCompetition()
	if err != nil {
		logger.Fatal(err)
	}
}
