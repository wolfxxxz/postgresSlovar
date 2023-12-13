package competition

import (
	"bufio"
	"fmt"
	"os"
	"postgresTakeWords/internal/models"
	"postgresTakeWords/internal/repositories"
	"strconv"
	"strings"
	"time"

	"github.com/agnivade/levenshtein"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Competition struct {
	stat            *repositories.Stat
	repoLearn       *repositories.RepoLearn
	repoTest        *repositories.RepoTest
	repoTXT         *repositories.ReserveTXTRepo
	repoUpdateByTXT *repositories.UpdateWordsFromTXTRepo
	log             *logrus.Logger
}

func NewCompetition(statPath string, reserveCopyPath string, newWordsPath string, sqlDB *sqlx.DB, log *logrus.Logger) *Competition {
	return &Competition{
		stat:            repositories.NewStatRepo(statPath),
		repoLearn:       repositories.NewRepoLearn(sqlDB),
		repoTest:        repositories.NewRepoTest(sqlDB),
		repoTXT:         repositories.NewReserveTXTRepo(reserveCopyPath),
		repoUpdateByTXT: repositories.NewUpdateWordsFromTXTRepo(newWordsPath),
		log:             log,
	}
}

func (c *Competition) WorkTest(s *[]models.Word) *[]models.Word {
	startTime := time.Now()
	maps, err := c.repoTest.GetWordsMap()
	if err != nil {
		c.log.Error(err)
	}

	var LearnSlice []models.Word
	quantity := len(*s)
	var capSlovar int = quantity
	NewSlovar := make(models.Slovarick, capSlovar)
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
			models.Preppend(s, v)
			continue
		}

		y, n, exit := Compare(v, maps)
		if exit {
			exit1 = true
			n = 1
		}
		if y > 0 && n > 0 {
			yes++
			models.Preppend(s, v)
			models.Preppend(&LearnSlice, v)
			continue
		}

		if y > 0 {
			yes++
			v.RightAnswer += 1
			c.repoTest.UpdateRightAnswer(&v)
			models.AppendWord(s, v)
		} else if n > 0 {
			not++
			models.Preppend(s, v)
			models.Preppend(&LearnSlice, v)
		} else {
			break
		}
	}

	duration := time.Since(startTime)
	PrintTime(duration)
	sStat := models.NewStatistick(quantity, yes, not)
	c.stat.WriteStatistic(*sStat)
	fmt.Println(yes, not)
	return &LearnSlice
}

func (c *Competition) LearnWords(s []models.Word) bool {
	maps, err := c.repoTest.GetWordsMap()
	if err != nil {
		c.log.Error(err)
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

func Compare(l models.Word, mapWord *map[string][]string) (yes int, not int, exit bool) {
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
