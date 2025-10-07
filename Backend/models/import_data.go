package models

import "gorm.io/gorm"

type ImportData struct {
	gorm.Model
	FileName string `gorm:"not null"`
	UserID   uint
	User     User
	Records  []Record `gorm:"foreignKey:ImportDataID"`
}

type Record struct {
	gorm.Model
	ImportDataID uint
	ExamName     string  // 試験名や回数
	Category     string  // 分野（例：ネットワーク、セキュリティなど）
	Score        float64 // 得点率など
	Date         string  // 実施日（CSVによる）
}
