package pkg

import (
	"github.com/go-playground/validator/v10"
	"github.com/ilyakaznacheev/cleanenv"
	"net"
	"strconv"
)

type Config struct {
	Host     string `env:"HOST" json:"host" env-default:"0.0.0.0""`
	Port     int    `env:"PORT" json:"port" env-default:"6379"`
	Password string `env:"PASSWORD" json:"password" env-default:"password"`
}

func (cfg *Config) Address() string {
	return net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port))
}

func LoadConfig(validator *validator.Validate) (*Config, error) {
	var cfg struct {
		Config Config `json:"redis" env-prefix:"REDIS_"`
	}
	err := cleanenv.ReadConfig("config.json", &cfg)

	if err != nil {
		err := cleanenv.ReadEnv(&cfg)
		if err != nil {
			return nil, err
		}
	}
	return &cfg.Config, nil
}
