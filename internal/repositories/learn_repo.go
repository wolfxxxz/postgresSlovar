package repositories

import (
	"fmt"
	"log"
	"postgresTakeWords/internal/models"

	"github.com/jmoiron/sqlx"
)

var collectionRepoLearn = "words_learn"

type RepoLearn struct {
	db         *sqlx.DB
	collection string
}

func NewRepoLearn(db *sqlx.DB) *RepoLearn {
	return &RepoLearn{db: db, collection: collectionRepoLearn}
}

func (rl *RepoLearn) InsertWordLearn(word *models.Word) error {
	_, err := rl.db.Exec("INSERT INTO words_learn (id, english, russian, theme) VALUES ($1, $2, $3, $4)",
		word.Id, word.English, word.Russian, word.Theme)
	if err != nil {
		return err
	}

	return nil
}

func (rl *RepoLearn) InsertWordsLearn(words *[]models.Word) error {
	var errr error
	for _, v := range *words {

		err := rl.InsertWordLearn(&v)
		if err != nil {
			errr = fmt.Errorf("%v, %v", errr, err)
		}
	}

	if errr != nil {
		return errr
	}

	return nil
}

func (rl *RepoLearn) GetWordsLearn(words *[]models.Word, quantity int) error {
	err := rl.db.Select(words, "SELECT * FROM words_learn limit $1", quantity)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (rl *RepoLearn) DeleteLearnWordsId(id int) error {
	_, err := rl.db.Exec("delete from words_learn where id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
