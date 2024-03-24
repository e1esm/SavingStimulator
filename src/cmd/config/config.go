package config

import (
	"github.com/caarlos0/env/v10"
)

type Config struct {
	WebPort  int `env:"PORT" envDefault:"8000"`
	Database struct {
		Address  string `env:"ADDRESS,required"`
		Port     int    `env:"PORT,required"`
		User     string `env:"USER,required"`
		Password string `env:"PASSWORD,required"`
		DB       string `env:"DATABASE:required"`
	} `envPrefix:"DB_"`
	Cache struct {
		Address string `env:"ADDRESS,required"`
		Port    int    `env:"PORT,required"`
		DB      int    `env:"DATABASE,required"`
	} `envPrefix:"CACHE_"`
}

func New() *Config {
	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		return nil
	}

	return cfg
}
