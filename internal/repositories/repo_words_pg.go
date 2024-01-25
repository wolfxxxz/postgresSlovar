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
	GetWordsMap() (*map[string][]string, error)
	InsertWordsLibrary(ctx context.Context, library []*models.Word) error
}

type repoWords struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewRepoWordsGorm(db *gorm.DB, log *logrus.Logger) RepoWordsPg { //Words {
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

func (rt *repoWords) InsertWord(ctx context.Context, word *models.Word) error {
	if word == nil {
		appErr := apperrors.InsertWordsLibraryErr.AppendMessage("lib == nil")
		rt.log.Error(appErr)
		return appErr
	}

	tx := rt.db.WithContext(ctx)
	if tx.Error != nil {
		appErr := apperrors.InsertWordsLibraryErr.AppendMessage(tx.Error)
		rt.log.Error(appErr)
		return appErr
	}

	result := tx.Create(word)
	if result.Error != nil {
		appErr := apperrors.InsertWordsLibraryErr.AppendMessage(result.Error)
		rt.log.Error(appErr)
		return appErr
	}

	if result.RowsAffected == 0 {
		appErr := apperrors.InsertWordsLibraryErr.AppendMessage("no rows affected")
		rt.log.Error(appErr)
		return appErr
	}

	createdLib := &models.Word{}
	if err := tx.First(createdLib, "id = ?", word.ID).Error; err != nil {
		appErr := apperrors.InsertWordsLibraryErr.AppendMessage(err)
		rt.log.Error(appErr)
		return appErr
	}

	return nil
}

func (rt *repoWords) GetTranslationRus(word string) ([]*models.Word, error) {
	var words []*models.Word
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

func (rt *repoWords) InsertWordsLibrary(ctx context.Context, library []*models.Word) error {
	if library == nil {
		appErr := apperrors.InsertWordsLibraryErr.AppendMessage("library == nil")
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
		return fmt.Errorf("no rows affected when updating word with ID %d", word.ID)
	}

	return nil
}

func (rt *repoWords) GetWordsMap() (*map[string][]string, error) {
	var words []*models.Word
	err := rt.db.Order("russian").Find(&words).Error
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
