package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/miminame-dev/miminame/backend/controller"
	"github.com/miminame-dev/miminame/backend/pkg/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %+v", err)
	}

	props, err := InitProps(cfg)
	if err != nil {
		log.Fatalf("failed to initialize props: %+v", err)
	}

	_ = controller.NewController(props)

	e := echo.New()
	e.Logger.Fatal(e.Start(":1323"))
}
