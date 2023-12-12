package words

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"postgresTakeWords/internal/domain/models"
	"strings"
	"unicode"
)

func DecodeJsonSliceWord(SliceWord *[]models.Word, file string) {
	filejson, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	defer filejson.Close()
	data, err := io.ReadAll(filejson)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(data, SliceWord)
}

func EncodeJson(s *[]models.Word, file string) {
	byteArr, err := json.MarshalIndent(s, "", "   ")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(byteArr))             0664
	err = os.WriteFile(file, byteArr, 0666) //-rw-rw-rw-
	if err != nil {
		log.Fatal(err)
	}
}

func DecodeTXT(s *[]models.Word, filetxt string) {
	data, err := os.ReadFile(filetxt)
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

func EncodeTXT(files string, s models.Slovarick) {
	file, err := os.Create(files)
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

func SaveEmptyTXT(files string, txt string) {
	file, err := os.Create(files)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}

	defer file.Close()
	file.WriteString(txt)
}

func SaveForLearningTxt(files string, s models.Slovarick) {
	file, err := os.Create(files)
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

func capitalizeFirstRune(str string) string {
	runes := []rune(str)
	for i, r := range runes {
		if i == 0 {
			runes[i] = unicode.ToUpper(r)
		}
	}

	return string(runes)
}
