package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Server struct {
		Host string `yaml:"host", envconfig:"SERVER_HOST"`
		Port int `yaml:"port", envconfig:"SERVER_PORT"`
	} `yaml:"server"`
}

func FromEnv(cfg *Config) error {
	return envconfig.Process("", cfg)
}