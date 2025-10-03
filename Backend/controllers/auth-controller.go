package controllers

import (
	"main/dto"
	"main/services"
	"net/http"

	logger "main/middleware"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type IAuthController interface {
	Signup(c *gin.Context)
	Login(c *gin.Context)
}

type AuthController struct {
	service services.IAuthService
}

func NewAuthController(service services.IAuthService) IAuthController {
	return &AuthController{service: service}
}

func (c *AuthController) Signup(ctx *gin.Context) {
	var input dto.SignupInput
	if err := ctx.ShouldBindBodyWithJSON(&input); err != nil {
		logger.Error("failed to bind signup input", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := c.service.Signup(input.Email, input.Password)
	if err != nil {
		logger.Error("failed to signup", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.Status(http.StatusCreated)
}

func (c *AuthController) Login(ctx *gin.Context) {
	var input dto.LoginInput
	err := ctx.ShouldBindBodyWithJSON(&input)
	if err != nil {
		logger.Error("failed to bind login input", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = c.service.Login(input.Email, input.Password)
	if err != nil {
		logger.Error("failed to login", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.Status(http.StatusCreated)
}
