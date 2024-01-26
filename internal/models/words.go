package models

import (
	"strings"

	"gorm.io/gorm"
)

type Word struct {
	gorm.Model
	ID           int    `json:"id" gorm:"primaryKey"`
	English      string `json:"english"`
	Russian      string `json:"russian"`
	Preposition  string `json:"preposition"`
	Theme        string `json:"theme"`
	PartOfSpeech string `json:"part_of_speech"`
	RightAnswer  int    `json:"rightAnswer" db:"right_answer"`
	CreatedAt    string `json:"created_at"`
}

type WordsLearn struct {
	gorm.Model
	ID           int    `json:"id" gorm:"primaryKey"`
	English      string `json:"english"` //gorm:"primaryKey"`
	Russian      string `json:"russian"`
	Preposition  string `json:"preposition"`
	Theme        string `json:"theme"`
	PartOfSpeech string `json:"part_of_speech"`
	RightAnswer  int    `json:"rightAnswer" db:"right_answer"`
	CreatedAt    string `json:"created_at"`
}

func NewWord() *Word {
	return &Word{}
}

func NewLibrary(newId int, newEnglish string, newRussian string, newTheme string) *Word {
	return &Word{ID: newId, English: newEnglish, Russian: newRussian, Theme: newTheme}
}

type Slovarick []Word

func NewSlovar(w Word) *Slovarick {
	s := &Slovarick{w}
	return s
}

func AppendWord(s []*Word, w *Word) []*Word {
	s = append(s, w)
	return s
}

func Preppend(s []*Word, w *Word) []*Word {
	SliceWord := []*Word{w}
	s = append(SliceWord, s...)
	return s
}

func (s *Slovarick) CreateAndInitMapWords() *map[string][]string {
	maps := make(map[string][]string)
	for _, w := range *s {
		maps[w.Russian] = append(maps[w.Russian], strings.ToLower(w.English))
	}

	return &maps
}
