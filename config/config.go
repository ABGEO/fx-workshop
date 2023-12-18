package config

import (
	"errors"
	"fmt"

	"github.com/mcuadros/go-defaults"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	ListenAddr     string   `mapstructure:"SERVER_ADDR" default:"0.0.0.0"`
	Port           string   `mapstructure:"SERVER_PORT" default:"8080"`
	TrustedProxies []string `mapstructure:"TRUSTED_PROXIES"`
}

type DatabaseConfig struct {
	FilePath string `mapstructure:"DB_FILE_PATH" default:"main.db"`
}

type Config struct {
	Env      string         `mapstructure:"ENV" default:"prod"`
	Server   ServerConfig   `mapstructure:",squash"`
	Database DatabaseConfig `mapstructure:",squash"`
}

func New() (*Config, error) {
	conf := new(Config)

	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AllowEmptyEnv(true)
	viper.AutomaticEnv()

	defaults.SetDefaults(conf)

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError

		if !errors.As(err, &configFileNotFoundError) {
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}
	}

	if err := viper.Unmarshal(conf); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return conf, nil
}
