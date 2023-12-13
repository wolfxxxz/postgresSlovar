package models

import (
	"fmt"
	"strings"
)

type Word struct {
	Id          int    `json:"id"`
	English     string `json:"english"`
	Russian     string `json:"russian"`
	Theme       string `json:"theme"`
	RightAnswer int    `json:"rightAnswer" db:"right_answer"`
}

func NewWord() *Word {
	return &Word{}
}

func NewLibrary(newId int, newEnglish string, newRussian string, newTheme string) *Word {
	return &Word{Id: newId, English: newEnglish, Russian: newRussian, Theme: newTheme}
}

type Slovarick []Word

func NewSlovar(w Word) *Slovarick {
	s := &Slovarick{w}
	return s
}

func AppendWord(s *[]Word, w Word) {
	*s = append(*s, w)
}

func Preppend(s *[]Word, w Word) {
	SliceWord := []Word{w}
	*s = append(SliceWord, *s...)
}

func (s *Slovarick) CreateAndInitMapWords() *map[string][]string {
	maps := make(map[string][]string)
	for _, w := range *s {
		maps[w.Russian] = append(maps[w.Russian], strings.ToLower(w.English))
	}

	return &maps
}

func (oldWords Slovarick) UpdateLibraryOnlyNewWords(NewWords Slovarick) {
	c := len(oldWords)
	oldWords = append(NewWords, oldWords...)
	d := len(oldWords)
	if d != c {
		fmt.Println("                   New Words Add:", d-c)
	} else {
		fmt.Println("Для загрузки слов списком необходимо упорядочить и вставить слова в файл `save/newWords.txt`")
		fmt.Println("english - перевод - тема")
		fmt.Println("в конце оставить пустую строчку")
		fmt.Println("I believe in you!!!")
	}
}
