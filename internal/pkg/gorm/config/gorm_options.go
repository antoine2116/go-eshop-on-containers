package config

import (
	"github.com/antoine2116/go-eshop-on-containers/internal/pkg/config"
	typeMapper "github.com/antoine2116/go-eshop-on-containers/internal/pkg/reflection/type_mapper"
	"github.com/iancoleman/strcase"
)

var optionName = strcase.ToLowerCamel(typeMapper.GetTypeNameByT[GormOptions]())

type GormOptions struct {
	Port     int    `mapstructure:"port"`
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Db       string `mapstructure:"db"`
}

func ProvideConfig() (*GormOptions, error) {
	return config.BindConfigKey[*GormOptions](optionName)
}
