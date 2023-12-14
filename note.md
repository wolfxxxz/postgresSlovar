// "postgres://localhost:5435/postgres?sslmode=disable&user=postgres&password=1"

type ConfigPostgress struct {
	Token             string `env:"TOKEN"`
	LogLevel          string `env:"LOGGER_LEVEL"`
	SqlHost           string `env:"SQLHost"`   //localhost
	SqlPort           string `env:"SQL_PORT"`  //5435
	SqlType           string `env:"SQL_TYPE"`  //postgres
	SqlMode           string `env:"SQL_MODE"`  //disable
	UserName          string `env:"USER_NAME"` // postgres
	Password          string `env:"PASSWORD"`  //1
	DBName            string `env:"DB_NAME"`
	TimeoutMongoQuery string `env:"TIMEOUT_MONGO_QUERY"`
}

/*
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
}*/


/*
	func (tr *BackUpCopyRepo) DecodeTXT(s *[]models.Word) {
		data, err := os.ReadFile(tr.reserveCopyPath)
		if err != nil {
			tr.log.Error(err)
			return
		}
		content := string(data)
		lines := strings.Split(content, "\n")
		for _, line := range lines {
			if line == "" {
				continue
			}

			words := strings.Split(line, "-")
			if len(words) <= 0 {
				continue
			}

			for i, v := range words {
				words[i] = strings.TrimSpace(v)
				words[i] = strings.ReplaceAll(words[i], ".", "")
				words[i] = capitalizeFirstRune(words[i])
			}

			id := 0
			theme := ""
			if len(words) > 3 {
				theme = words[2]
			}

			word := models.NewLibrary(id, words[0], words[1], theme)
			*s = append(*s, *word)
		}
	}
*/

/*
func (tr *BackUpCopyRepo) SaveEmptyTXT(txt string) {
	file, err := os.Create(tr.reserveCopyPath)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}

	defer file.Close()
	file.WriteString(txt)
}*/

/*

func (tr *BackUpCopyRepo) SaveForLearningTxt(s models.Slovarick) {
	file, err := os.Create(tr.reserveCopyPath)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}

	defer file.Close()
	if len(s) != 0 {
		for _, v := range s {
			file.WriteString(v.English)
			var length = len(v.English)
			if length <= 30 {
				for i := 0; i+length <= 25; i++ {
					file.WriteString(" ")
				}
			}
			file.WriteString(" - ")
			file.WriteString(v.Russian)
			file.WriteString("\n")
		}
	} else {
		fmt.Println("empty learn words")
	}
}*/
/*
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
*/
/*
func Print(l models.Slovarick) {
	for _, v := range l {
		fmt.Print(v.English, " - ", v.Russian, " - ", v.Theme)
		fmt.Println()
	}
}*/

/*

 */
/*
func (tr *UpdateWordsFromTXTRepo) EncodeTXT(s models.Slovarick) {
	file, err := os.Create(tr.newWordsPath)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}

	defer file.Close()
	for _, v := range s {
		file.WriteString(v.English)
		file.WriteString(" - ")
		file.WriteString(v.Russian)
		if v.Theme != "" {
			file.WriteString(" - ")
			file.WriteString(v.Theme)
		}

		file.WriteString("\n")
	}
}*/

/*
func (tr *UpdateWordsFromTXTRepo) DownloadFromJsonAndDecodeModel() *[]models.Word {
	filejson, err := os.Open(tr.newWordsPath)
	if err != nil {
		log.Fatal(err)
	}

	defer filejson.Close()
	data, err := io.ReadAll(filejson)
	if err != nil {
		log.Fatal(err)
	}

	var words []models.Word
	json.Unmarshal(data, &words)
	return &words
}*/
/*

func (tr *UpdateWordsFromTXTRepo) EncodeJson(s *[]models.Word) {
	byteArr, err := json.MarshalIndent(s, "", "   ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(tr.newWordsPath, byteArr, 0666) //-rw-rw-rw- 0664
	if err != nil {
		log.Fatal(err)
	}
}*/

/*
func (tr *UpdateWordsFromTXTRepo) SaveForLearningTxt(s models.Slovarick) {
	file, err := os.Create(tr.newWordsPath)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}

	defer file.Close()
	if len(s) != 0 {
		for _, v := range s {
			file.WriteString(v.English)
			var length = len(v.English)
			if length <= 30 {
				for i := 0; i+length <= 25; i++ {
					file.WriteString(" ")
				}
			}

			file.WriteString(" - ")
			file.WriteString(v.Russian)
			file.WriteString("\n")
		}
	}

	fmt.Println("empty learn words")
}
*/
/*
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

func Print(l models.Slovarick) {
	for _, v := range l {
		fmt.Print(v.English, " - ", v.Russian, " - ", v.Theme)
		fmt.Println()
	}
}

func capitalizeFirstRune(line string) string {
	runes := []rune(line)
	for i, r := range runes {
		if i == 0 {
			runes[i] = unicode.ToUpper(r)
		}
	}

	return string(runes)
}
*/
/*
func (rt *RepoWordsPg) GetWordById(word *models.Word) (*models.Word, error) {
	var getWord models.Word
	err := rt.db.Get(&getWord, "SELECT * FROM words WHERE id=$1", word.Id)
	if err != nil {
		return nil, err
	}
	return &getWord, nil
}*/


func UpdateLibrary(NewWords *[]models.Word, oldWords *[]models.Word) {
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