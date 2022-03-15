package controller

import "github.com/miminame-dev/miminame/backend/pkg/props"

type Controller struct {
	props *props.Props
}

func NewController(p *props.Props) *Controller {
	return &Controller{
		props: p,
	}
}
