package controllers

import (
	"main/middleware"
	"main/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type IUserCotroller interface {
	GetUserInfo(ctx *gin.Context)
}

type UserController struct {
	service services.IUserService
}

func NewUserController(service services.IUserService) IUserCotroller {
	return &UserController{service: service}
}

func (c *UserController) GetUserInfo(ctx *gin.Context) {
	email, exsist := ctx.Get("email")
	if !exsist {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		middleware.Log.Error("Unauthorized")
		ctx.Abort()
		return
	}

	user, err := c.service.FindUser(email.(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		middleware.Log.Error("failed to find user", zap.Error(err))
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"email": user.Email,
		"id":    user.ID,
	})
}
