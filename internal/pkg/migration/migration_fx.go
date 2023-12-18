package migration

import (
	"github.com/antoine2116/go-eshop-on-containers/internal/pkg/migration/config"
	"go.uber.org/fx"
)

var (
	Module = fx.Module(
		"migrationfx",

		fx.Provide(
			config.ProvideConfig,
			NewMigrator,
		),
	)
)
