package words

import "fmt"

// Соединяем два среза в один
// Return oldWords
func UpdateLibrary(NewWords *[]Word, oldWords *[]Word) {

	c := len(*oldWords)

	//--------Соединяем два среза в один--------------
	*oldWords = append(*NewWords, *oldWords...)
	//NewWords.Words = append(NewWords.Words, oldWords.Words...)

	d := len(*oldWords)
	// Записать в filetxt пустой

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
	// Проверить на дубликаты

	//--------Соединяем два среза в один--------------
	oldWords = append(NewWords, oldWords...)
	//NewWords.Words = append(NewWords.Words, oldWords.Words...)

	d := len(oldWords)
	// Записать в filetxt пустой

	if d != c {
		fmt.Println("                   New Words Add:", d-c)
	} else {
		fmt.Println("Для загрузки слов списком необходимо упорядочить и вставить слова в файл `save/newWords.txt`")
		fmt.Println("english - перевод - тема")
		fmt.Println("в конце оставить пустую строчку")
		fmt.Println("I believe in you!!!")
	}
}
