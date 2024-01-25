package repositories

/*
var collectionRepoLearn = "words_learns"

type repoLearn struct {
	db         *sqlx.DB
	collection string
	log        *logrus.Logger
}

func NewRepoLearn(db *sqlx.DB, log *logrus.Logger) *repoLearn {
	return &repoLearn{db: db, collection: collectionRepoLearn, log: log}
}

func (rl *repoLearn) InsertWordLearn(ctx context.Context, word *models.Word) error {
	_, err := rl.db.Exec("INSERT INTO words_learns (id, english, russian, theme) VALUES ($1, $2, $3, $4)",
		word.Id, word.English, word.Russian, word.Theme)
	if err != nil {
		appErr := apperrors.InsertWordLearnErr.AppendMessage(err, word)
		rl.log.Error(appErr)
		//return appErr
	}

	return nil
}

func (rl *repoLearn) InsertWordsLearn(words []*models.Word) error {
	for _, v := range words {
		err := rl.InsertWordLearn(context.TODO(), v)
		if err != nil {
			return err
		}
	}

	return nil
}

func (rl *repoLearn) GetWordsLearn(quantity int) ([]*models.Word, error) {
	words := []*models.Word{}
	err := rl.db.Select(&words, "SELECT * FROM words_learns limit $1", quantity)
	if err != nil {
		appErr := apperrors.GetWordsLearnErr.AppendMessage(err)
		rl.log.Error(appErr)
		return nil, appErr
	}

	return words, nil
}

func (rl *repoLearn) DeleteLearnWordsId(id int) error {
	_, err := rl.db.Exec("delete from words_learns where id = $1", id)
	if err != nil {
		appErr := apperrors.DeleteLearnWordsIdErr.AppendMessage(err)
		rl.log.Error(appErr)
		return appErr
	}

	return nil
}
*/
