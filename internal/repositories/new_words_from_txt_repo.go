package repositories

import (
	"errors"
	"fmt"
	"os"
	"postgresTakeWords/internal/models"
	"strings"
	"unicode"

	"github.com/sirupsen/logrus"
)

type UpdateWordsFromTXTRepo struct {
	newWordsPath string
	log          *logrus.Logger
}

func NewUpdateWordsFromTXTRepo(newWordsPath string, log *logrus.Logger) *UpdateWordsFromTXTRepo {
	return &UpdateWordsFromTXTRepo{newWordsPath: newWordsPath, log: log}
}

func (tr *UpdateWordsFromTXTRepo) GetAllFromTXT() (*[]models.Word, error) {
	data, err := os.ReadFile(tr.newWordsPath)
	if err != nil {
		return nil, err
	}

	content := string(data)
	if content == "You need to add your words here" {
		err := errors.New("you need to add your words here")
		fmt.Println("You need to add new words in save/newWords.txt")
		return nil, err
	}

	lines := strings.Split(content, "\n")
	var words []models.Word
	for _, line := range lines {
		if line == "" {
			continue
		}

		lines := strings.Split(line, "-")
		if len(lines) <= 0 {
			continue
		}

		for i, v := range lines {
			lines[i] = strings.TrimSpace(v)
			lines[i] = strings.ReplaceAll(lines[i], ".", "")
			lines[i] = capitalizeFirstRune(lines[i])
		}

		id := 0
		theme := ""
		if len(lines) > 3 {
			theme = lines[2]
		}

		word := models.Word{Id: id, English: lines[0], Russian: lines[1], Theme: theme}
		words = append(words, word)
	}

	return &words, nil
}

func (tr *UpdateWordsFromTXTRepo) CleanNewWords(txt string) error {
	file, err := os.Create(tr.newWordsPath)
	if err != nil {
		tr.log.Error("Unable to create file:", err)
		os.Exit(1)
		return err
	}

	defer file.Close()
	file.WriteString(txt)
	return nil
}

func capitalizeFirstRune(line string) string {
	runes := []rune(line)
	for i, r := range runes {
		if i == 0 {
			runes[i] = unicode.ToUpper(r)
		}
	}

	return string(runes)
}
