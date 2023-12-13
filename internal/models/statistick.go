package models

import (
	"fmt"
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
