package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress string `mapstructure:"server_address"`
	BaseUrl       string `mapstructure:"base_url"`
}

func ParseFlags() (*Config, error) {
	pflag.String("a", ":8080", "Адрес запуска HTTP сервера")
	pflag.String("b", "example.com/", "Базовый адрес для сокращенных URL")
	pflag.Parse()

	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		return nil, err
	}
	viper.BindPFlag("server_address", pflag.Lookup("a"))
	viper.BindPFlag("base_url", pflag.Lookup("b"))

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
