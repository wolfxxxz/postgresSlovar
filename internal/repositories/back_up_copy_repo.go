package repositories

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"postgresTakeWords/internal/apperrors"
	"postgresTakeWords/internal/models"

	"github.com/sirupsen/logrus"
	"github.com/tealeg/xlsx"
)

const xlsxPath = "save/library.xlsx"

type BackUpCopyRepo struct {
	reserveCopyPath    string
	reserveCopyPathTXT string
	copyPathXLSX       string
	log                *logrus.Logger
}

func NewBackUpCopyRepo(path string, pathTXT string, log *logrus.Logger) *BackUpCopyRepo {
	return &BackUpCopyRepo{reserveCopyPath: path, copyPathXLSX: xlsxPath, reserveCopyPathTXT: pathTXT, log: log}
}

func (tr *BackUpCopyRepo) GetAllFromBackUp() ([]*models.Word, error) {
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

	words := []*models.Word{}
	err = json.Unmarshal(data, words)
	if err != nil {
		appErr := apperrors.GetAllFromBackUpErr.AppendMessage(err)
		tr.log.Error(appErr)
		return nil, appErr
	}

	return words, nil
}

func (tr *BackUpCopyRepo) SaveAllAsJson(s []*models.Word) error {
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

func (tr *BackUpCopyRepo) SaveAllAsTXT(s []*models.Word) error {
	file, err := os.Create(tr.reserveCopyPathTXT)
	if err != nil {
		appErr := apperrors.SaveAllAsTXTErr.AppendMessage(fmt.Sprintf("Unable to create file: [%v]", err))
		tr.log.Error(appErr)
		os.Exit(1)
		return err
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

	return nil
}

func (tr *BackUpCopyRepo) GetAllWordsFromBackUpXlsx() ([]*models.Word, error) {
	xlFile, err := xlsx.OpenFile(tr.copyPathXLSX)
	if err != nil {
		appErr := apperrors.GetAllWordsXLSXErr.AppendMessage(err)
		tr.log.Error(appErr)
		return nil, appErr
	}

	wordNew := []*models.Word{}
	// Проход по всем листам в файле
	for _, sheet := range xlFile.Sheets {
		fmt.Println("Sheet Name:", sheet.Name)
		// Проход по всем строкам в листе
		for _, row := range sheet.Rows {
			// Проход по всем ячейкам в строке
			for _, cell := range row.Cells {
				text := cell.String()
				fmt.Printf("%s\t", text)
			}
			fmt.Println()
		}
	}

	return wordNew, nil
}

func (tr *BackUpCopyRepo) SaveWordNewAsXLSX(words []*models.Word) error {
	file := xlsx.NewFile()

	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		appErr := apperrors.SaveAllAsXLSXErr.AppendMessage(err)
		tr.log.Error(err)
		return appErr
	}

	for _, word := range words {
		row := sheet.AddRow()
		cell := row.AddCell()
		cell.SetInt(word.Id)
		cell = row.AddCell()
		cell.Value = word.English
		cell = row.AddCell()
		cell.Value = word.Russian
		cell = row.AddCell()
		cell.Value = word.Theme
		cell = row.AddCell()
		cell.SetInt(word.RightAnswer)
	}

	err = file.Save(tr.copyPathXLSX)
	if err != nil {
		appErr := apperrors.SaveAllAsXLSXErr.AppendMessage(err)
		tr.log.Error(err)
		return appErr
	}

	return nil
}
