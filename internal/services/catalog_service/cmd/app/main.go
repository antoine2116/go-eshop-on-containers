package main

import (
	"github.com/antoine2116/go-eshop-on-containers/internal/pkg/customGin"
	"github.com/antoine2116/go-eshop-on-containers/internal/pkg/migration"
	catalog "github.com/antoine2116/go-eshop-on-containers/internal/services/catalogservice/internal"
	"github.com/antoine2116/go-eshop-on-containers/internal/services/catalogservice/internal/configurations/mappings"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	app := fx.New(
		catalog.Module,

		// Logger
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),

		// Migrate database
		fx.Invoke(func(m migration.Runner, logger *zap.Logger) {
			if err := m.Up(); err != nil {
				logger.Fatal("Error migrating the database :", zap.Error(err))
			}
		}),

		// Map the endpoints
		fx.Invoke(
			fx.Annotate(
				func(endpoints []customGin.Endpoint) {
					for _, ep := range endpoints {
						ep.MapEndpoint()
					}
				},
				fx.ParamTags(`group:"catalog-route"`),
			)),

		// Configure mappings
		fx.Invoke(func() error {
			err := mappings.ConfigureMappings()
			if err != nil {
				return err
			}
			return nil
		}),
	)

	app.Run()
}
