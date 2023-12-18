package config

import (
	"github.com/antoine2116/go-eshop-on-containers/internal/pkg/config"
	typeMapper "github.com/antoine2116/go-eshop-on-containers/internal/pkg/reflection/type_mapper"
	"github.com/iancoleman/strcase"
)

var optionName = strcase.ToLowerCamel(typeMapper.GetTypeNameByT[GinHttpServerOptions]())

type GinHttpServerOptions struct {
	Port string `mapstructure:"port" validate:"required"`
}

func ProvideConfig() (*GinHttpServerOptions, error) {
	return config.BindConfigKey[*GinHttpServerOptions](optionName)
}
