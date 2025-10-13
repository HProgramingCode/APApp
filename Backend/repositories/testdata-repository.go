package repositories

import (
	"main/models"

	"gorm.io/gorm"
)

type TestDataRepository struct {
	db *gorm.DB
}

type ITestDataRepository interface {
	CreateImportData(importData *models.ImportData) error
}

func NewTestDataRepository(db *gorm.DB) ITestDataRepository {
	return &TestDataRepository{db: db}
}

func (r *TestDataRepository) CreateImportData(importData *models.ImportData) error {
	result := r.db.Create(&importData)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
