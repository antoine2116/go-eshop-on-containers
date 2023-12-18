package config

import "github.com/antoine2116/go-eshop-on-containers/internal/pkg/config"

type Config struct {
	AppOptions AppOptions `mapstructure:"appOptions"`
}

func NewConfig() (*Config, error) {
	cfg, err := config.BindConfig[*Config]()
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

type AppOptions struct {
	ServiceName string `mapstructure:"serviceName"`
}
