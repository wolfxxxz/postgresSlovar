package repositories

import (
	"context"
	"fmt"
	"postgresTakeWords/internal/apperrors"
	"postgresTakeWords/internal/models"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RepoLearn interface {
	InsertWordLearn(ctx context.Context, word *models.WordsLearn) error
	InsertWordsLearn(words []*models.WordsLearn) error
	GetWordsLearn(quantity int) ([]*models.WordsLearn, error)
	DeleteLearnWordsId(id int) error
}

type repoLearnGorm struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewLearnGormRepo(db *gorm.DB, log *logrus.Logger) RepoLearn {
	return &repoLearnGorm{db: db, log: log}
}

func (rl *repoLearnGorm) InsertWordLearn(ctx context.Context, word *models.WordsLearn) error {
	if word == nil {
		appErr := apperrors.InsertWordsLearnErr.AppendMessage("lib == nil")
		rl.log.Error(appErr)
		return appErr
	}

	tx := rl.db.WithContext(ctx)
	if tx.Error != nil {
		appErr := apperrors.InsertWordsLearnErr.AppendMessage(tx.Error)
		rl.log.Error(appErr)
		return appErr
	}

	result := tx.Create(word)
	if result.Error != nil {
		appErr := apperrors.InsertWordsLearnErr.AppendMessage(result.Error)
		rl.log.Error(appErr)
		return nil
	}

	if result.RowsAffected == 0 {
		appErr := apperrors.InsertWordsLearnErr.AppendMessage("no rows affected")
		rl.log.Error(appErr)
		return appErr
	}

	createdLib := &models.WordsLearn{}
	if err := tx.First(createdLib, "id = ?", word.ID).Error; err != nil {
		appErr := apperrors.InsertWordsLearnErr.AppendMessage(err)
		rl.log.Error(appErr)
		return appErr
	}

	return nil
}

func (rl *repoLearnGorm) InsertWordsLearn(words []*models.WordsLearn) error {
	for _, word := range words {
		err := rl.InsertWordLearn(context.TODO(), word)
		if err != nil {
			return err
		}
	}

	return nil
}

func (rl *repoLearnGorm) GetWordsLearn(quantity int) ([]*models.WordsLearn, error) {
	var words []*models.WordsLearn
	err := rl.db.Limit(quantity).Find(&words).Error
	if err != nil {
		appErr := apperrors.GetWordsLearnErr.AppendMessage(err)
		rl.log.Error(appErr)
		return nil, appErr
	}

	return words, nil
}

func (rl *repoLearnGorm) DeleteLearnWordsId(id int) error {
	result := rl.db.Unscoped().Delete(&models.WordsLearn{}, id)
	if result.Error != nil {
		appErr := apperrors.DeleteLearnWordsIdErr.AppendMessage(result.Error)
		rl.log.Error(appErr)
		return appErr
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no records deleted for id %d", id)
	}

	return nil
}
