package config

import (
	"github.com/Netflix/go-env"
)

type Configuration struct {
	DB_USER     string `env:"MYSQL_USER"`
	DB_PASSWORD string `env:"MYSQL_PASSWORD"`
	DB_HOST     string `env:"MYSQL_HOST"`
	DB_PORT     string `env:"MYSQL_PORT"`
	DB_DATABASE string `env:"MYSQL_DATABASE"`
}

func Environment() *Configuration {
	var cfg Configuration
	_, err := env.UnmarshalFromEnviron(&cfg)
	if err != nil {
		panic(err)
	}

	return &cfg
}
