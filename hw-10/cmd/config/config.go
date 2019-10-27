package config

import (
	"github.com/caarlos0/env/v6"
)

// AppConfig structure to store all general parameters...
type AppConfig struct {
	Port         int  `env:"PORT" envDefault:"8080"`
	IsProduction bool `env:"PRODUCTION"`
}

// NewConfig return new configuration struct received from the env variables
func NewConfig() (*AppConfig, error) {

	config := new(AppConfig)
	if err := env.Parse(config); err != nil {
		return nil, err
	}

	return config, nil

}
