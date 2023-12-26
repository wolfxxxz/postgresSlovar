package repositories

import (
	"postgresTakeWords/internal/apperrors"
	"postgresTakeWords/internal/models"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

var collectionRepoLearn = "words_learn"

type RepoLearn struct {
	db         *sqlx.DB
	collection string
	log        *logrus.Logger
}

func NewRepoLearn(db *sqlx.DB, log *logrus.Logger) *RepoLearn {
	return &RepoLearn{db: db, collection: collectionRepoLearn, log: log}
}

func (rl *RepoLearn) InsertWordLearn(word *models.Word) error {
	_, err := rl.db.Exec("INSERT INTO words_learn (id, english, russian, theme) VALUES ($1, $2, $3, $4)",
		word.Id, word.English, word.Russian, word.Theme)
	if err != nil {
		appErr := apperrors.InsertWordLearnErr.AppendMessage(err, word)
		rl.log.Error(appErr)
		return appErr
	}

	return nil
}

func (rl *RepoLearn) InsertWordsLearn(words *[]models.Word) error {
	for _, v := range *words {
		err := rl.InsertWordLearn(&v)
		if err != nil {
			return err
		}
	}

	return nil
}

func (rl *RepoLearn) GetWordsLearn(quantity int) (*[]models.Word, error) {
	var words []models.Word
	err := rl.db.Select(&words, "SELECT * FROM words_learn limit $1", quantity)
	if err != nil {
		appErr := apperrors.GetWordsLearnErr.AppendMessage(err)
		rl.log.Error(appErr)
		return nil, appErr
	}

	return &words, nil
}

func (rl *RepoLearn) DeleteLearnWordsId(id int) error {
	_, err := rl.db.Exec("delete from words_learn where id = $1", id)
	if err != nil {
		appErr := apperrors.DeleteLearnWordsIdErr.AppendMessage(err)
		rl.log.Error(appErr)
		return appErr
	}

	return nil
}
