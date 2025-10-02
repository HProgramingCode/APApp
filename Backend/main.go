package main

import (
	"fmt"
	routers "main/Routers"
	"main/config"
	"main/controllers"
	"main/infra"
	logger "main/middleware"
	"main/migrations"
	"main/repositories"
	"main/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {

	r, err := setupApp()
	if err != nil {
		logger.Fatal("failed to setup app", zap.Error(err))
	}

	addr := fmt.Sprintf(":%s", config.AppConfig.Port)
	defer logger.Log.Sync()
	r.Run(addr)
}

func setupApp() (*gin.Engine, error) {

	config.LoadConfig()
	if err := logger.InitLogger(config.AppConfig.Env); err != nil {
		return nil, err
	}
	if err := infra.InitDB(); err != nil {
		return nil, err
	}
	if err := migrations.Migrate(); err != nil {
		return nil, err
	}

	authRepository := repositories.NewAuthRepository(infra.DB)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(logger.GinZapMiddleware())

	routers.SetupRouter(r, routers.RouterSetting{
		Auth: authController,
	})

	return r, nil
}
