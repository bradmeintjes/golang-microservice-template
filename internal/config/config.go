package config

import (
	env "github.com/Netflix/go-env"
)

type Config struct {
	HTTP     HTTP
	Postgres Postgres
}

func Parse() (Config, error) {
	var c Config
	_, err := env.UnmarshalFromEnviron(&c)
	return c, err
}

type HTTP struct {
	Host string `env:"HTTP_HOST,default=localhost"`
	Port int    `env:"HTTP_PORT,default=8080"`
}

type Postgres struct {
	DBName   string `env:"POSTGRES_DBNAME,required=true"`
	Host     string `env:"POSTGRES_HOST,required=true"`
	Port     int    `env:"POSTGRES_PORT,default=5432"`
	User     string `env:"POSTGRES_USER,required=true"`
	Password string `env:"POSTGRES_PASSWORD,required=true"`
	SSL      string `env:"POSTGRES_SSL,default=enabled"`
}
