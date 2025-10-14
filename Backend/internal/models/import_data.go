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
	ImportDataID uint   // 外部キー
	No           string // 問題番号
	IsCorrect    bool   // 正誤（○: true, ×: false）
	FieldName    string // 分野名
	MainCategory string // 大問分類
	SubCategory  string // 中分類
	SourceURL    string // 出典URL
}
