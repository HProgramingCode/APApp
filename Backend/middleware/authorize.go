package middleware

import (
	"main/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {
	tokenString, err := ctx.Cookie("token")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
			"token": tokenString,
		})
		ctx.Abort()
		return
	}

	token, err := utils.ParseToken(tokenString)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid token",
		})
		ctx.Abort()
		return
	}

	claims, err := utils.GetTokenClaims(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid token",
		})
		ctx.Abort()
		return
	}

	email := claims["email"].(string)
	userId := uint(claims["sub"].(float64))

	ctx.Set("email", email)
	ctx.Set("userId", userId)

	ctx.Next()
	ctx.Next()
}
