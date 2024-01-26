package service

import (
	"postgresTakeWords/internal/models"
	"postgresTakeWords/internal/repositories"

	"github.com/sirupsen/logrus"
)

type ServiceTranslator struct {
	repoWordsPg repositories.RepoWordsPg
	log         *logrus.Logger
}

func NewServiceTranslator(repoWordsPg repositories.RepoWordsPg,
	log *logrus.Logger) *ServiceTranslator {
	return &ServiceTranslator{
		repoWordsPg: repoWordsPg,
		log:         log,
	}
}

func (c *ServiceTranslator) GetTranslation(word string) ([]*models.Word, error) {
	capitalizedWord := capitalizeFirstRune(word)
	if isCyrillic(capitalizedWord) {
		words, err := c.repoWordsPg.GetTranslationRus(capitalizedWord)
		if err != nil {
			return nil, err
		}

		if len(words) == 0 {
			words, err = c.repoWordsPg.GetTranslationRusLike(capitalizedWord)
			if err != nil {
				return nil, err
			}

		}

		return words, nil
	}

	if !isCyrillic(capitalizedWord) {
		words, err := c.repoWordsPg.GetTranslationEngl(capitalizedWord)
		if err != nil {
			return nil, err
		}

		if len(words) == 0 {
			words, err = c.repoWordsPg.GetTranslationEnglLike(capitalizedWord)
			if err != nil {
				return nil, err
			}
		}

		return words, nil
	}

	return nil, nil
}
