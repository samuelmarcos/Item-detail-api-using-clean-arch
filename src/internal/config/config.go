package config

import "github.com/Netflix/go-env"

type Configuration struct {
	MYSQL_USER     string `env:"MYSQL_USER"`
	MYSQL_PASSWORD string `env:"MYSQL_PASSWORD"`
	MYSQL_HOST     string `env:"MYSQL_HOST"`
	MYSQL_PORT     string `env:"MYSQL_PORT"`
	MYSQL_DATABASE string `env:"MYSQL_DATABASE"`
}

func Environment() *Configuration {
	var cfg Configuration
	_, err := env.UnmarshalFromEnviron(&cfg)
	if err != nil {
		panic(err)
	}

	return &cfg
}
