package services

import (
	"encoding/csv"
	"main/models"
	"main/repositories"
	"mime/multipart"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

type IImportService interface {
	ImportCSV(userId uint, file *multipart.FileHeader) error
}

type ImportService struct {
	repo repositories.ITestDataRepository
}

func NewImportService(repo repositories.ITestDataRepository) IImportService {
	return &ImportService{repo: repo}
}

func (s *ImportService) ImportCSV(userId uint, file *multipart.FileHeader) error {
	f, err := file.Open()
	if err != nil {
		return err
	}
	defer f.Close()

	reader := transform.NewReader(f, japanese.ShiftJIS.NewDecoder())
	csvReader := csv.NewReader(reader)
	csvReader.FieldsPerRecord = -1

	records, err := csvReader.ReadAll()
	if err != nil {
		return err
	}

	var recordList []models.Record
	for _, row := range records[1:] { // 1行目はヘッダーを除外
		// 正誤欄（例: ○ / ×）を bool に変換
		isCorrect := false
		if row[1] == "○" {
			isCorrect = true
		}

		record := models.Record{
			No:           row[0],
			IsCorrect:    isCorrect,
			FieldName:    row[2],
			MainCategory: row[3],
			SubCategory:  row[4],
			SourceURL:    row[5],
		}

		recordList = append(recordList, record)
	}

	importData := models.ImportData{
		FileName: file.Filename,
		UserID:   userId,
		Records:  recordList,
	}

	err = s.repo.CreateImportData(&importData)
	if err != nil {
		return err
	}

	return nil

}
