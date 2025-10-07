package routers

import (
	"main/controllers"
	"main/middleware"

	"github.com/gin-gonic/gin"
)

type RouterSetting struct {
	Auth controllers.IAuthController
}

func SetupRouter(r *gin.Engine, setting RouterSetting) {
	r.GET("/", controllers.Home)
	r.GET("/example", middleware.AuthMiddleware, controllers.Example)

	auth := r.Group("/auth")
	{
		auth.POST("/signup", setting.Auth.Signup)
		auth.POST("/login", setting.Auth.Login)
	}
}
