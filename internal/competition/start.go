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
			var LibraryOldWords []models.Word
			c.repoUpdateByTXT.DecodeJsonSliceWord(&LibraryOldWords)
			fmt.Println(len(LibraryOldWords))
			var LibraryWords []models.Word
			c.repoUpdateByTXT.DecodeTXT(&LibraryWords)
			c.repoUpdateByTXT.SaveEmptyTXT("You need to add your words here")
			c.log.Info(LibraryWords)
			err := c.InsertWords(&LibraryWords)
			if err != nil {
				c.log.Errorf("main %v", err)
			}

			UpdateLibrary(&LibraryWords, &LibraryOldWords)
			c.repoTXT.EncodeJson(&LibraryOldWords)
		case "test":
			var quantity int
			fmt.Println("Количество слов для теста")
			fmt.Scan(&quantity)
			var LibraryWords []models.Word
			c.repoTest.GetWordsWhereRA(&LibraryWords, quantity)
			LibraryLearn := c.WorkTest(&LibraryWords)
			err := c.repoLearn.InsertWordsLearn(LibraryLearn)
			if err != nil {
				log.Println("You need to learn words because there are lots of words accumulated in library")
			}

		case "learn":
			var quantity int
			fmt.Println("Количество слов to learn")
			fmt.Scan(&quantity)
			var Learn []models.Word
			c.repoLearn.GetWordsLearn(&Learn, quantity)
			c.LearnWords(Learn)
			fmt.Println("After learn :", len(Learn))
			for _, v := range Learn {
				err := c.repoLearn.DeleteLearnWordsId(v.Id)
				if err != nil {
					fmt.Println(err)
				}
			}

		case "upload_json":
			var LibraryOldWords []models.Word
			c.repoTXT.DecodeJsonSliceWord(&LibraryOldWords)
			fmt.Println(len(LibraryOldWords))
			err := c.InsertWords(&LibraryOldWords)
			if err != nil {
				fmt.Println("main ", err)
			}

			c.repoTXT.EncodeJson(&LibraryOldWords)
		case "update_json":
			var LibraryOldWords []models.Word
			c.repoTXT.DecodeJsonSliceWord(&LibraryOldWords)
			fmt.Println(len(LibraryOldWords))
			for _, v := range LibraryOldWords {
				err := c.repoTest.UpdateWord(&v)
				if err != nil {
					fmt.Println(err)
				}
			}

		case "downloadFromDB":
			var LibraryOldWords []models.Word
			err := c.repoTest.GetAllWords(&LibraryOldWords)
			if err != nil {
				fmt.Println(err)
			}

			c.repoTXT.EncodeJson(&LibraryOldWords)
		case "map":
			maps, err := c.repoTest.GetWordsMap()
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println((*maps)["Большой"])
		case "exit":
			fmt.Println("You have to do it, your dream wait")
			return nil
		}
	}

}

func (c *Competition) InsertWords(words *[]models.Word) error {
	for i, v := range *words {
		id, err := c.repoTest.CheckWordByEnglish(&v)
		if err != nil {
			log.Println(err)
			vCopy := v
			vCopy.Id = id
			err = c.repoLearn.InsertWordLearn(&vCopy)
			if err != nil {
				log.Println(err)
			}
		} else {
			err = c.repoTest.InsertWord(&(*words)[i])
			if err != nil {
				log.Println(err)
			}
		}
	}
	return nil
}
