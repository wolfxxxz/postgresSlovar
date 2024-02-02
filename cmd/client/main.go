package main

import (
	"context"
	"postgresTakeWords/internal/competition"
	"postgresTakeWords/internal/config"
	"postgresTakeWords/internal/database"
	"postgresTakeWords/internal/log"
	"postgresTakeWords/internal/models"
	"postgresTakeWords/internal/repositories"

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
		logger.Fatal(err)
	}

	if err = log.SetLevel(logger, conf.LogLevel); err != nil {
		logger.Fatal(err)
	}

	ctx := context.Background()

	db, err := database.NewPostgresDB().SetupDatabase(ctx, conf, logger)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("SETUP DATABASE SUCCESS")

	if !db.Migrator().HasTable(&models.Word{}) {
		err = db.AutoMigrate(&models.Word{})
		if err != nil {
			logger.Fatal(err)
		}

		logger.Info("TABLE words CREATED SUCCESS")

		backUpRepo := repositories.NewBackUpCopyRepo(libJson, libTxt, logger)
		words, err := backUpRepo.GetAllWordsFromBackUpXlsx()
		if err != nil {
			logger.Fatal(err)
		}

		logger.Infof("GET WORDS FROM XLSx SUCCESS %+v", len(words))

		repoWords := repositories.NewRepoWordsGorm(db, logger)
		err = repoWords.InsertWords(ctx, words)
		if err != nil {
			logger.Fatal(err)
		}

		logger.Info("INSERT WORDS IN DB SUCCESS")
	}

	if !db.Migrator().HasTable(&models.WordsLearn{}) {
		err = db.AutoMigrate(&models.WordsLearn{})
		if err != nil {
			logger.Fatal(err)
		}
	}

	logger.Info("CREATE TABLE words_learn SUCCESS")

	logger.Info("Postgres has been connected")

	compet := competition.NewCompetition(stat, libJson, libTxt, newWords, db, logger)
	err = compet.StartCompetition()
	if err != nil {
		logger.Fatal(err)
	}
}
