package words

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

func StartCompetition(db *sqlx.DB) error {
	for {
		fmt.Println("  Update Library_by_txt_new_words: `update`")
		time.Sleep(20 * time.Millisecond)
		fmt.Println("    Test knowlige: `test`")
		time.Sleep(20 * time.Millisecond)
		fmt.Println("      Learn Words: `learn`")
		time.Sleep(20 * time.Millisecond)
		fmt.Println("      Update by json: `update_json`")
		time.Sleep(20 * time.Millisecond)
		fmt.Println("      Download all words: `download`")
		time.Sleep(20 * time.Millisecond)
		fmt.Println("      Words map: `map`")
		time.Sleep(20 * time.Millisecond)
		fmt.Println("      Exit: `exit`")
		var command string
		fmt.Scan(&command)
		switch command {
		case "update":
			var LibraryOldWords []Word
			DecodeJsonSliceWord(&LibraryOldWords, "save/library.json")
			fmt.Println(len(LibraryOldWords))
			var LibraryWords []Word
			DecodeTXT(&LibraryWords, "save/newWords.txt")
			SaveEmptyTXT("save/newWords.txt", "You need to add your words here")
			fmt.Println(LibraryWords)
			err := InsertWords(db, &LibraryWords)
			if err != nil {
				fmt.Println("main ", err)
			}

			UpdateLibrary(&LibraryWords, &LibraryOldWords)
			EncodeJson(&LibraryOldWords, "save/library.json")
		case "test":
			var quantity int
			fmt.Println("Количество слов для теста")
			fmt.Scan(&quantity)
			var LibraryWords []Word
			GetWordsWhereRA(db, &LibraryWords, quantity)
			LibraryLearn := WorkTest(&LibraryWords)
			err := InsertWordsLearn(db, LibraryLearn)
			if err != nil {
				log.Println("You need to learn words because there are lots of words accumulated in library")
			}

		case "learn":
			var quantity int
			fmt.Println("Количество слов to learn")
			fmt.Scan(&quantity)
			var Learn []Word
			GetWordsLearn(db, &Learn, quantity)
			LearnWords(Learn)
			fmt.Println("After learn :", len(Learn))
			for _, v := range Learn {
				err := DeleteLearnWordsId(db, v.Id)
				if err != nil {
					fmt.Println(err)
				}
			}

		case "upload_json":
			var LibraryOldWords []Word
			DecodeJsonSliceWord(&LibraryOldWords, "save/library.json")
			fmt.Println(len(LibraryOldWords))
			err := InsertWords(db, &LibraryOldWords)
			if err != nil {
				fmt.Println("main ", err)
			}

			EncodeJson(&LibraryOldWords, "save/library.json")
		case "update_json":
			var LibraryOldWords []Word
			DecodeJsonSliceWord(&LibraryOldWords, "save/library.json")
			fmt.Println(len(LibraryOldWords))
			for _, v := range LibraryOldWords {
				err := UpdateWord(db, &v)
				if err != nil {
					fmt.Println(err)
				}
			}

		case "download":
			var LibraryOldWords []Word
			err := GetAllWords(db, &LibraryOldWords)
			if err != nil {
				fmt.Println(err)
			}

			EncodeJson(&LibraryOldWords, "save/library.json")
		case "map":
			maps, err := GetWordsMap(db)
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
