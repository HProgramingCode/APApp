package routers

import (
	"main/controllers"

	"github.com/gin-gonic/gin"
)

type RouterSetting struct {
	Auth controllers.IAuthController
}

func SetupRouter(r *gin.Engine, setting RouterSetting) {
	r.GET("/", controllers.Home)
	r.GET("/example", controllers.Example)

	auth := r.Group("/auth")
	{
		auth.POST("/signup", setting.Auth.Signup)
		auth.POST("/login", setting.Auth.Login)
	}

}
