package repositories

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"postgresTakeWords/internal/models"
	"strings"

	"github.com/sirupsen/logrus"
)

type UpdateWordsFromTXTRepo struct {
	newWordsPath string
	log          *logrus.Logger
}

func NewUpdateWordsFromTXTRepo(newWordsPath string, log *logrus.Logger) *UpdateWordsFromTXTRepo {
	return &UpdateWordsFromTXTRepo{newWordsPath: newWordsPath, log: log}
}

func (tr *UpdateWordsFromTXTRepo) DownloadFromJsonAndDecodeModel() *[]models.Word {
	filejson, err := os.Open(tr.newWordsPath)
	if err != nil {
		log.Fatal(err)
	}

	defer filejson.Close()
	data, err := io.ReadAll(filejson)
	if err != nil {
		log.Fatal(err)
	}

	var words []models.Word
	json.Unmarshal(data, &words)
	return &words
}

func (tr *UpdateWordsFromTXTRepo) EncodeJson(s *[]models.Word) {
	byteArr, err := json.MarshalIndent(s, "", "   ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(tr.newWordsPath, byteArr, 0666) //-rw-rw-rw- 0664
	if err != nil {
		log.Fatal(err)
	}
}

func (tr *UpdateWordsFromTXTRepo) GetAllFromTXT() (*[]models.Word, error) {
	data, err := os.ReadFile(tr.newWordsPath)
	if err != nil {
		return nil, err
	}

	content := string(data)
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

		word := models.NewLibrary(id, lines[0], lines[1], theme)
		words = append(words, *word)
	}

	return &words, nil
}

func (tr *UpdateWordsFromTXTRepo) EncodeTXT(s models.Slovarick) {
	file, err := os.Create(tr.newWordsPath)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}

	defer file.Close()
	for _, v := range s {
		file.WriteString(v.English)
		file.WriteString(" - ")
		file.WriteString(v.Russian)
		if v.Theme != "" {
			file.WriteString(" - ")
			file.WriteString(v.Theme)
		}

		file.WriteString("\n")
	}
}

func (tr *UpdateWordsFromTXTRepo) SaveEmptyTXT(txt string) {
	file, err := os.Create(tr.newWordsPath)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}

	defer file.Close()
	file.WriteString(txt)
}

func (tr *UpdateWordsFromTXTRepo) SaveForLearningTxt(s models.Slovarick) {
	file, err := os.Create(tr.newWordsPath)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}

	defer file.Close()
	if len(s) != 0 {
		for _, v := range s {
			file.WriteString(v.English)
			var length = len(v.English)
			if length <= 30 {
				for i := 0; i+length <= 25; i++ {
					file.WriteString(" ")
				}
			}
			file.WriteString(" - ")
			file.WriteString(v.Russian)
			file.WriteString("\n")
		}
	} else {
		fmt.Println("empty learn words")
	}
}

/*
func ScanStringOne() (string, error) {
	fmt.Print("       ...")
	in := bufio.NewScanner(os.Stdin)
	if in.Scan() {
		return in.Text(), nil
	}

	if err := in.Err(); err != nil {
		return "", err
	}

	return "", nil
}

func Print(l models.Slovarick) {
	for _, v := range l {
		fmt.Print(v.English, " - ", v.Russian, " - ", v.Theme)
		fmt.Println()
	}
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
*/
