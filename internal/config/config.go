package config

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	PORT     int      `en:"PORT"`
	HOST     string   `en:"HOST"`
	Database Database `en:"prefix=MONGO_"`
}

type Database struct {
	Host     string `env:"HOST"`
	Port     int    `env:"PORT"`
	Username string `env:"USERNAME"`
	Password string `env:"PASSWORD"`
}

func NewConfig() *Config {
	var cfg Config
	if err := envconfig.ProcessWith(context.Background(), &cfg, envconfig.OsLookuper()); err != nil {
		panic(err)
	}

	return &cfg
}
