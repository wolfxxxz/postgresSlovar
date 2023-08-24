package words

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

func GetAllWords(db *sqlx.DB, words *[]Word) error {
	//var words []Word
	err := db.Select(words, "SELECT * FROM words order by theme")

	if err != nil {
		log.Fatal(err)
	}

	// Выводим результаты
	return nil
}

/*
func InsertWord(db *sqlx.DB, word *Word) error {
	_, err := db.Exec("INSERT INTO words ( english, russian, theme, right_answer) VALUES ($1, $2, $3, $4) RETURNING id",
		word.Id, word.English, word.Russian, word.Theme, word.RightAnswer)
	if err != nil {
		return err
	}
	return nil
}*/

// Проверка по Id
func GetWordById(db *sqlx.DB, word *Word) (*Word, error) {
	var getWord Word
	err := db.Get(&getWord, "SELECT * FROM words WHERE id=$1", word.Id)
	if err != nil {
		return nil, err
	}
	return &getWord, nil
}

/*
// Проверить существование
func CheckWordByEnglish(db *sqlx.DB, word *Word) (error, int) {
	err := db.Get(word, "SELECT * FROM words WHERE english=$1", word.English)
	if err != nil {
		return nil, 0
	}
	return fmt.Errorf("Word [%v] is already exist in DB", word.English), id
}*/
// Проверить существование
func CheckWordByEnglish(db *sqlx.DB, word *Word) (int, error) {
	var id int
	query := "SELECT id FROM words WHERE english=$1"
	err := db.QueryRow(query, word.English).Scan(&id)
	if err != nil {
		return 0, nil
	}
	return id, fmt.Errorf("Word [%v] is already exist in DB", word.English)
}

// +
func InsertWord(db *sqlx.DB, word *Word) error {
	var insertedId int // Переменная для хранения сгенерированного Id
	query := "INSERT INTO words (english, russian, theme, right_answer) VALUES ($1, $2, $3, $4) RETURNING id"
	err := db.QueryRow(query, word.English, word.Russian, word.Theme, word.RightAnswer).Scan(&insertedId)
	if err != nil {
		return err
	}
	fmt.Println(insertedId)

	word.Id = insertedId // Присвоить сгенерированный Id обратно в структуру
	return nil
}

// Отправить слова в db
func InsertWords(db *sqlx.DB, words *[]Word) error {

	for i, v := range *words {
		id, err := CheckWordByEnglish(db, &v)
		if err != nil {
			log.Println(err)
			vCopy := v
			vCopy.Id = id
			err = InsertWordLearn(db, &vCopy)
			if err != nil {
				log.Println(err)
			}
		} else {
			err = InsertWord(db, &(*words)[i])
			if err != nil {
				log.Println(err)
			}
		}
	}
	return nil
}

func GetWordsWhereRA(db *sqlx.DB, words *[]Word, quantity int) error {
	err := db.Select(words, "SELECT * FROM words order by right_answer limit $1", quantity)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func InsertWordLearn(db *sqlx.DB, word *Word) error {
	//query := "INSERT INTO words (id, english, russian, theme) VALUES ($1, $2, $3, $4) RETURNING id"
	_, err := db.Exec("INSERT INTO words_learn (id, english, russian, theme) VALUES ($1, $2, $3, $4)",
		word.Id, word.English, word.Russian, word.Theme)
	if err != nil {
		return err
	}

	return nil
}

// Отправить слова в db
func InsertWordsLearn(db *sqlx.DB, words *[]Word) error {
	var errr error
	for _, v := range *words {

		err := InsertWordLearn(db, &v)
		if err != nil {
			//log.Println(err)
			errr = fmt.Errorf("%v, %v", errr, err)
		}
	}
	if errr != nil {
		return errr
	}
	return nil
}

func GetWordsLearn(db *sqlx.DB, words *[]Word, quantity int) error {

	err := db.Select(words, "SELECT * FROM words_learn limit $1", quantity)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func DeleteLearnWordsId(db *sqlx.DB, id int) error {
	_, err := db.Exec("delete from words_learn where id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateRightAnswer(db *sqlx.DB, word *Word) error {
	res, err := db.Exec("update words set right_answer=$1 where id=$2", word.RightAnswer, word.Id)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func UpdateWord(db *sqlx.DB, word *Word) error {
	res, err := db.Exec("update words set english=$1, russian=$2, theme=$3 where id=$4",
		word.English, word.Russian, word.Theme, word.Id)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func GetWordsMap(db *sqlx.DB) (*map[string][]string, error) {
	var words Slovarick
	err := db.Select(&words, "SELECT english, russian FROM words order by russian")
	if err != nil {
		log.Fatal("something wrong ", err)
	}

	wordsLib := words.CreateAndInitMapWords()
	return wordsLib, nil
}
