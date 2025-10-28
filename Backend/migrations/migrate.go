package migrations

import (
	"main/config"
	"main/infra"
	logger "main/internal/middleware"
	"main/internal/models"

	"go.uber.org/zap"
)

func Migrate() error {
	err := infra.DB.AutoMigrate(&models.User{}, &models.ImportData{}, &models.Record{})
	if err != nil {
		return err
	}
	logger.Info("Database connected",
		zap.String("driver", "sqlite3"),
		zap.String("name", config.AppConfig.DBName),
	)
	return nil
}
