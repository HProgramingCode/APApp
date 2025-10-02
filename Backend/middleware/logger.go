package logger

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func InitLogger(env string) error {
	var err error
	switch env {
	case "production":
		Log, err = zap.NewProduction()
		if err != nil {
			return err
		}
	case "development":
		cfg := zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		Log, err = cfg.Build()
		if err != nil {
			return err
		}
	default:
		cfg := zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		Log, err = cfg.Build()
		if err != nil {
			return err
		}
	}
	return nil
}

func GinZapMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next() // ハンドラーを実行

		status := c.Writer.Status()
		Log.Info("request completed",
			zap.String("method", method),
			zap.String("path", path),
			zap.Int("status", status),
		)
	}
}
func Info(msg string, fields ...zap.Field) {
	Log.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Log.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	Log.Fatal(msg, fields...)
}
