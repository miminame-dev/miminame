package main

import (
	"context"
	"encoding/json"

	"cloud.google.com/go/pubsub"
	"github.com/miminame-dev/miminame/video-processor/pkg/config"
	"github.com/miminame-dev/miminame/video-processor/pkg/message"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	cfg, err := config.Load()
	if err != nil {
		zap.S().Fatalf("failed to load config: %+v", err)
	}

	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, cfg.ProjectID)
	if err != nil {
		zap.S().Fatalf("failed to create client: %+v", err)
	}

	defer client.Close()

	sub := client.Subscription("process-video-sub")

	sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		defer m.Ack()

		msg := new(message.ProcessVideoMessage)
		if err := json.Unmarshal(m.Data, msg); err != nil {
			zap.S().Errorf("failed to unmarshal json: %+v", err)
			return
		}

		zap.S().Infof("Start process: %s", msg.Name)
		process(msg.Name)
	})
}

func process(name string) error {
	return nil
}
