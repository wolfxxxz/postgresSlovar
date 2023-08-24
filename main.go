package main

import (
	"fmt"
	"log"
	"postgresTakeWords/internal/words"
	"postgresTakeWords/pkg/repository"

	_ "github.com/lib/pq"
)

var Conf = repository.Conf

func main() {

	fmt.Println("Conf connect")

	err := words.StartCompetition(Conf.DB)
	if err != nil {
		log.Println(err)
	}

}
