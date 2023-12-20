package middlewares

import (
	customErrors "github.com/antoine2116/go-eshop-on-containers/internal/pkg/http/custom_errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ErrorHandler(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, e := range c.Errors {
			err := customErrors.GetCustomError(e.Err)
			if err != nil {
				logger.Sugar().Error(err)
				c.AbortWithStatusJSON(err.GetStatus(), err.GetMessage())
			}
		}
	}
}
