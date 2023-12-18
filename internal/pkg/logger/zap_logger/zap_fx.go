package zapLogger

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewZapLogger() *zap.Logger {
	// TODO : don't forget to change to zap.NewProduction()
	logger, _ := zap.NewDevelopment()
	return logger
}

var Module = fx.Module(
	"zapfx",
	fx.Provide(NewZapLogger),
)
