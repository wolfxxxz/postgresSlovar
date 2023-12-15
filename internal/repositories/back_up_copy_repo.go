package repositories

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"postgresTakeWords/internal/apperrors"
	"postgresTakeWords/internal/models"

	"github.com/sirupsen/logrus"
)

type BackUpCopyRepo struct {
	reserveCopyPath    string
	reserveCopyPathTXT string
	log                *logrus.Logger
}

func NewBackUpCopyRepo(path string, pathTXT string, log *logrus.Logger) *BackUpCopyRepo {
	return &BackUpCopyRepo{reserveCopyPath: path, reserveCopyPathTXT: pathTXT, log: log}
}

func (tr *BackUpCopyRepo) GetAllFromBackUp() (*[]models.Word, error) {
	filejson, err := os.Open(tr.reserveCopyPath)
	if err != nil {
		appErr := apperrors.GetAllFromBackUpErr.AppendMessage(err)
		tr.log.Error(appErr)
		return nil, appErr
	}

	defer filejson.Close()
	data, err := io.ReadAll(filejson)
	if err != nil {
		appErr := apperrors.GetAllFromBackUpErr.AppendMessage(err)
		tr.log.Error(appErr)
		return nil, appErr
	}

	var words []models.Word
	err = json.Unmarshal(data, &words)
	if err != nil {
		appErr := apperrors.GetAllFromBackUpErr.AppendMessage(err)
		tr.log.Error(appErr)
		return nil, appErr
	}

	return &words, nil
}

func (tr *BackUpCopyRepo) SaveAllAsJson(s *[]models.Word) error {
	byteArr, err := json.MarshalIndent(s, "", "   ")
	if err != nil {
		appErr := apperrors.SaveAllAsJsonErr.AppendMessage(err)
		tr.log.Error(appErr)
		return err
	}

	err = os.WriteFile(tr.reserveCopyPath, byteArr, 0666) //-rw-rw-rw- 0664
	if err != nil {
		appErr := apperrors.SaveAllAsJsonErr.AppendMessage(err)
		tr.log.Error(appErr)
		return err
	}

	return nil
}

func (tr *BackUpCopyRepo) SaveAllAsTXT(s *[]models.Word) error {
	file, err := os.Create(tr.reserveCopyPathTXT)
	if err != nil {
		appErr := apperrors.SaveAllAsTXTErr.AppendMessage(fmt.Sprintf("Unable to create file: [%v]", err))
		tr.log.Error(appErr)
		os.Exit(1)
		return err
	}

	defer file.Close()
	for _, v := range *s {
		file.WriteString(v.English)
		file.WriteString(" - ")
		file.WriteString(v.Russian)
		if v.Theme != "" {
			file.WriteString(" - ")
			file.WriteString(v.Theme)
		}

		file.WriteString("\n")
	}

	return nil
}
