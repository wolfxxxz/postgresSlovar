package repositories

import (
	"context"
	"fmt"
	"postgresTakeWords/internal/apperrors"
	"postgresTakeWords/internal/models"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RepoWordsPg interface {
	GetAllWords() ([]*models.Word, error)
	InsertWord(ctx context.Context, word *models.Word) error
	GetTranslationRus(word string) ([]*models.Word, error)
	GetTranslationEngl(word string) ([]*models.Word, error)
	GetTranslationEnglLike(word string) ([]*models.Word, error)
	GetTranslationRusLike(word string) ([]*models.Word, error)
	CheckWordByEnglish(word *models.Word) (int, error)
	GetWordsWhereRA(quantity int) ([]*models.Word, error)
	UpdateRightAnswer(word *models.Word) error
	UpdateWord(word *models.Word) error
	GetWordsMap(russian string) (*map[string][]string, error)
	InsertWords(ctx context.Context, library []*models.Word) error
	GetLenWords(ctx context.Context) (int, error)
}

type repoWords struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewRepoWordsGorm(db *gorm.DB, log *logrus.Logger) RepoWordsPg {
	return &repoWords{db: db, log: log}
}

func (rt *repoWords) GetAllWords() ([]*models.Word, error) {
	var words []*models.Word
	err := rt.db.Order("right_answer").Find(&words).Error
	if err != nil {
		appErr := apperrors.GetAllWordsErr.AppendMessage(err)
		rt.log.Error(appErr)
		return nil, appErr
	}

	return words, nil
}

func (rt *repoWords) GetLenWords(ctx context.Context) (int, error) {
	var count int64
	err := rt.db.Model(&models.Word{}).Count(&count).Error
	if err != nil {
		appErr := apperrors.GetAllWordsErr.AppendMessage(err)
		rt.log.Error(appErr)
		return 0, appErr
	}

	return int(count), nil
}

func (rt *repoWords) InsertWord(ctx context.Context, word *models.Word) error {
	if word == nil {
		appErr := apperrors.InsertWordErr.AppendMessage("lib == nil")
		rt.log.Error(appErr)
		return appErr
	}

	tx := rt.db.WithContext(ctx)
	if tx.Error != nil {
		appErr := apperrors.InsertWordErr.AppendMessage(tx.Error)
		rt.log.Error(appErr)
		return appErr
	}

	result := tx.Create(word)
	if result.Error != nil {
		appErr := apperrors.InsertWordErr.AppendMessage(result.Error, word.English)
		rt.log.Error(appErr, word.ID)
		return appErr
	}

	if result.RowsAffected == 0 {
		appErr := apperrors.InsertWordErr.AppendMessage("no rows affected")
		rt.log.Error(appErr)
		return appErr
	}

	createdLib := &models.Word{}
	if err := tx.First(createdLib, "id = ?", word.ID).Error; err != nil {
		appErr := apperrors.InsertWordErr.AppendMessage(err)
		rt.log.Error(appErr)
		return appErr
	}

	return nil
}

func (rt *repoWords) GetTranslationRus(word string) ([]*models.Word, error) {
	rt.log.Info("GetTrRus")
	words := []*models.Word{}
	err := rt.db.Where("russian = ?", word).Find(&words).Error
	if err != nil {
		appErr := apperrors.GetTranslationRusErr.AppendMessage(err)
		rt.log.Error(appErr)
		return nil, appErr
	}

	return words, nil
}

func (rt *repoWords) GetTranslationRusLike(word string) ([]*models.Word, error) {
	var words []*models.Word
	err := rt.db.Where("russian LIKE ?", "%"+word+"%").Find(&words).Error
	if err != nil {
		appErr := apperrors.GetTranslationRusLikeErr.AppendMessage(err)
		rt.log.Error(appErr)
		return nil, appErr
	}

	return words, nil
}

func (rt *repoWords) GetTranslationEngl(word string) ([]*models.Word, error) {
	var words []*models.Word
	err := rt.db.Where("english = ?", word).Find(&words).Error
	if err != nil {
		appErr := apperrors.GetTranslationEnglErr.AppendMessage(err)
		rt.log.Error(appErr)
		return nil, appErr
	}

	return words, nil
}

func (rt *repoWords) GetTranslationEnglLike(word string) ([]*models.Word, error) {
	var words []*models.Word
	err := rt.db.Where("english LIKE ?", "%"+word+"%").Find(&words).Error
	if err != nil {
		appErr := apperrors.GetTranslationEnglLikeErr.AppendMessage(err)
		rt.log.Error(appErr)
		return nil, appErr
	}

	return words, nil
}

func (rt *repoWords) InsertWords(ctx context.Context, library []*models.Word) error {
	if library == nil {
		appErr := apperrors.InsertWordsErr.AppendMessage("library == nil")
		rt.log.Error(appErr)
		return appErr
	}

	for _, word := range library {
		rt.InsertWord(context.TODO(), word)
	}

	return nil
}

func (rt *repoWords) CheckWordByEnglish(word *models.Word) (int, error) {
	var id int
	result := rt.db.Model(&models.Word{}).Select("id").Where("english = ?", word.English).First(&id)
	if result.Error != nil {
		appErr := apperrors.CheckWordByEnglishErr.AppendMessage(result.Error)
		rt.log.Info(appErr)
		return 0, nil
	}

	return id, fmt.Errorf("word [%v] already exists in the database", word.English)
}

func (rt *repoWords) GetWordsWhereRA(quantity int) ([]*models.Word, error) {
	words := []*models.Word{}
	err := rt.db.Order("right_answer").Limit(quantity).Find(&words).Error
	if err != nil {
		appErr := apperrors.GetWordsWhereRAErr.AppendMessage(err)
		rt.log.Error(appErr)
		return nil, appErr
	}

	return words, nil
}

func (rt *repoWords) UpdateRightAnswer(word *models.Word) error {
	result := rt.db.Model(&models.Word{}).Where("id = ?", word.ID).Update("right_answer", word.RightAnswer)
	if result.Error != nil {
		appErr := apperrors.UpdateRightAnswerErr.AppendMessage(result.Error)
		rt.log.Error(appErr)
		return appErr
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows affected when updating right_answer for word with ID %d", word.ID)
	}

	return nil
}

func (rt *repoWords) UpdateWord(word *models.Word) error {
	result := rt.db.Model(&models.Word{}).Where("id = ?", word.ID).
		Updates(map[string]interface{}{
			"english":        word.English,
			"russian":        word.Russian,
			"theme":          word.Theme,
			"preposition":    word.Preposition,
			"part_of_speech": word.PartOfSpeech,
		})
	if result.Error != nil {
		appErr := apperrors.UpdateWordErr.AppendMessage(result.Error)
		rt.log.Error(appErr)
		return appErr
	}

	if result.RowsAffected == 0 {
		appErr := &apperrors.UpdateWordRowAffectedErr
		rt.log.Info(appErr)
		return appErr
	}

	return nil
}

func (rt *repoWords) GetWordsMap(russian string) (*map[string][]string, error) {
	var words []*models.Word
	err := rt.db.Order("russian").Where("russian = ?", russian).Find(&words).Error
	if err != nil {
		appErr := apperrors.GetWordsMapErr.AppendMessage(err)
		rt.log.Error(appErr)
		return nil, appErr
	}

	wordsMap := make(map[string][]string)
	for _, word := range words {
		wordsMap[word.Russian] = append(wordsMap[word.Russian], word.English)
	}

	return &wordsMap, nil
}
