package config

import (
	"context"

	"github.com/heetch/confita"
	"golang.org/x/xerrors"
)

type Config struct {
	RabbitMQHost     string `config:"RABBITMQ_HOST"`
	RabbitMQUser     string `config:"RABBITMQ_USER"`
	RabbitMQPassword string `config:"RABBITMQ_PASSWORD"`
}

func Load() (*Config, error) {
	cfg := new(Config)

	loader := confita.NewLoader()

	if err := loader.Load(context.Background(), cfg); err != nil {
		return nil, xerrors.Errorf("failed to load config: %w", err)
	}

	return cfg, nil
}
