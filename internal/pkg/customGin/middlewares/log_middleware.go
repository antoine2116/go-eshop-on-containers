package middlewares

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func GinLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		req := ctx.Request
		ctx.Next()

		end := time.Now()
		latency := end.Sub(start)

		fields := []zap.Field{
			zap.Int("status", ctx.Writer.Status()),
			zap.String("method", req.Method),
			zap.String("path", req.URL.Path),
			zap.String("query", req.URL.RawQuery),
			zap.String("ip", ctx.ClientIP()),
			zap.String("user-agent", req.UserAgent()),
			zap.Duration("latency", latency),
		}

		if len(ctx.Errors) > 0 {
			for _, err := range ctx.Errors.Errors() {
				logger.Error(err, fields...)
			}
		} else {
			logger.Info(req.URL.Path, fields...)
		}
	}
}
