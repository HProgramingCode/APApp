package main

import (
	"fmt"
	"log"
	"main/config"
	"main/infra"
	logger "main/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {

	config.LoadConfig()
	infra.InitDB()

	if err := logger.Init(config.AppConfig.Env); err != nil {
		log.Fatalf("failed to init logger: %v", err)
	}
	defer logger.Log.Sync()
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(GinZapMiddleware())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	addr := fmt.Sprintf(":%s", config.AppConfig.Port)
	r.Run(addr)
}

func GinZapMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next() // ハンドラーを実行

		status := c.Writer.Status()
		logger.Log.Info("request completed",
			zap.String("method", method),
			zap.String("path", path),
			zap.Int("status", status),
		)
	}
}
