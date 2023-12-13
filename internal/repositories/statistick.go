package repositories

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"postgresTakeWords/internal/models"
)

type Stat struct {
	path string
}

func NewStatRepo(path string) *Stat {
	return &Stat{path: path}
}

func (st Stat) Println(stat models.Statistick) {
	fmt.Printf("Right Answer: %v\nWrong Answer: %v\nYour Average is: %v\n",
		stat.RightAnswer, stat.WrongAnswer, stat.Average)
	fmt.Println(stat.Data)
}

func (st Stat) StringStatistic(stat models.Statistick) string {
	todaysStat := fmt.Sprintf("Words tested: %v || Right Answer: %v || Wrong Answer: %v || Average is: %v%v ||",
		stat.WordsTested, stat.RightAnswer, stat.WrongAnswer, stat.Average, "%") + stat.Data
	return todaysStat
}

func (st Stat) WriteStatistic(stat models.Statistick) {
	file, err := os.OpenFile(st.path, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		file, err = os.Create(st.path)
		if err != nil {
			fmt.Println("Unable to create file:", err)
			os.Exit(1)
		}
	}

	defer file.Close()
	writer := bufio.NewWriter(file)
	newStat := st.StringStatistic(stat)
	_, err = writer.WriteString(newStat + "\n")
	if err != nil {
		log.Fatal(err)
	}

	err = writer.Flush()
	if err != nil {
		log.Fatal(err)
	}
}

func (st Stat) DecodeJson(stat models.Statistick) {
	filejson, err := os.Open(st.path)
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

func (st Stat) EncodeJson(stat models.Statistick) {
	byteArr, err := json.MarshalIndent(stat, "", "   ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(st.path, byteArr, 0666)
	if err != nil {
		log.Fatal(err)
	}
}
