package words

import (
	"bufio"
	"fmt"
	"os"
	"postgresTakeWords/pkg/repository"
	"strconv"
	"strings"
	"time"

	"github.com/agnivade/levenshtein"
)

var DB = repository.Conf.DB

func WorkTest(s *[]Word) *[]Word {
	startTime := time.Now()
	maps, err := GetWordsMap(DB)
	if err != nil {
		fmt.Println(err)
	}

	var LearnSlice []Word
	quantity := len(*s)
	var capSlovar int = quantity
	NewSlovar := make(Slovarick, capSlovar)
	copy(NewSlovar, (*s)[:quantity])
	*s = (*s)[quantity:]
	fmt.Println("                     START")
	var yes int
	var not int
	var exit1 bool
	fmt.Println("range NewSlovar")
	for _, v := range NewSlovar {
		if exit1 {
			not++
			Preppend(s, v)
			continue
		}

		y, n, exit := Compare(v, maps)
		if exit {
			exit1 = true
			n = 1
		}
		if y > 0 && n > 0 {
			yes++
			Preppend(s, v)
			Preppend(&LearnSlice, v)
			continue
		}

		if y > 0 {
			yes++
			v.RightAnswer += 1
			UpdateRightAnswer(DB, &v)
			AppendWord(s, v)
		} else if n > 0 {
			not++
			Preppend(s, v)
			Preppend(&LearnSlice, v)
		} else {
			break
		}
	}
	duration := time.Since(startTime)
	PrintTime(duration)
	sStat := NewStatistick(quantity, yes, not)
	sStat.WriteStatistic("statistic.txt")
	fmt.Println(yes, not)
	return &LearnSlice
}

func LearnWords(s []Word) bool {
	maps, err := GetWordsMap(DB)
	if err != nil {
		fmt.Println(err)
	}

	startTime := time.Now()
	fmt.Println("                 Learn Words")
	for {
		if len(s) == 0 {
			break
		}
		v := s[0]
		y, _, exit := Compare(v, maps)
		if exit {
			break
		}

		if y > 0 && len(s) != 1 {
			s = s[1:]
		} else if y < 1 {
			copy(s, s[1:])
			s[len(s)-1] = v
			PrintXpen(v.English)
		} else if y > 0 && len(s) == 1 {
			break
		}
	}
	duration := time.Since(startTime)
	PrintTime(duration)
	return true
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

func Compare(l Word, mapWord *map[string][]string) (yes int, not int, exit bool) {
	fmt.Println(l.Russian, " ||Тема: ", l.Theme)
	c := IgnorSpace(l.English)

	a, _ := ScanStringOne()
	if a == "exit" {
		exit = true
		return yes, not, exit
	}

	s := IgnorSpace(a)
	if strings.EqualFold(c, s) {
		yes++
		fmt.Println("Yes")
	} else if CompareWithMap(l.Russian, s, mapWord) {
		fmt.Println("Не совсем правильно ")
		y, n, exit := Compare(l, mapWord)
		if y == 1 {
			yes++
		} else if n == 1 || exit {
			not++
		}
	} else if compareStringsLevenshtein(c, s) {
		yes++
		fmt.Println("Не совсем правильно ", l.English)
	} else {
		not++
		fmt.Println("Incorect:", l.English)
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

/*
// Сравнение строк
func CompareTime(l Word) (yes int, not int) {
	fmt.Println(l.Russian, " ||Тема: ", l.Theme)
	c := ""
	//Игнорировать пробелы
	for _, v := range l.English {
		if v != ' ' {
			c = c + string(v)
		}
	}
	var a string
	s := ""
	//Mistake-----------------------------------------------------------
	go ScanTime(&a)
	time.Sleep(10 * time.Second)
	for _, v := range a {
		if v != ' ' {
			s = s + string(v)
		}
	}

	if strings.EqualFold(c, s) {
		yes++
		fmt.Println("Yes")
	} else {
		not++
		fmt.Println("Incorect:", l.English)
	}
	return yes, not

}*/
