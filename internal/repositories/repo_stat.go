package repositories

import (
	"bufio"
	"fmt"
	"os"
	"postgresTakeWords/internal/models"

	"github.com/sirupsen/logrus"
)

type Stat struct {
	path string
	log  *logrus.Logger
}

func NewStatRepo(path string, log *logrus.Logger) *Stat {
	return &Stat{path: path, log: log}
}

func (st Stat) Println(stat models.Statistick) {
	fmt.Printf("Right Answer: %v\nWrong Answer: %v\nYour Average is: %v\n",
		stat.RightAnswer, stat.WrongAnswer, stat.Average)
	fmt.Println(stat.Data)
}

func (st Stat) WriteStatistic(stat models.Statistick) {
	file, err := os.OpenFile(st.path, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		st.log.Errorf("STAT_OPEN_FILE, path [%v] is empty", st.path)
		file, err = os.Create(st.path)
		if err != nil {
			st.log.Error("Unable to create file:", err)
			os.Exit(1)
		}
	}

	defer file.Close()
	writer := bufio.NewWriter(file)
	newStat := st.lineStatistic(stat)
	_, err = writer.WriteString(newStat + "\n")
	if err != nil {
		st.log.Error("STAT_REPO_ERR ", err)
		return
	}

	err = writer.Flush()
	if err != nil {
		st.log.Error("STAT_REPO_ERR ", err)
		return
	}
}

func (st Stat) lineStatistic(stat models.Statistick) string {
	todaysStat := fmt.Sprintf("Words tested: %v || Right Answer: %v || Wrong Answer: %v || Average is: %v%v ||",
		stat.WordsTested, stat.RightAnswer, stat.WrongAnswer, stat.Average, "%") + stat.Data
	return todaysStat
}
