package service

import (
	"fmt"
	"postgresTakeWords/internal/models"
	"postgresTakeWords/internal/repositories"
	"time"

	"github.com/sirupsen/logrus"
)

const exit = "exit"

type ServiceCompetition struct {
	stat        *repositories.Stat
	repoLearn   repositories.RepoLearn
	repoWordsPg repositories.RepoWordsPg
	log         *logrus.Logger
}

func NewServiceCompetition(stat *repositories.Stat, repoLearn repositories.RepoLearn, repoWordsPg repositories.RepoWordsPg,
	log *logrus.Logger) *ServiceCompetition {
	return &ServiceCompetition{
		stat:        stat,
		repoLearn:   repoLearn,
		repoWordsPg: repoWordsPg,
		log:         log,
	}
}

func (c *ServiceCompetition) WorkTest(num int) error {
	s, err := c.repoWordsPg.GetWordsWhereRA(num)
	if err != nil {
		c.log.Error(err)
		return err
	}

	startTime := time.Now()
	LearnSlice := []*models.Word{}
	quantity := len(s)
	var capSlovar int = quantity
	NewSlovar := make([]*models.Word, capSlovar)
	copy(NewSlovar, (s)[:quantity])
	s = (s)[quantity:]
	fmt.Println("                     START")
	var yes int
	var not int
	var exit1 bool
	fmt.Println("range NewSlovar")
	for _, word := range NewSlovar {
		if exit1 {
			not++
			s = models.Preppend(s, word)
			continue
		}

		maps, err := c.repoWordsPg.GetWordsMap(word.Russian)
		if err != nil {
			c.log.Error(err)
			return err
		}

		y, n, exit := Compare(word, maps)
		if exit {
			exit1 = true
			n = 1
		}

		if y > 0 && n > 0 {
			yes++
			s = models.Preppend(s, word)
			LearnSlice = models.Preppend(LearnSlice, word)
			continue
		}

		if y > 0 {
			yes++
			word.RightAnswer += 1
			c.repoWordsPg.UpdateRightAnswer(word)
			s = models.AppendWord(s, word)
		} else if n > 0 {
			not++
			s = models.Preppend(s, word)
			LearnSlice = models.Preppend(LearnSlice, word)
		} else {
			break
		}
	}

	duration := time.Since(startTime)
	PrintTime(duration)
	sStat := models.NewStatistick(quantity, yes, not)
	c.stat.WriteStatistic(*sStat)
	fmt.Println(yes, not)

	wordsLearn := []*models.WordsLearn{}
	for _, word := range LearnSlice {
		wordLearn := &models.WordsLearn{
			ID:           word.ID,
			English:      word.English,
			Russian:      word.Russian,
			Preposition:  word.Preposition,
			Theme:        word.Theme,
			PartOfSpeech: word.PartOfSpeech,
			RightAnswer:  word.RightAnswer,
		}

		wordsLearn = append(wordsLearn, wordLearn)
	}

	err = c.repoLearn.InsertWordsLearn(wordsLearn)
	if err != nil {
		c.log.Error(err)
		return nil
	}

	return nil
}

func (c *ServiceCompetition) LearnWords(quantity int) error {
	wordsLearn, err := c.repoLearn.GetWordsLearn(quantity)
	if err != nil {
		c.log.Error(err)
		return err
	}

	s := []*models.Word{}
	for _, word := range wordsLearn {
		wordLearn := &models.Word{
			ID:           word.ID,
			English:      word.English,
			Russian:      word.Russian,
			Preposition:  word.Preposition,
			Theme:        word.Theme,
			PartOfSpeech: word.PartOfSpeech,
			RightAnswer:  word.RightAnswer,
		}

		s = append(s, wordLearn)
	}

	if ok := c.LearnWordsTest(s); !ok {
		c.log.Info("!ok)")
	}

	fmt.Println("After learn :", len(wordsLearn))
	for _, v := range wordsLearn {
		err := c.repoLearn.DeleteLearnWordsId(v.ID)
		if err != nil {
			c.log.Error(err)
			return err
		}
	}

	return nil
}

func (c *ServiceCompetition) LearnWordsTest(s []*models.Word) bool {
	startTime := time.Now()
	fmt.Println("                 Learn Words")
	for {
		if len(s) == 0 {
			break
		}

		maps, err := c.repoWordsPg.GetWordsMap(s[0].Russian)
		if err != nil {
			c.log.Error(err)
		}

		v := s[0]
		y, _, exit := Compare(v, maps)
		if exit {
			break
		}

		if y > 0 && len(s) != 1 {
			s = s[1:]
		} else if y < 1 {
			copy(s, s[1:])
			s[len(s)-1] = v
			PrintXpen(v.English)
		} else if y > 0 && len(s) == 1 {
			break
		}
	}

	duration := time.Since(startTime)
	PrintTime(duration)
	return true
}
