package main

import (
	"context"
	"encoding/json"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/miminame-dev/miminame/pkg/message"
	"github.com/miminame-dev/miminame/video-processor/pkg/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %+v", err)
	}

	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, cfg.ProjectID)
	if err != nil {
		log.Fatalf("failed to create client: %+v", err)
	}

	defer client.Close()

	sub := client.Subscription("process-video-sub")

	sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		defer m.Ack()

		msg := new(message.ProcessVideoMessage)
		if err := json.Unmarshal(m.Data, msg); err != nil {
			log.Println("failed to unmarshal json: %+v", err)
			return
		}

		log.Println("Start process:", msg.VideoID)
		process(msg.VideoID)
	})
}

func process(videoID string) error {
	return nil
}
