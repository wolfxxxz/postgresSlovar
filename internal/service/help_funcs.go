package service

import (
	"bufio"
	"fmt"
	"os"
	"postgresTakeWords/internal/models"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/agnivade/levenshtein"
)

func Compare(word *models.Word, mapWord *map[string][]string) (yes int, not int, trExit bool) {
	fmt.Println(word.Russian, " ||Theme:", word.Theme, "||Part of Speech ", word.PartOfSpeech)
	c := IgnorSpace(word.English)

	a, _ := ScanStringOne()
	if a == exit {
		trExit = true
		return yes, not, trExit
	}

	s := IgnorSpace(a)
	if strings.EqualFold(c, s) {
		yes++
		fmt.Println("Yes")
	} else if CompareWithMap(word.Russian, s, mapWord) {
		fmt.Println("Another one ")
		y, n, exit := Compare(word, mapWord)
		if y == 1 {
			yes++
		} else if n == 1 || exit {
			not++
		}
	} else if compareStringsLevenshtein(c, s) {
		yes++
		fmt.Println("Some different ", word.English)
	} else {
		not++
		fmt.Println("Incorect:", word.English)
		for {
			fmt.Println("Please enter correct: ")
			j, _ := ScanStringOne()
			jj := IgnorSpace(j)
			if strings.EqualFold(c, jj) {
				break
			}
		}
	}
	return yes, not, false
}

func compareStringsLevenshtein(str1, str2 string) bool {
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)
	mistakes := 1
	if distance := levenshtein.ComputeDistance(str1, str2); distance <= mistakes {
		return true
	} else {
		return false
	}
}

func IgnorSpace(s string) (c string) {
	for _, v := range s {
		if v != ' ' {
			c = c + string(v)
		}
	}
	return
}

func CompareWithMap(russian, answer string, mapWords *map[string][]string) bool {
	englishWords, ok := (*mapWords)[russian]
	if !ok {
		return false
	}

	for _, word := range englishWords {
		if answer == word {
			return true
		}
	}

	fmt.Println("__________________________________", englishWords)
	return false
}

func ScanTime(a *string) {
	fmt.Print("    ")
	in := bufio.NewScanner(os.Stdin)
	if in.Scan() {
		*a = in.Text()
	}
}

func PrintXpen(s string) {
	var d int
	for i := 0; i <= 15; i++ {
		d++
		for i := 0; i <= d; i++ {
			fmt.Print(" ")
		}

		fmt.Println(s, "   ", s, "   ", s)
	}
}

func PrintTime(duration time.Duration) {
	minutes := int(duration.Minutes())
	seconds := int(duration.Seconds()) % 60
	fmt.Printf("Time: %d minutes %d seconds\n", minutes, seconds)
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

func isCyrillic(s string) bool {
	for _, r := range s {
		if unicode.Is(unicode.Cyrillic, r) {
			return true
		}
	}
	return false
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

func ScanInt() (n int) {
	for {
		cc, _ := ScanStringOne()
		i, err := strconv.Atoi(cc)
		if err != nil {
			fmt.Println("Incorect, please enter number")
		} else {
			n = i
			break
		}
	}

	return
}
