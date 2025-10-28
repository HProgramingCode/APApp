package main

import (
	"fmt"
	"main/config"
	"main/infra"
	"main/internal/controllers"
	logger "main/internal/middleware"
	"main/internal/repositories"
	"main/internal/services"
	"main/migrations"

	routers "main/internal/routers"

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

	userRepository := repositories.NewUserRepository(infra.DB)

	authService := services.NewAuthService(userRepository)
	authController := controllers.NewAuthController(authService)

	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	csvRepository := repositories.NewTestDataRepository(infra.DB)
	csvService := services.NewImportService(csvRepository)
	csvController := controllers.NewImportController(csvService)

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(logger.GinZapMiddleware())

	routers.SetupRouter(r, routers.RouterSetting{
		Auth: authController,
		User: userController,
		CSV:  csvController,
	})
	fmt.Printf("CSV Controller type: %T\n", csvController)

	return r, nil
}
