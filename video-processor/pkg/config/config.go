package config

import (
	"context"

	"github.com/heetch/confita"
	"golang.org/x/xerrors"
)

type Config struct {
	ProjectID string `config:"PROJECT_ID"`
}

func Load() (*Config, error) {
	cfg := new(Config)

	loader := confita.NewLoader()

	if err := loader.Load(context.Background(), cfg); err != nil {
		return nil, xerrors.Errorf("failed to load config: %w", err)
	}

	return cfg, nil
}
