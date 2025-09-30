package infra

import (
	"main/config"
	logger "main/middleware"

	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error

	DB, err = gorm.Open(sqlite.Open(config.AppConfig.DBName), &gorm.Config{})
	if err != nil {
		logger.Log.Fatal("failed to connect database", zap.Error(err))
	}

	logger.Log.Info("Database connected",
		zap.String("driver", "sqlite3"),
		zap.String("name", config.AppConfig.DBName),
	)
}
