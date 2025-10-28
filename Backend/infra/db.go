package infra

import (
	"main/config"
	logger "main/internal/middleware"

	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	var err error

	DB, err = gorm.Open(sqlite.Open(config.AppConfig.DBName), &gorm.Config{})
	if err != nil {
		return err
	}

	logger.Info("Database connected",
		zap.String("driver", "sqlite3"),
		zap.String("name", config.AppConfig.DBName),
	)
	return nil
}
