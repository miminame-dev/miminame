package service

import (
	"github.com/streadway/amqp"
	"golang.org/x/xerrors"
)

type IProcessVideoService interface {
	StartProcess(videoID string) error
}

type ProcessVideoService struct {
	AMQPConn *amqp.Connection
}

func NewProcessVideoService(amqpConn *amqp.Connection) IProcessVideoService {
	return &ProcessVideoService{
		AMQPConn: amqpConn,
	}
}

func (s *ProcessVideoService) StartProcess(videoID string) error {
	ch, err := s.AMQPConn.Channel()
	if err != nil {
		return xerrors.Errorf("failed to open channel: %w", err)
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"ProcessVideo",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return xerrors.Errorf("failed to declare a queue: %w", err)
	}

	body := videoID

	if err := ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	); err != nil {
		return xerrors.Errorf("failed to publish a message: %w", err)
	}

	return nil
}
