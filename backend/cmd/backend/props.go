package main

import (
	"fmt"

	"github.com/miminame-dev/miminame/backend/pkg/config"
	"github.com/miminame-dev/miminame/backend/pkg/props"
	"github.com/miminame-dev/miminame/backend/service"
	"github.com/streadway/amqp"
	"golang.org/x/xerrors"
)

func InitProps(cfg *config.Config) (*props.Props, error) {
	amqpConn, err := amqp.Dial(fmt.Sprintf(
		"amqp://%s:%s@%s:5672/",
		cfg.RabbitMQUser,
		cfg.RabbitMQPassword,
		cfg.RabbitMQHost,
	))
	if err != nil {
		return nil, xerrors.Errorf("failed to dial amqp: %w", err)
	}

	processVideoService := service.NewProcessVideoService(amqpConn)

	return &props.Props{
		Config:              cfg,
		ProcessVideoService: processVideoService,
	}, nil
}
