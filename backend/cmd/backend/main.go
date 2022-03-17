package main

import (
	"github.com/brpaz/echozap"
	"github.com/labstack/echo/v4"
	"github.com/miminame-dev/miminame/backend/controller"
	"github.com/miminame-dev/miminame/backend/pkg/config"
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

	props, err := InitProps(cfg)
	if err != nil {
		zap.S().Fatalf("failed to initialize props: %+v", err)
	}

	_ = controller.NewController(props)

	e := echo.New()

	e.Use(echozap.ZapLogger(logger))

	e.Logger.Fatal(e.Start(":1323"))
}
