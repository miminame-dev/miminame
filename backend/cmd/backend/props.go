package main

import (
	"github.com/miminame-dev/miminame/backend/pkg/config"
	"github.com/miminame-dev/miminame/backend/pkg/props"
)

func InitProps(cfg *config.Config) (*props.Props, error) {
	return &props.Props{
		Config: cfg,
	}, nil
}
