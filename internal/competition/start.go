package competition

import (
	"fmt"
	"log"
	"postgresTakeWords/internal/models"
	"time"
)

func (c *Competition) StartCompetition() error {
	for {
		fmt.Println("  Update Library_by_txt_new_words: `update`")
		time.Sleep(20 * time.Millisecond)
		fmt.Println("    Test knowlige: `test`")
		time.Sleep(20 * time.Millisecond)
		fmt.Println("      Learn Words: `learn`")
		time.Sleep(20 * time.Millisecond)
		fmt.Println("      Update by json: `update_json`")
		time.Sleep(20 * time.Millisecond)
		fmt.Println("      Download all words: `downloadFromDB`")
		time.Sleep(20 * time.Millisecond)
		fmt.Println("      Words map: `map`")
		time.Sleep(20 * time.Millisecond)
		fmt.Println("      Exit: `exit`")
		var command string
		fmt.Scan(&command)
		switch command {
		case "update":
			newWords, err := c.repoUpdateByTXT.GetAllFromTXT()
			if err != nil {
				c.log.Error(err)
				return err
			}

			c.log.Info(newWords)
			c.repoUpdateByTXT.SaveEmptyTXT("You need to add your words here")

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

			err = c.repoBackUpCopy.SaveAll(words)
			if err != nil {
				c.log.Error(err)
				return err
			}
		case "test":
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

		case "learn":
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

		case "upload_json":
			oldWords, err := c.repoBackUpCopy.DecodeJsonSliceWord()
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
		case "update_json":
			wordsFromBackUp, err := c.repoBackUpCopy.DecodeJsonSliceWord()
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

		case "downloadFromDB":
			c.log.Info("download All words from db")
			oldWords, err := c.repoWordsPg.GetAllWords()
			if err != nil {
				c.log.Error(err)
				return err
			}

			c.log.Infof("Get All From DB len [%v]", len(*oldWords))

			err = c.repoBackUpCopy.SaveAll(oldWords)
			if err != nil {
				c.log.Error(err)
				return err
			}
		case "map":
			maps, err := c.repoWordsPg.GetWordsMap()
			if err != nil {
				c.log.Error(err)
				return err
			}

			c.log.Info((*maps)["Большой"])
		case "exit":
			fmt.Println("You have to do it, your dream wait")
			return nil
		}
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
