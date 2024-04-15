package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

// Config struct
type Config struct {
	HTTPServer   `yaml:"http_server"`
	DataProvider `yaml:"data_provider"`
}

// HTTPServer struct
type HTTPServer struct {
	//Address     string        `yaml:"address" env-default:"localhost:8080"`
	//Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	//IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
	//User        string        `yaml:"user" env-required:"true"`
	//Password    string        `yaml:"password" env-required:"true" env:"HTTP_SERVER_PASSWORD"`
}

// DataProvider struct
type DataProvider struct {
	Secret string `yaml:"secret" env-required:"true"`
}

// Load config
func Load() (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadConfig("config/prod.yaml", &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
