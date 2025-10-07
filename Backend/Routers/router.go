package routers

import (
	"main/controllers"
	"main/middleware"

	"github.com/gin-gonic/gin"
)

type RouterSetting struct {
	Auth controllers.IAuthController
	User controllers.IUserCotroller
}

func SetupRouter(r *gin.Engine, setting RouterSetting) {
	r.GET("/", controllers.Home)
	r.GET("/example", middleware.AuthMiddleware, controllers.Example)

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
	}
}
