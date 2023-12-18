package config

import (
	typeMapper "github.com/antoine2116/go-eshop-on-containers/internal/pkg/reflection/type_mapper"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func BindConfig[T any]() (T, error) {
	return BindConfigKey[T]("")
}

func BindConfigKey[T any](configKey string) (T, error) {
	configPath, err := getConfigRootPath()
	if err != nil {
		return *new(T), err
	}

	viper.AddConfigPath(configPath)
	viper.SetConfigName("config.dev")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		return *new(T), err
	}

	cfg := typeMapper.GenericInstanceByT[T]()

	if len(configKey) == 0 {
		if err := viper.Unmarshal(cfg); err != nil {
			return *new(T), err
		}
	} else {
		if err := viper.UnmarshalKey(configKey, cfg); err != nil {
			return *new(T), err
		}
	}

	return cfg, nil
}

func getConfigRootPath() (string, error) {
	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Get the absolute path of the executed project directory
	absCurrentDir, err := filepath.Abs(wd)
	if err != nil {
		return "", err
	}

	// Get the path to the "config" folder within the project directory
	configPath := filepath.Join(absCurrentDir, "config")

	return configPath, nil
}
