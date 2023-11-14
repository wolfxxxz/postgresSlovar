package words

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type Statistick struct {
	Data        string  `json:"data"`
	WordsTested int     `json:"allWords"`
	RightAnswer int     `json:"right"`
	WrongAnswer int     `json:"wrong"`
	Average     float64 `json:"average"`
}

func NewStatistick(AllWords, NewRight, NewWrong int) *Statistick {
	NewData := timeStamp()
	NewAverage := (float64(NewRight) / float64(AllWords)) * 100
	return &Statistick{Data: NewData, WordsTested: AllWords, RightAnswer: NewRight, WrongAnswer: NewWrong, Average: NewAverage}
}

func timeStamp() string {
	c := time.Now()
	return fmt.Sprintf("%02.f %v %v %02.f:%02.f:%02.f",
		float64(c.Day()), c.Month(), c.Year(), float64(c.Hour()),
		float64(c.Minute()), float64(c.Second()))
}

func (stat Statistick) Println() {
	fmt.Printf("Right Answer: %v\nWrong Answer: %v\nYour Average is: %v\n",
		stat.RightAnswer, stat.WrongAnswer, stat.Average)
	fmt.Println(stat.Data)
}

func (st Statistick) StringStatistic() (stat string) {
	stat = fmt.Sprintf("Words tested: %v || Right Answer: %v || Wrong Answer: %v || Average is: %v%v ||",
		st.WordsTested, st.RightAnswer, st.WrongAnswer, st.Average, "%") + st.Data
	return
}
func (stat Statistick) WriteStatistic(files string) {
	file, err := os.OpenFile(files, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		file, err = os.Create(files)
		if err != nil {
			fmt.Println("Unable to create file:", err)
			os.Exit(1)
		}
	}

	defer file.Close()
	writer := bufio.NewWriter(file)
	newStat := stat.StringStatistic()
	_, err = writer.WriteString(newStat + "\n")
	if err != nil {
		log.Fatal(err)
	}

	err = writer.Flush()
	if err != nil {
		log.Fatal(err)
	}
}

func (stat *Statistick) DecodeJson(file string) {
	filejson, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	defer filejson.Close()
	f := make([]byte, 64)
	var data2 string
	for {
		n, err := filejson.Read(f)
		if err == io.EOF {
			break
		}

		data2 = data2 + string(f[:n])
	}

	data := []byte(data2)
	json.Unmarshal(data, &stat)

}

func (stat *Statistick) EncodeJson(file string) {
	byteArr, err := json.MarshalIndent(stat, "", "   ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(file, byteArr, 0666)
	if err != nil {
		log.Fatal(err)
	}
}
