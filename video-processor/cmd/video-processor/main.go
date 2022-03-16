package main

import (
	"fmt"
	"log"

	"github.com/miminame-dev/miminame/video-processor/pkg/config"
	"github.com/streadway/amqp"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %+v", err)
	}

	conn, err := amqp.Dial(fmt.Sprintf(
		"amqp://%s:%s@%s:5672/",
		cfg.RabbitMQUser,
		cfg.RabbitMQPassword,
		cfg.RabbitMQHost,
	))
	if err != nil {
		log.Fatalf("failed to open connection: %+v", err)
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("failed to open channel: %+v", err)
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
		log.Fatalf("failed to declare a queue: %+v", err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("failed to register a consumer: %+v", err)
	}

	go func() {
		for d := range msgs {
			videoID := string(d.Body)
			log.Println("Start process:", videoID)
			if err := process(videoID); err != nil {
				log.Printf("failed to process %s: %+v", videoID, err)
			}
		}
	}()

	forever := make(chan struct{})
	<-forever
}

func process(videoID string) error {
	return nil
}
