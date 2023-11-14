package words

import "fmt"

func UpdateLibrary(NewWords *[]Word, oldWords *[]Word) {
	c := len(*oldWords)
	*oldWords = append(*NewWords, *oldWords...)
	d := len(*oldWords)
	if d != c {
		fmt.Println("                   New Words Add:", d-c)
	} else {
		fmt.Println("Для загрузки слов списком необходимо упорядочить и вставить слова в файл `save/newWords.txt`")
		fmt.Println("english - перевод - тема")
		fmt.Println("в конце оставить пустую строчку")
		fmt.Println("I believe in you!!!")
	}
}

func (oldWords Slovarick) UpdateLibraryOnlyNewWords(NewWords Slovarick) {
	c := len(oldWords)
	oldWords = append(NewWords, oldWords...)
	d := len(oldWords)
	if d != c {
		fmt.Println("                   New Words Add:", d-c)
	} else {
		fmt.Println("Для загрузки слов списком необходимо упорядочить и вставить слова в файл `save/newWords.txt`")
		fmt.Println("english - перевод - тема")
		fmt.Println("в конце оставить пустую строчку")
		fmt.Println("I believe in you!!!")
	}
}
