package service

import (
	"context"
	"fmt"
	"log"
	"postgresTakeWords/internal/apperrors"
	"postgresTakeWords/internal/models"
	"postgresTakeWords/internal/repositories"

	"github.com/sirupsen/logrus"
)

type ServiceLocal struct {
	stat            *repositories.Stat
	repoLearn       repositories.RepoLearn
	repoWordsPg     repositories.RepoWordsPg
	repoBackUpCopy  *repositories.BackUpCopyRepo
	repoUpdateByTXT *repositories.UpdateWordsFromTXTRepo
	log             *logrus.Logger
}

func NewServiceLocal(stat *repositories.Stat, repoLearn repositories.RepoLearn, repoWordsPg repositories.RepoWordsPg,
	repoBackUpCopy *repositories.BackUpCopyRepo, repoUpdateByTXT *repositories.UpdateWordsFromTXTRepo,
	log *logrus.Logger) *ServiceLocal {
	return &ServiceLocal{
		stat:            stat,
		repoLearn:       repoLearn,
		repoWordsPg:     repoWordsPg,
		repoBackUpCopy:  repoBackUpCopy,
		repoUpdateByTXT: repoUpdateByTXT,
		log:             log,
	}
}

func (c *ServiceLocal) Backup() error {
	oldWords, err := c.repoWordsPg.GetAllWords()
	if err != nil {
		c.log.Error(err)
		return err
	}

	err = c.repoBackUpCopy.SaveAllAsJson(oldWords)
	if err != nil {
		c.log.Error(err)
		return err
	}

	err = c.repoBackUpCopy.SaveAllAsTXT(oldWords)
	if err != nil {
		c.log.Error(err)
		return err
	}

	err = c.repoBackUpCopy.SaveWordNewAsXLSX(oldWords)
	if err != nil {
		c.log.Error(err)
		return err
	}

	return nil
}

func (c *ServiceLocal) UpdateFromBackUp() error {
	wordsFromBackUp, err := c.repoBackUpCopy.GetAllWordsFromBackUpXlsx()
	if err != nil {
		c.log.Error(err)
		return err
	}

	for _, word := range wordsFromBackUp {
		err := c.repoWordsPg.UpdateWord(word)
		if err != nil {
			if err == &apperrors.UpdateWordRowAffectedErr {
				c.log.Info(fmt.Sprintf("insert word %v, ID %v", word.English, word.ID))
				err := c.repoWordsPg.InsertWord(context.TODO(), word)
				if err != nil {
					return err
				}

				continue
			}

			c.log.Error(err)
			return err
		}
	}

	return nil
}

func (c *ServiceLocal) Restore() error {
	oldWords, err := c.repoBackUpCopy.GetAllWordsFromBackUpXlsx()
	if err != nil {
		c.log.Error(err)
		return err
	}

	for _, word := range oldWords {
		err = c.repoWordsPg.InsertWord(context.TODO(), word)
		if err != nil {
			log.Println(err)
			return err
		}

	}

	if err != nil {
		c.log.Errorf("main %v", err)
		return err
	}

	c.log.Info("All words have been inserted in DB")
	return nil
}

func (c *ServiceLocal) Update() error {
	newWords, err := c.repoUpdateByTXT.GetAllFromTXT()
	if err != nil {
		c.log.Error(err)
		return err
	}

	c.log.Info("new words: ", len(newWords))

	countWords, err := c.repoWordsPg.GetAllWords()
	if err != nil {
		c.log.Error(err)
		return err
	}

	count := len(countWords)

	for i, v := range newWords {
		_ = v
		count++
		newWords[i].ID = count
	}

	for _, word := range newWords {
		id, err := c.repoWordsPg.CheckWordByEnglish(word)
		if err != nil {
			c.log.Error(err)
			return err
		}

		if id != 0 {
			wordLearn := &models.WordsLearn{
				ID:           word.ID,
				English:      word.English,
				Russian:      word.Russian,
				Preposition:  word.Preposition,
				Theme:        word.Theme,
				PartOfSpeech: word.PartOfSpeech,
				RightAnswer:  word.RightAnswer,
			}
			err = c.repoLearn.InsertWordLearn(context.TODO(), wordLearn)
			if err != nil {
				c.log.Error(err)
				return err
			}
		}

		if id == 0 {
			lenWords, err := c.repoWordsPg.GetLenWords(context.TODO())
			if err != nil {
				c.log.Error(err)
				return err
			}

			word.ID = lenWords + 1
			err = c.repoWordsPg.InsertWord(context.TODO(), word)
			if err != nil {
				c.log.Error(err)
				return err
			}
		}
	}

	err = c.repoUpdateByTXT.CleanNewWords("You need to type your words here")
	if err != nil {
		c.log.Error(err)
		return err
	}

	return nil
}
