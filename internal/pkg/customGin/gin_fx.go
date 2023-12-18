package customGin

import (
	"context"
	"github.com/antoine2116/go-eshop-on-containers/internal/pkg/customGin/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var (
	Module = fx.Module(
		"ginfx",

		fx.Provide(
			config.ProvideConfig,
			NewGinHttpServer,
		),
		fx.Invoke(registerHook),
	)
)

func registerHook(lc fx.Lifecycle, ginServer GinHttpServer, logger *zap.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := ginServer.Run(); err != nil {
					logger.Fatal("Failed to start server:", zap.Error(err))
				}
				logger.Info("Server is listening on:", zap.String("addr", ginServer.Cfg().Port))
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := ginServer.GracefulShutdown(ctx); err != nil {
				logger.Error("Error shutting down server:", zap.Error(err))
			} else {
				logger.Info("Server shutdown gracefully")
			}

			return nil
		},
	})
}
