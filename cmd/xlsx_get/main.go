package main

import (
	"fmt"
	"postgresTakeWords/internal/log"
	"postgresTakeWords/internal/repositories"
)

const (
	libJson  = "save/library.json"
	libTxt   = "save/library.txt"
	newWords = "save/newWords.txt"
)

func main() {
	log, err := log.NewLogAndSetLevel("info")
	if err != nil {
		log.Fatal(err)
	}

	backupRepo := repositories.NewBackUpCopyRepo(libJson, libTxt, log)

	words, err := backupRepo.GetAllWordsFromBackUpXlsx()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(words))
	fmt.Printf("%+v\n", words[0])

	err = backupRepo.SaveWordNewAsXLSX(words)
	if err != nil {
		log.Fatal(err)
	}
}
