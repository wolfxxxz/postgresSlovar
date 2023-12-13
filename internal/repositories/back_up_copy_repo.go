package repositories

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"postgresTakeWords/internal/models"
	"strings"
	"unicode"

	"github.com/sirupsen/logrus"
)

type BackUpCopyRepo struct {
	reserveCopyPath    string
	reserveCopyPathTXT string
	log                *logrus.Logger
}

func NewBackUpCopyRepo(path string, pathTXT string, log *logrus.Logger) *BackUpCopyRepo {
	return &BackUpCopyRepo{reserveCopyPath: path, reserveCopyPathTXT: pathTXT, log: log}
}

func (tr *BackUpCopyRepo) DecodeJsonSliceWord() (*[]models.Word, error) {
	filejson, err := os.Open(tr.reserveCopyPath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer filejson.Close()
	data, err := io.ReadAll(filejson)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var words []models.Word
	err = json.Unmarshal(data, &words)
	if err != nil {
		return nil, err
	}

	return &words, nil
}

func (tr *BackUpCopyRepo) SaveAll(s *[]models.Word) error {
	byteArr, err := json.MarshalIndent(s, "", "   ")
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = os.WriteFile(tr.reserveCopyPath, byteArr, 0666) //-rw-rw-rw- 0664
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (tr *BackUpCopyRepo) DecodeTXT(s *[]models.Word) {
	data, err := os.ReadFile(tr.reserveCopyPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	content := string(data)
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		words := strings.Split(line, "-")
		if len(words) <= 0 {
			continue
		}

		for i, v := range words {
			words[i] = strings.TrimSpace(v)
			words[i] = strings.ReplaceAll(words[i], ".", "")
			words[i] = capitalizeFirstRune(words[i])
		}

		id := 0
		theme := ""
		if len(words) > 3 {
			theme = words[2]
		}

		word := models.NewLibrary(id, words[0], words[1], theme)
		*s = append(*s, *word)
	}
}

func (tr *BackUpCopyRepo) EncodeTXT(s models.Slovarick) {
	file, err := os.Create(tr.reserveCopyPathTXT)
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

func (tr *BackUpCopyRepo) SaveEmptyTXT(txt string) {
	file, err := os.Create(tr.reserveCopyPath)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}

	defer file.Close()
	file.WriteString(txt)
}

func (tr *BackUpCopyRepo) SaveForLearningTxt(s models.Slovarick) {
	file, err := os.Create(tr.reserveCopyPath)
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