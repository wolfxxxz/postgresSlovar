package competition

import (
	"fmt"
	"log"
	"postgresTakeWords/internal/models"
	"time"
	"unicode"
)

func (c *Competition) StartCompetition() error {
	for {
		printMenu()
		var command string
		fmt.Scan(&command)
		switch command {
		case update:
			c.log.Info(update)
			if err := c.update(); err != nil {
				return err
			}

		case test:
			if err := c.test(); err != nil {
				return err
			}

		case learn:
			if err := c.learn(); err != nil {
				return err
			}

		case restore:
			if err := c.restore(); err != nil {
				return err
			}

		case updateFromBackUp:
			if err := c.updateFromBackUp(); err != nil {
				return err
			}

		case backup:
			if err := c.backup(); err != nil {
				return err
			}

		case mapWords:
			if err := c.translator(); err != nil {
				c.log.Error(err)
			}

		case exit:
			fmt.Println("You have to do it, your dream wait")
			return nil
		}
	}

}

func (c Competition) translator() error {
	fmt.Printf("для выхода введите [%v]\n", exit)
	Word := ""
	for {
		fmt.Println()
		fmt.Scan(&Word)
		if Word == exit {
			break
		}

		capitalizedWord := capitalizeFirstRune(Word)
		if isCyrillic(capitalizedWord) {
			words, err := c.repoWordsPg.GetTranslationRus(capitalizedWord)
			if err != nil {
				return err
			}

			if words == nil {
				words, err = c.repoWordsPg.GetTranslationRusLike(capitalizedWord)
				if err != nil {
					return err
				}

			}

			printAll(words)
			continue
		}

		if !isCyrillic(capitalizedWord) {
			words, err := c.repoWordsPg.GetTranslationEngl(capitalizedWord)
			if err != nil {
				return err
			}

			if words == nil {
				words, err = c.repoWordsPg.GetTranslationEnglLike(capitalizedWord)
				if err != nil {
					return err
				}

			}

			printAll(words)
			continue

		}
	}

	return nil
}

func printAll(words []*models.Word) {
	for _, word := range words {
		fmt.Printf(" %v -- %v \n", word.Russian, word.English)
	}
}

func printEngl(words []*models.Word) {
	for i, word := range words {
		fmt.Printf("num %v and value %v | ", i+1, word.English)
	}
}

func printRuss(words []*models.Word) {
	for i, word := range words {
		fmt.Printf("num %v and value %v | ", i+1, word.Russian)
	}
}

func isCyrillic(s string) bool {
	for _, r := range s {
		if unicode.Is(unicode.Cyrillic, r) {
			return true
		}
	}
	return false
}

// func (c Competition) translator() error {
// 	maps, err := c.repoWordsPg.GetWordsMap()
// 	if err != nil {
// 		c.log.Error(err)
// 		return err
// 	}

// 	fmt.Printf("для выхода введите [%v]\n", exit)
// 	rusWord := ""
// 	for {
// 		fmt.Scan(&rusWord)
// 		if rusWord == exit {
// 			break
// 		}

// 		capitalizedWord := capitalizeFirstRune(rusWord)

// 		fmt.Println((*maps)[capitalizedWord])
// 	}

// 	return nil
// }

func (c Competition) backup() error {
	c.log.Info("download All words from db")
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

	c.log.Info("backup has been saved")
	return nil
}

func (c Competition) updateFromBackUp() error {
	wordsFromBackUp, err := c.repoBackUpCopy.GetAllFromBackUp()
	if err != nil {
		c.log.Error(err)
		return err
	}

	for _, v := range *wordsFromBackUp {
		err := c.repoWordsPg.UpdateWord(&v)
		if err != nil {
			c.log.Error(err)
			return err
		}
	}

	return nil
}

func (c Competition) restore() error {
	oldWords, err := c.repoBackUpCopy.GetAllFromBackUp()
	if err != nil {
		c.log.Error(err)
		return err
	}

	fmt.Println(len(*oldWords))
	err = c.InsertWordsIfNotExist(oldWords)
	if err != nil {
		c.log.Errorf("main %v", err)
		return err
	}

	c.log.Info("All words have been inserted in DB")
	return nil
}

func (c Competition) learn() error {
	var quantity int
	fmt.Println("Количество слов to learn")
	fmt.Scan(&quantity)
	wordsLearn, err := c.repoLearn.GetWordsLearn(quantity)
	if err != nil {
		c.log.Error(err)
		return err
	}

	if ok := c.LearnWords(*wordsLearn); !ok {
		c.log.Info("!ok)")
	}

	fmt.Println("After learn :", len(*wordsLearn))
	for _, v := range *wordsLearn {
		err := c.repoLearn.DeleteLearnWordsId(v.Id)
		if err != nil {
			c.log.Error(err)
			return err
		}
	}

	return nil
}

func (c *Competition) test() error {
	var quantity int
	fmt.Println("Количество слов для теста")
	fmt.Scan(&quantity)
	testWords, err := c.repoWordsPg.GetWordsWhereRA(quantity)
	if err != nil {
		c.log.Error(err)
		return err
	}

	maps, err := c.repoWordsPg.GetWordsMap()
	if err != nil {
		c.log.Error(err)
		return err
	}

	wrongWords, err := c.WorkTest(testWords, maps)
	if err != nil {
		c.log.Error(err)
		return err
	}
	err = c.repoLearn.InsertWordsLearn(wrongWords)
	if err != nil {
		c.log.Error(err)
		return err
	}

	return nil
}

func (c *Competition) update() error {
	newWords, err := c.repoUpdateByTXT.GetAllFromTXT()
	if err != nil {
		c.log.Error(err)
		return err
	}

	c.log.Info(newWords)
	c.repoUpdateByTXT.CleanNewWords("You need to add your words here")

	err = c.InsertWordsIfNotExist(newWords)
	if err != nil {
		c.log.Errorf("main %v", err)
		return err
	}

	words, err := c.repoWordsPg.GetAllWords()
	if err != nil {
		c.log.Error(err)
		return err
	}

	err = c.repoBackUpCopy.SaveAllAsJson(words)
	if err != nil {
		c.log.Error(err)
		return err
	}

	return nil
}

func printMenu() {
	menu := []string{
		fmt.Sprintf("      Update library (newWords.txt): [%v]\n", update),
		fmt.Sprintf("      Test knowlige:   [%v]\n", test),
		fmt.Sprintf("      Learn words:     [%v]\n", learn),
		fmt.Sprintf("      Restore by json: [%v]\n", restore),
		fmt.Sprintf("      Update by json:  [%v]\n", updateFromBackUp),
		fmt.Sprintf("      Backup:          [%v]\n", backup),
		fmt.Sprintf("      Init map words:  [%v]\n", mapWords),
		fmt.Sprintf("          Exit:        [%v]\n", exit),
	}

	for _, pos := range menu {
		fmt.Println(pos)
		time.Sleep(20 * time.Millisecond)
	}
}

func (c *Competition) InsertWordsIfNotExist(words *[]models.Word) error {
	for _, word := range *words {
		id, err := c.repoWordsPg.CheckWordByEnglish(&word)
		if err != nil {
			c.log.Error(err)
			vCopy := word
			vCopy.Id = id
			err = c.repoLearn.InsertWordLearn(&vCopy)
			if err != nil {
				c.log.Error(err)
				return err
			}
		}

		if id == 0 {
			err = c.repoWordsPg.InsertWord(&word)
			if err != nil {
				log.Println(err)
				return err
			}
		}
	}

	return nil
}
