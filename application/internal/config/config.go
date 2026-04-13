package config

import (
	"errors"
	"os"
)

type Config struct {
	Token   string
	AdminID string
}

func LoadConfig() (*Config, error) {
	token := os.Getenv("TOKEN")
	if token == "" {
		return nil, errors.New("TOKEN env not defined")
	}

	adminID := os.Getenv("ADMIN_ID")
	if adminID == "" {
		return nil, errors.New("ADMIN_ID env not defined")
	}

	return &Config{
		Token:   token,
		AdminID: adminID,
	}, nil
}
