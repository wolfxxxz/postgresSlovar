package repositories

/*
var collectionRepoTest = "words"


type repoWordsPg struct {
	db         *sqlx.DB
	collection string
	log        *logrus.Logger
}

func NewRepoWordsPg(db *sqlx.DB, log *logrus.Logger) RepoWordsPg {
	return &repoWordsPg{db: db, collection: collectionRepoTest, log: log}
}

func (rt *repoWordsPg) GetAllWords() ([]*models.Word, error) {
	words := []*models.Word{}
	err := rt.db.Select(&words, "SELECT * FROM words order by theme")
	if err != nil {
		appErr := apperrors.GetAllWordsErr.AppendMessage(err)
		rt.log.Error(appErr)
		return nil, appErr
	}

	return words, nil
}

func (rt *repoWordsPg) InsertWord(ctx context.Context, word *models.Word) error {
	var insertedId int
	query := "INSERT INTO words (english, russian, theme, right_answer) VALUES ($1, $2, $3, $4) RETURNING id"
	err := rt.db.QueryRow(query, word.English, word.Russian, word.Theme, word.RightAnswer).Scan(&insertedId)
	if err != nil {
		appErr := apperrors.InsertWordErr.AppendMessage(err)
		rt.log.Error(appErr)
		return appErr
	}

	rt.log.Info(insertedId)
	word.Id = insertedId
	return nil
}

func (rt *repoWordsPg) GetTranslationRus(word string) ([]*models.Word, error) {
	var words []*models.Word
	err := rt.db.Select(&words, "SELECT * FROM words where russian=$1", word)
	if err != nil {
		appErr := apperrors.GetWordsWhereRAErr.AppendMessage(err)
		rt.log.Error(appErr)
		return nil, appErr
	}

	return words, nil
}

func (rt *repoWordsPg) GetTranslationEngl(word string) ([]*models.Word, error) {
	var words []*models.Word
	err := rt.db.Select(&words, "SELECT * FROM words where english=$1", word)
	if err != nil {
		appErr := apperrors.GetWordsWhereRAErr.AppendMessage(err)
		rt.log.Error(appErr)
		return nil, appErr
	}

	return words, nil
}

func (rt *repoWordsPg) GetTranslationEnglLike(word string) ([]*models.Word, error) {
	var words []*models.Word
	err := rt.db.Select(&words, "SELECT * FROM words WHERE english LIKE '%' || $1 || '%'", word)
	if err != nil {
		appErr := apperrors.GetWordsWhereRAErr.AppendMessage(err)
		rt.log.Error(appErr)
		return nil, appErr
	}

	return words, nil
}

func (rt *repoWordsPg) GetTranslationRusLike(word string) ([]*models.Word, error) {
	var words []*models.Word
	err := rt.db.Select(&words, "SELECT * FROM words WHERE russian LIKE '%' || $1 || '%'", word)
	if err != nil {
		appErr := apperrors.GetWordsWhereRAErr.AppendMessage(err)
		rt.log.Error(appErr)
		return nil, appErr
	}

	return words, nil
}

func (rt *repoWordsPg) CheckWordByEnglish(word *models.Word) (int, error) {
	var id int
	query := "SELECT id FROM words WHERE english=$1"
	err := rt.db.QueryRow(query, word.English).Scan(&id)
	if err != nil {
		appErr := apperrors.CheckWordByEnglishErr.AppendMessage(err)
		rt.log.Info(appErr)
		return 0, nil
	}

	return id, fmt.Errorf("word [%v] is already exist in DB", word.English)
}

func (rt *repoWordsPg) GetWordsWhereRA(quantity int) ([]*models.Word, error) {
	words := []*models.Word{}
	err := rt.db.Select(&words, "SELECT * FROM words order by right_answer limit $1", quantity)
	if err != nil {
		appErr := apperrors.GetWordsWhereRAErr.AppendMessage(err)
		rt.log.Error(appErr)
		return nil, appErr
	}

	return words, nil
}

func (rt *repoWordsPg) UpdateRightAnswer(word *models.Word) error {
	res, err := rt.db.Exec("update words set right_answer=$1 where id=$2", word.RightAnswer, word.Id)
	if err != nil {
		appErr := apperrors.UpdateRightAnswerErr.AppendMessage(err)
		rt.log.Error(appErr)
		return appErr
	}

	_, err = res.RowsAffected()
	if err != nil {
		appErr := apperrors.UpdateRightAnswerErr.AppendMessage(err)
		rt.log.Error(appErr)
		return appErr
	}

	return nil
}

func (rt *repoWordsPg) UpdateWord(word *models.Word) error {
	res, err := rt.db.Exec("update words set english=$1, russian=$2, theme=$3 where id=$4",
		word.English, word.Russian, word.Theme, word.Id)
	if err != nil {
		appErr := apperrors.UpdateWordErr.AppendMessage(err)
		rt.log.Error(appErr)
		return appErr
	}

	_, err = res.RowsAffected()
	if err != nil {
		appErr := apperrors.UpdateWordErr.AppendMessage(err)
		rt.log.Error(appErr)
		return appErr
	}

	return nil
}

func (rt *repoWordsPg) GetWordsMap() (*map[string][]string, error) {
	var words models.Slovarick
	err := rt.db.Select(&words, "SELECT english, russian FROM words order by russian")
	if err != nil {
		appErr := apperrors.GetWordsMapErr.AppendMessage(err)
		rt.log.Error(appErr)
	}

	wordsLib := words.CreateAndInitMapWords()
	return wordsLib, nil
}

func (rt *repoWordsPg) InsertWordsLibrary(ctx context.Context, library []*models.Word) error {

	return nil
}
*/
