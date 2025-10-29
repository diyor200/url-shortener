package config

import (
	"context"
	"fmt"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	PORT     string    `env:"SERVER_PORT, required"`
	HOST     string    `env:"SERVER_HOST, required"`
	Database *Database `env:", prefix=MONGO_"`
	Cache    *Cache    `env:", prefix=REDIS_"`
}

type Database struct {
	Host     string `env:"HOST, required"`
	Name     string `env:"NAME, required"`
	Port     string `env:"PORT, required"`
	Username string `env:"USERNAME, required"`
	Password string `env:"PASSWORD, required"`
}

func (d *Database) URI() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s", d.Username, d.Password, d.Host, d.Port)
}

type Cache struct {
	Addr     string `env:"ADDR"`
	Password string `env:"PASSWORD"`
	DB       string `env:"DB"`
}

func NewConfig() *Config {
	var cfg Config
	if err := envconfig.Process(context.Background(), &envconfig.Config{
		Target:   &cfg,
		Lookuper: envconfig.OsLookuper(),
	}); err != nil {
		panic(err)
	}

	return &cfg
}
