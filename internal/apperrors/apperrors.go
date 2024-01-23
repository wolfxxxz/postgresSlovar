package apperrors

import "fmt"

type AppError struct {
	Message string
	Code    string
}

func NewAppError() *AppError {
	return &AppError{}
}

var (
	EnvConfigLoadError = AppError{
		Message: "Failed to load env file",
		Code:    EnvInit,
	}
	EnvConfigParseError = AppError{
		Message: "Failed to parse env file",
		Code:    EnvParse,
	}
	InitPostgressErr = AppError{
		Message: "Failed to InitPostgress",
		Code:    EnvParse,
	}
	NewLoggerErr = AppError{
		Message: "Failed to NewLog",
		Code:    Log,
	}
	GetAllFromBackUpErr = AppError{
		Message: "Failed to GetAllFromBackUp",
		Code:    BackUpRepo,
	}
	SaveAllAsJsonErr = AppError{
		Message: "Failed to SaveAllAsJson",
		Code:    BackUpRepo,
	}
	SaveAllAsTXTErr = AppError{
		Message: "Failed to SaveAllAsTXT",
		Code:    BackUpRepo,
	}
	InsertWordLearnErr = AppError{
		Message: "Failed to InsertWordLearn",
		Code:    LearnRepo,
	}
	GetWordsLearnErr = AppError{
		Message: "Failed to GetWordsLearn",
		Code:    LearnRepo,
	}
	DeleteLearnWordsIdErr = AppError{
		Message: "Failed to DeleteLearnWordsId",
		Code:    LearnRepo,
	}
	GetAllFromTXTErr = AppError{
		Message: "Failed to GetAllFromTXT",
		Code:    NewWordsTXTRepo,
	}
	GetAllWordsXLSXErr = AppError{
		Message: "Failed to GetAllWordsXLSXErr",
		Code:    NewWordsTXTRepo,
	}
	SaveAllAsXLSXErr = AppError{
		Message: "Failed to SaveAllAsXLSXErr",
		Code:    NewWordsTXTRepo,
	}
	CleanNewWordsErr = AppError{
		Message: "Failed to CleanNewWords",
		Code:    NewWordsTXTRepo,
	}
	GetAllWordsErr = AppError{
		Message: "Failed to CleanNewWords",
		Code:    RepoWordsPg,
	}
	CheckWordByEnglishErr = AppError{
		Message: "Failed to CheckWordByEnglish [THIS WORLD IS NEW]",
		Code:    RepoWordsPg,
	}
	InsertWordErr = AppError{
		Message: "Failed to InsertWord",
		Code:    RepoWordsPg,
	}
	GetWordsWhereRAErr = AppError{
		Message: "Failed to GetWordsWhereRA",
		Code:    RepoWordsPg,
	}
	UpdateRightAnswerErr = AppError{
		Message: "Failed to UpdateRightAnswer",
		Code:    RepoWordsPg,
	}
	UpdateWordErr = AppError{
		Message: "Failed to UpdateWord",
		Code:    RepoWordsPg,
	}
	GetWordsMapErr = AppError{
		Message: "Failed to GetWordsMap",
		Code:    RepoWordsPg,
	}
)

func (appError *AppError) Error() string {
	return appError.Code + ": " + appError.Message
}

func (appError *AppError) AppendMessage(anyErrs ...interface{}) *AppError {
	return &AppError{
		Message: fmt.Sprintf("%v : %v", appError.Message, anyErrs),
		Code:    appError.Code,
	}
}

func IsAppError(err1 error, err2 *AppError) bool {
	err, ok := err1.(*AppError)
	if !ok {
		return false
	}

	return err.Code == err2.Code
}
