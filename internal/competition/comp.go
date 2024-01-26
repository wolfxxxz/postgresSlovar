package competition

import (
	"fmt"
	"postgresTakeWords/internal/models"
	"postgresTakeWords/internal/repositories"
	"postgresTakeWords/internal/service"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Competition struct {
	stat            *repositories.Stat
	repoLearn       repositories.RepoLearn
	repoWordsPg     repositories.RepoWordsPg
	repoBackUpCopy  *repositories.BackUpCopyRepo
	repoUpdateByTXT *repositories.UpdateWordsFromTXTRepo
	log             *logrus.Logger
}

func NewCompetition(statPath string, reserveCopyPath string, reserveCopyPathTXT string, newWordsPath string, sqlDB *gorm.DB, log *logrus.Logger) *Competition {
	return &Competition{
		stat:            repositories.NewStatRepo(statPath, log),
		repoLearn:       repositories.NewLearnGormRepo(sqlDB, log),
		repoWordsPg:     repositories.NewRepoWordsGorm(sqlDB, log),
		repoBackUpCopy:  repositories.NewBackUpCopyRepo(reserveCopyPath, reserveCopyPathTXT, log),
		repoUpdateByTXT: repositories.NewUpdateWordsFromTXTRepo(newWordsPath, log),
		log:             log,
	}
}

func (c *Competition) test() error {
	var quantity int
	fmt.Println("Количество слов для теста")
	fmt.Scan(&quantity)
	servWord := service.NewServiceCompetition(c.stat, c.repoLearn, c.repoWordsPg, c.log)
	return servWord.WorkTest(quantity)
}

func (c *Competition) translator() error {
	fmt.Printf("для выхода введите [%v]\n", exit)
	Word := ""
	for {
		fmt.Println()
		fmt.Scan(&Word)
		if Word == exit {
			break
		}

		servTranslate := service.NewServiceTranslator(c.repoWordsPg, c.log)
		words, err := servTranslate.GetTranslation(Word)
		if err != nil {
			return err
		}

		if len(words) == 0 {
			fmt.Println("there is not such a word in this library")
		}

		printAll(words)
	}

	return nil
}

func (c *Competition) backup() error {
	c.log.Info("download All words from db")

	servLocal := service.NewServiceLocal(c.stat, c.repoLearn, c.repoWordsPg, c.repoBackUpCopy, c.repoUpdateByTXT, c.log)
	err := servLocal.Backup()
	if err != nil {
		return err
	}

	c.log.Info("backup has been saved")
	return nil
}

func (c *Competition) updateFromBackUp() error {
	servLocal := service.NewServiceLocal(c.stat, c.repoLearn, c.repoWordsPg, c.repoBackUpCopy, c.repoUpdateByTXT, c.log)
	err := servLocal.UpdateFromBackUp()
	if err != nil {
		return err
	}

	return nil
}

func (c *Competition) restore() error {
	servWord := service.NewServiceLocal(c.stat, c.repoLearn, c.repoWordsPg, c.repoBackUpCopy, c.repoUpdateByTXT, c.log)
	err := servWord.Restore()
	if err != nil {
		return err
	}

	return nil
}

func (c *Competition) learn() error {
	var quantity int
	fmt.Println("Количество слов to learn")
	fmt.Scan(&quantity)

	servWord := service.NewServiceCompetition(c.stat, c.repoLearn, c.repoWordsPg, c.log)
	err := servWord.LearnWords(quantity)
	if err != nil {
		c.log.Info("!ok)")
	}

	return nil
}

func (c *Competition) update() error {
	servLocal := service.NewServiceLocal(c.stat, c.repoLearn, c.repoWordsPg, c.repoBackUpCopy, c.repoUpdateByTXT, c.log)
	err := servLocal.Update()
	if err != nil {
		return err
	}

	return nil
}

func printAll(words []*models.Word) {
	for _, word := range words {
		fmt.Printf(" %v -- %v \n", word.Russian, word.English)
	}
}
