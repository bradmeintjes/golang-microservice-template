package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	HTTP     HTTP
	Postgres Postgres
}

func Parse() (Config, error) {
	var c Config
	return c, env.Parse(&c)
}

type HTTP struct {
	Host string `env:"HTTP_HOST,required" envDefault:"localhost"`
	Port int    `env:"HTTP_PORT,required" envDefault:"8080"`
}

type Postgres struct {
	DBName   string `env:"POSTGRES_DBNAME,required" envDefault:"postgres"`
	Host     string `env:"POSTGRES_HOST,required" envDefault:"localhost"`
	Port     int    `env:"POSTGRES_PORT,required" envDefault:"5432"`
	User     string `env:"POSTGRES_USER,required" envDefault:"postgres"`
	Password string `env:"POSTGRES_PASSWORD,required" envDefault:"pass"`
	SSL      string `env:"POSTGRES_SSL,required" envDefault:"disable"`
}
