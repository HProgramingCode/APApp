package repositories

import (
	"main/internal/models"

	"gorm.io/gorm"
)

type TestDataRepository struct {
	db *gorm.DB
}

type ITestDataRepository interface {
	CreateImportData(importData *models.ImportData) error
	GetImportDataListByUserId(userId uint) ([]models.ImportData, error)
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

func (r *TestDataRepository) GetImportDataListByUserId(userId uint) ([]models.ImportData, error) {
	var importData []models.ImportData
	err := r.db.Preload("Records").Where("user_id = ?", userId).Find(&importData).Error
	if err != nil {
		return nil, err
	}

	return importData, nil
}
