package routers

import (
	"main/internal/controllers"
	"main/internal/middleware"

	"github.com/gin-gonic/gin"
)

type RouterSetting struct {
	Auth controllers.IAuthController
	User controllers.IUserCotroller
	CSV  controllers.IImportController
}

func SetupRouter(r *gin.Engine, setting RouterSetting) {
	auth := r.Group("/auth")
	{
		auth.POST("/signup", setting.Auth.Signup)
		auth.POST("/login", setting.Auth.Login)
	}
	private := r.Group("/private")
	private.Use(middleware.AuthMiddleware)
	user := private.Group("/user")
	{
		user.GET("/info", setting.User.GetUserInfo)
		user.POST("/import", setting.CSV.ImportCSV)
		user.GET("/getdata", setting.CSV.GetImportedCSV)
	}
}
