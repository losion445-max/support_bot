package config

import (
	"errors"
	"os"
)

type Config struct {
	Token string
}

func LoadConfig() (*Config, error) {
	token := os.Getenv("TOKEN")
	if token == "" {
		return nil, errors.New("TOKEN env not defined")
	}

	return &Config{
		Token: token,
	}, nil
}
