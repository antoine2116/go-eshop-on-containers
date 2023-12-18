package gormPostgres

import (
	"github.com/antoine2116/go-eshop-on-containers/internal/pkg/gorm/config"
	"go.uber.org/fx"
)

var (
	Module = fx.Module(
		"gormfx",

		fx.Provide(
			config.ProvideConfig,
			NewGorm,
		),
	)
)
