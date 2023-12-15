package repositories

import (
	"fmt"
	"postgresTakeWords/internal/apperrors"
	"postgresTakeWords/internal/models"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

var collectionRepoTest = "words"

type RepoWordsPg struct {
	db         *sqlx.DB
	collection string
	log        *logrus.Logger
}

func NewRepoWordsPg(db *sqlx.DB, log *logrus.Logger) *RepoWordsPg {
	return &RepoWordsPg{db: db, collection: collectionRepoTest, log: log}
}

func (rt *RepoWordsPg) GetAllWords() (*[]models.Word, error) {
	var words []models.Word
	err := rt.db.Select(&words, "SELECT * FROM words order by theme")
	if err != nil {
		appErr := apperrors.GetAllWordsErr.AppendMessage(err)
		rt.log.Error(appErr)
		return nil, appErr
	}

	return &words, nil
}

func (rt *RepoWordsPg) CheckWordByEnglish(word *models.Word) (int, error) {
	var id int
	query := "SELECT id FROM words WHERE english=$1"
	err := rt.db.QueryRow(query, word.English).Scan(&id)
	if err != nil {
		appErr := apperrors.CheckWordByEnglishErr.AppendMessage(err)
		rt.log.Error(appErr)
		return 0, nil
	}

	return id, fmt.Errorf("word [%v] is already exist in DB", word.English)
}

func (rt *RepoWordsPg) InsertWord(word *models.Word) error {
	var insertedId int
	query := "INSERT INTO words (english, russian, theme, right_answer) VALUES ($1, $2, $3, $4) RETURNING id"
	err := rt.db.QueryRow(query, word.English, word.Russian, word.Theme, word.RightAnswer).Scan(&insertedId)
	if err != nil {
		appErr := apperrors.InsertWordErr.AppendMessage(err)
		rt.log.Error(appErr)
		return appErr
	}

	rt.log.Info(insertedId)
	word.Id = insertedId
	return nil
}

func (rt *RepoWordsPg) GetWordsWhereRA(quantity int) (*[]models.Word, error) {
	var words []models.Word
	err := rt.db.Select(&words, "SELECT * FROM words order by right_answer limit $1", quantity)
	if err != nil {
		appErr := apperrors.GetWordsWhereRAErr.AppendMessage(err)
		rt.log.Error(appErr)
		return nil, appErr
	}

	return &words, nil
}

func (rt *RepoWordsPg) UpdateRightAnswer(word *models.Word) error {
	res, err := rt.db.Exec("update words set right_answer=$1 where id=$2", word.RightAnswer, word.Id)
	if err != nil {
		appErr := apperrors.UpdateRightAnswerErr.AppendMessage(err)
		rt.log.Error(appErr)
		return appErr
	}

	_, err = res.RowsAffected()
	if err != nil {
		appErr := apperrors.UpdateRightAnswerErr.AppendMessage(err)
		rt.log.Error(appErr)
		return appErr
	}

	return nil
}

func (rt *RepoWordsPg) UpdateWord(word *models.Word) error {
	res, err := rt.db.Exec("update words set english=$1, russian=$2, theme=$3 where id=$4",
		word.English, word.Russian, word.Theme, word.Id)
	if err != nil {
		appErr := apperrors.UpdateWordErr.AppendMessage(err)
		rt.log.Error(appErr)
		return appErr
	}

	_, err = res.RowsAffected()
	if err != nil {
		appErr := apperrors.UpdateWordErr.AppendMessage(err)
		rt.log.Error(appErr)
		return appErr
	}

	return nil
}

func (rt *RepoWordsPg) GetWordsMap() (*map[string][]string, error) {
	var words models.Slovarick
	err := rt.db.Select(&words, "SELECT english, russian FROM words order by russian")
	if err != nil {
		appErr := apperrors.GetWordsMapErr.AppendMessage(err)
		rt.log.Error(appErr)
	}

	wordsLib := words.CreateAndInitMapWords()
	return wordsLib, nil
}
